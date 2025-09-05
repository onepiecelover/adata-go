package tests

import (
	"errors"
	"testing"

	adataErrors "github.com/onepiecelover/adata-go/pkg/common/errors"
	"github.com/stretchr/testify/assert"
)

func TestADataError_Error(t *testing.T) {
	// 测试带详细信息的错误
	err1 := &adataErrors.ADataError{
		Code:    10001,
		Message: "测试错误",
		Detail:  "详细错误信息",
	}
	expected1 := "[10001] 测试错误: 详细错误信息"
	assert.Equal(t, expected1, err1.Error())

	// 测试不带详细信息的错误
	err2 := &adataErrors.ADataError{
		Code:    10002,
		Message: "测试错误2",
	}
	expected2 := "[10002] 测试错误2"
	assert.Equal(t, expected2, err2.Error())
}

func TestNewADataError(t *testing.T) {
	err := adataErrors.NewADataError(10001, "测试错误", "详细信息")

	assert.Equal(t, 10001, err.Code)
	assert.Equal(t, "测试错误", err.Message)
	assert.Equal(t, "详细信息", err.Detail)
}

func TestWrapError(t *testing.T) {
	originalErr := errors.New("原始错误")
	wrappedErr := adataErrors.WrapError(originalErr, "包装错误")

	assert.Equal(t, 99999, wrappedErr.Code)
	assert.Equal(t, "包装错误", wrappedErr.Message)
	assert.Equal(t, "原始错误", wrappedErr.Detail)
}

func TestPredefinedErrors(t *testing.T) {
	// 测试预定义错误
	assert.Equal(t, 10001, adataErrors.ErrInvalidStockCode.Code)
	assert.Equal(t, 10002, adataErrors.ErrInvalidDateFormat.Code)
	assert.Equal(t, 20001, adataErrors.ErrRequestFailed.Code)
	assert.Equal(t, 20002, adataErrors.ErrParseResponseFailed.Code)
	assert.Equal(t, 30001, adataErrors.ErrNoDataFound.Code)
	assert.Equal(t, 30002, adataErrors.ErrDataSourceUnavailable.Code)
}
