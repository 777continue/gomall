package service

import (
	"context"

	common "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "coke", "Price": 100, "Picture": "/static/image/coke.jpeg"},
		{"Name": "coke1", "Price": 110, "Picture": "/static/image/coke.jpeg"},
		{"Name": "coke1", "Price": 120, "Picture": "/static/image/coke.jpeg"},
		{"Name": "coke1", "Price": 130, "Picture": "/static/image/coke.jpeg"},
		{"Name": "coke1", "Price": 140, "Picture": "/static/image/coke.jpeg"},
		{"Name": "coke2", "Price": 150, "Picture": "/static/image/coke.jpeg"},
	}
	resp["Title"] = "杨镇泽的小店"
	resp["Items"] = items
	//session := sessions.Default(h.RequestContext)
	//resp["user_id"] = session.Get("user_id")
	return resp, nil
}
