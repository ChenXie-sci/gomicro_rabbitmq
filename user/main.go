package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"main/conf"
	"main/core"
	"main/services"
)

func main() {
	conf.Init()
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg), // etcd register
	)

	microService.Init()

	_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	_ = microService.Run()
}
