package checkout

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	//"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestCheckout(t *testing.T) {
	h := server.Default()
	h.POST("/checkout", Checkout)

	// 构造测试请求体
	reqBody := map[string]interface{}{
		"email":   "test@example.com",
		"address": "123 Main St",
		"name":    "yy",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	path := "/checkout"
	body := &ut.Body{Body: bytes.NewBuffer(bodyBytes), Len: len(bodyBytes)}
	// 正确设置header
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/json",
	}

	// 执行请求
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()

	// 验证响应
	assert.DeepEqual(t, 200, resp.StatusCode())

	// 解析响应体
	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body(), &respBody)

}
