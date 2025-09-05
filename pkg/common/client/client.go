// Package client 提供HTTP客户端功能
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/onepiecelover/adata-go/pkg/common/errors"
	"github.com/onepiecelover/adata-go/pkg/common/headers"
)

// ProxyConfig 代理配置
type ProxyConfig struct {
	Enabled   bool
	ProxyURL  string
	ProxyList []string
	mu        sync.RWMutex
}

// Client HTTP客户端
type Client struct {
	client      *resty.Client
	proxyConfig *ProxyConfig
	retryTimes  int
	waitTime    time.Duration
}

// NewClient 创建新的HTTP客户端
func NewClient() *Client {
	client := resty.New()

	// 设置超时时间
	client.SetTimeout(30 * time.Second)

	// 设置默认请求头
	client.SetHeaders(headers.GetCommonHeaders())

	// 设置重试机制
	client.SetRetryCount(3)
	client.SetRetryWaitTime(1 * time.Second)
	client.SetRetryMaxWaitTime(5 * time.Second)

	return &Client{
		client:      client,
		proxyConfig: &ProxyConfig{},
		retryTimes:  3,
		waitTime:    100 * time.Millisecond,
	}
}

// SetProxy 设置代理
func (c *Client) SetProxy(enabled bool, proxyURL string) {
	c.proxyConfig.mu.Lock()
	defer c.proxyConfig.mu.Unlock()

	c.proxyConfig.Enabled = enabled
	c.proxyConfig.ProxyURL = proxyURL

	if enabled && proxyURL != "" {
		c.client.SetProxy(proxyURL)
	} else {
		c.client.RemoveProxy()
	}
}

// SetProxyList 设置代理列表
func (c *Client) SetProxyList(proxyList []string) {
	c.proxyConfig.mu.Lock()
	defer c.proxyConfig.mu.Unlock()

	c.proxyConfig.ProxyList = proxyList
}

// getRandomProxy 获取随机代理
func (c *Client) getRandomProxy() string {
	c.proxyConfig.mu.RLock()
	defer c.proxyConfig.mu.RUnlock()

	if len(c.proxyConfig.ProxyList) == 0 {
		return c.proxyConfig.ProxyURL
	}

	rand.Seed(time.Now().UnixNano())
	return c.proxyConfig.ProxyList[rand.Intn(len(c.proxyConfig.ProxyList))]
}

// SetRetryConfig 设置重试配置
func (c *Client) SetRetryConfig(retryTimes int, waitTime time.Duration) {
	c.retryTimes = retryTimes
	c.waitTime = waitTime
}

// Get 发送GET请求
func (c *Client) Get(url string, params map[string]string, customHeaders map[string]string) (*http.Response, []byte, error) {
	return c.request("GET", url, params, nil, customHeaders)
}

// Post 发送POST请求
func (c *Client) Post(url string, data interface{}, customHeaders map[string]string) (*http.Response, []byte, error) {
	return c.request("POST", url, nil, data, customHeaders)
}

// GetJSON 发送GET请求并解析JSON响应
func (c *Client) GetJSON(url string, params map[string]string, customHeaders map[string]string, result interface{}) error {
	_, body, err := c.Get(url, params, customHeaders)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, result); err != nil {
		return errors.NewADataError(errors.ErrParseResponseFailed.Code, "JSON解析失败", err.Error())
	}

	return nil
}

// PostJSON 发送POST请求并解析JSON响应
func (c *Client) PostJSON(url string, data interface{}, customHeaders map[string]string, result interface{}) error {
	_, body, err := c.Post(url, data, customHeaders)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, result); err != nil {
		return errors.NewADataError(errors.ErrParseResponseFailed.Code, "JSON解析失败", err.Error())
	}

	return nil
}

