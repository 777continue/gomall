package client

import (
	"sync"

	"github.com/777continue/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		InitUserClient()
	})
}
func InitUserClient() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		hlog.Fatal(err)
	}
	UserClient, err = userservice.NewClient(
		"user",
		client.WithResolver(r),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}