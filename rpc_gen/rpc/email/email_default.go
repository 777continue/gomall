package email

import (
	"context"

	email "github.com/777continue/gomall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Send(ctx context.Context, req *email.EmailReq, callOptions ...callopt.Option) (resp *email.EmailResp, err error) {
	resp, err = Client.Send(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Send call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