// request 通用请求方法
func (c *Client) request(method, reqURL string, params map[string]string, data interface{}, customHeaders map[string]string) (*http.Response, []byte, error) {
	var lastErr error

	for i := 0; i < c.retryTimes; i++ {
		if i > 0 {
			// 等待后重试
			time.Sleep(c.waitTime * time.Duration(i))
		}

		// 创建请求
		req := c.client.R()

		// 设置自定义请求头
		if customHeaders != nil {
			req.SetHeaders(customHeaders)
		}

		// 设置代理
		if c.proxyConfig.Enabled {
			proxyURL := c.getRandomProxy()
			if proxyURL != "" {
				c.client.SetProxy(proxyURL)
			}
		}

		// 设置查询参数
		if params != nil {
			req.SetQueryParams(params)
		}

		// 设置请求体
		if data != nil {
			req.SetBody(data)
		}

		// 发送请求
		var resp *resty.Response
		var err error

		switch strings.ToUpper(method) {
		case "GET":
			resp, err = req.Get(reqURL)
		case "POST":
			resp, err = req.Post(reqURL)
		case "PUT":
			resp, err = req.Put(reqURL)
		case "DELETE":
			resp, err = req.Delete(reqURL)
		default:
			return nil, nil, fmt.Errorf("unsupported HTTP method: %s", method)
		}

		if err != nil {
			lastErr = err
			continue
		}

		// 检查HTTP状态码
		if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
			return &http.Response{
				StatusCode: resp.StatusCode(),
				Header:     resp.Header(),
			}, resp.Body(), nil
		}

		// 如果是404，直接返回，不重试
		if resp.StatusCode() == 404 {
			return &http.Response{
				StatusCode: resp.StatusCode(),
				Header:     resp.Header(),
			}, resp.Body(), errors.NewADataError(errors.ErrNoDataFound.Code, "数据不存在", fmt.Sprintf("HTTP %d", resp.StatusCode()))
		}

		lastErr = fmt.Errorf("HTTP %d: %s", resp.StatusCode(), string(resp.Body()))
	}

	return nil, nil, errors.WrapError(lastErr, "请求失败")
}

// GetWithContext 带上下文的GET请求
func (c *Client) GetWithContext(ctx context.Context, url string, params map[string]string, customHeaders map[string]string) (*http.Response, []byte, error) {
	// 创建请求
	req := c.client.R()
	req.SetContext(ctx)

	// 设置自定义请求头
	if customHeaders != nil {
		req.SetHeaders(customHeaders)
	}

	// 设置查询参数
	if params != nil {
		req.SetQueryParams(params)
	}

	// 发送请求
	resp, err := req.Get(url)
	if err != nil {
		return nil, nil, errors.WrapError(err, "请求失败")
	}

	return &http.Response{
		StatusCode: resp.StatusCode(),
		Header:     resp.Header(),
	}, resp.Body(), nil
}

// DownloadFile 下载文件
func (c *Client) DownloadFile(url, filepath string) error {
	resp, err := c.client.R().SetOutput(filepath).Get(url)
	if err != nil {
		return errors.WrapError(err, "下载文件失败")
	}

	if resp.StatusCode() != 200 {
		return errors.NewADataError(errors.ErrRequestFailed.Code, "下载失败", fmt.Sprintf("HTTP %d", resp.StatusCode()))
	}

	return nil
}

// GetText 获取文本响应
func (c *Client) GetText(url string, params map[string]string, customHeaders map[string]string) (string, error) {
	_, body, err := c.Get(url, params, customHeaders)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Close 关闭客户端
func (c *Client) Close() error {
	// resty客户端不需要显式关闭
	return nil
}

// 全局默认客户端
var defaultClient = NewClient()

// Get 使用默认客户端发送GET请求
func Get(url string, params map[string]string, customHeaders map[string]string) (*http.Response, []byte, error) {
	return defaultClient.Get(url, params, customHeaders)
}

// Post 使用默认客户端发送POST请求
func Post(url string, data interface{}, customHeaders map[string]string) (*http.Response, []byte, error) {
	return defaultClient.Post(url, data, customHeaders)
}

// GetJSON 使用默认客户端发送GET请求并解析JSON
func GetJSON(url string, params map[string]string, customHeaders map[string]string, result interface{}) error {
	return defaultClient.GetJSON(url, params, customHeaders, result)
}

// PostJSON 使用默认客户端发送POST请求并解析JSON
func PostJSON(url string, data interface{}, customHeaders map[string]string, result interface{}) error {
	return defaultClient.PostJSON(url, data, customHeaders, result)
}

// GetText 使用默认客户端获取文本响应
func GetText(url string, params map[string]string, customHeaders map[string]string) (string, error) {
	return defaultClient.GetText(url, params, customHeaders)
}

// SetProxy 设置默认客户端代理
func SetProxy(enabled bool, proxyURL string) {
	defaultClient.SetProxy(enabled, proxyURL)
}

// SetProxyList 设置默认客户端代理列表
func SetProxyList(proxyList []string) {
	defaultClient.SetProxyList(proxyList)
}
