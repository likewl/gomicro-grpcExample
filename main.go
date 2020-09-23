package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"gomicro-grpcExample/Services"
	"gomicro-grpcExample/serviceImp"
)


func main() {
	newRegistry := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	service := micro.NewService(
		micro.Name("test2"),
		micro.Address(":8011"),
		micro.Registry(newRegistry),
	)
	service.Init()
	Services.RegisterProdServiceHandler(service.Server(),serviceImp.GetServer())
	service.Run()
}
