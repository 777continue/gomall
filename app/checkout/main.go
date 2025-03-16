package main

import (
	"net"
	"time"

	"github.com/777continue/gomall/app/checkout/conf"
	"github.com/777continue/gomall/app/checkout/infra/mq"
	"github.com/777continue/gomall/common/mtl"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	cart_client "github.com/777continue/gomall/rpc_gen/rpc/cart"
	order_client "github.com/777continue/gomall/rpc_gen/rpc/order"
	payment_client "github.com/777continue/gomall/rpc_gen/rpc/payment"
	product_client "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	serviceName  = conf.GetConf().Kitex.Service
	registryAddr = conf.GetConf().Registry.RegistryAddress[0]
	metricsPort  = conf.GetConf().Kitex.MetricsPort
)

func main() {
	mtl.InitMetric(serviceName, metricsPort, registryAddr)
	mq.Init()

	InitClient()

	opts := kitexInit()

	svr := checkoutservice.NewServer(new(CheckoutServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}
func InitClient() {
	payment_client.DefaultClient()
	product_client.DefaultClient()
	cart_client.DefaultClient()
	order_client.DefaultClient()
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))
	// sevice register
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithServiceAddr(addr), server.WithRegistry(r))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
