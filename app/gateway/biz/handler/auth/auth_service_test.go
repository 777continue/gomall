package auth

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	//"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestLogin(t *testing.T) {
	h := server.Default()
	h.POST("/auth/login", Login)
	path := "/auth/login"                                     // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}

func TestRegister(t *testing.T) {
	h := server.Default()
	h.POST("/auth/register", Register)

	// 构造测试请求
	jsonStr := `{"email":"test@example.com","password":"123456","password_confirm":"123456"}`
	body := &ut.Body{
		Body: bytes.NewBufferString(jsonStr),
		Len:  len(jsonStr), // 添加长度信息
	}

	// 正确设置header
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/json",
	}

	// 执行请求
	w := ut.PerformRequest(h.Engine, "POST", "/auth/register", body, header)
	resp := w.Result()

	// 输出结果
	t.Logf("Status: %d, Body: %s", resp.StatusCode(), string(resp.Body()))
}
