package payment

import (
	"sync"

	"github.com/cloudwego/kitex/client"
)

var (
	// todo edit custom config
	Client            RPCClient
	defaultDstService = "payment"
	defaultClientOpts = []client.Option{
		client.WithHostPorts("127.0.0.1:8884"),
	}
	once sync.Once
)

func init() {
	DefaultClient()
}

func DefaultClient() RPCClient {
	once.Do(func() {
		Client = newClient(defaultDstService, defaultClientOpts...)
	})
	return Client
}

func newClient(dstService string, opts ...client.Option) RPCClient {
	c, err := NewRPCClient(dstService, opts...)
	if err != nil {
		panic("failed to init client: " + err.Error())
	}
	return c
}

func InitClient(dstService string, opts ...client.Option) {
	Client = newClient(dstService, opts...)
}
