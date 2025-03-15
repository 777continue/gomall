package cart

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	//"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestAddCartItem(t *testing.T) {
	h := server.Default()
	h.POST("/cart", AddCartItem)

	// 构造测试请求
	formData := "productId=1&productNum=2"
	body := &ut.Body{
		Body: bytes.NewBufferString(formData),
		Len:  len(formData),
	}
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}
	w := ut.PerformRequest(h.Engine, "POST", "/cart", body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// 添加断言
	if resp.StatusCode() != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode())
	}
}
func TestGetCart(t *testing.T) {
	h := server.Default()
	h.GET("/cart", GetCart)
	path := "/cart"                                           // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "GET", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}
