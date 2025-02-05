package main

import (
	"context"
	"fmt"

	"github.com/777continue/gomall/demo/thrift/kitex_gen/api"
	"github.com/777continue/gomall/demo/thrift/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}
	c, err := echo.NewClient("thrift", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
	//rpc call server's method
	res, err := c.Echo(context.Background(), &api.Request{Message: "hello"})
	if err != nil {
		klog.Fatal(err)
	}
	fmt.Println(res)
}
