package tests

import (
	"testing"

	"github.com/onepiecelover/adata-go"
)

func TestBasicStructure(t *testing.T) {
	// 基本结构测试
	t.Log("AData-Go 基础结构测试")

	// 测试GetVersion函数
	version := adata.GetVersion()
	if version != "1.0.0" {
		t.Errorf("Expected GetVersion() to return 1.0.0, got %s", version)
	}

	t.Log("✅ 基础结构测试通过")
}

func TestModuleInstances(t *testing.T) {
	// 测试模块实例是否正确初始化
	if adata.Stock == nil {
		t.Error("Stock 模块实例不应为 nil")
	}

	if adata.Fund == nil {
		t.Error("Fund 模块实例不应为 nil")
	}

	if adata.Bond == nil {
		t.Error("Bond 模块实例不应为 nil")
	}

	if adata.Sentiment == nil {
		t.Error("Sentiment 模块实例不应为 nil")
	}

	t.Log("✅ 模块实例测试通过")
}

func TestSetProxy(t *testing.T) {
	// 测试代理设置函数
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("SetProxy 函数不应引发 panic: %v", r)
		}
	}()

	// 测试代理设置
	adata.SetProxy(false, "")
	adata.SetProxy(true, "http://proxy.example.com:8080")

	t.Log("✅ 代理设置测试通过")
}
