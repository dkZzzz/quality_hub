package etcd

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/dkZzzz/quality_hub/config"
)

var (
	Client *clientv3.Client
)

func init() {
	InitEtcd()
}

func InitEtcd() {
	// 连接etcd
	config := clientv3.Config{
		Endpoints:   []string{config.Cfg.EtcdHost}, // Etcd 服务器地址
		DialTimeout: 5 * time.Second,
	}

	// 创建 Etcd 客户端
	client, err := clientv3.New(config)
	Client = client
	if err != nil {
		log.Fatal(err)
	}
}

func Get(service string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	service = "services" + "/" + service
	resp, err := Client.Get(ctx, service)
	cancel()
	if err != nil {
		return false, err
	}
	if len(resp.Kvs) == 0 {
		return false, nil
	}
	if string(resp.Kvs[0].Value) == "true" {
		return true, nil
	}
	return false, nil
}

func Register(service string) error {
	service = "services" + "/" + service
	_, err := Client.Put(context.Background(), service, "true")
	if err != nil {
		return err
	}
	return nil
}
