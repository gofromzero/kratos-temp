// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2021/10/25

package server

import (
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	etcd "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/gofromzero/kratos-temp/app/user/service/internal/conf"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewEtcdRegistrar)

func NewConsulRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewEtcdRegistrar(conf *conf.Registry) registry.Registrar {
	c := clientv3.Config{
		Endpoints: []string{conf.Etcd.Address},
	}
	cli, err := clientv3.New(c)
	if err != nil {
		panic(err)
	}

	r := etcd.New(cli, etcd.MaxRetry(5))
	return r
}
