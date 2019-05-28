package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//heidsoft@qq.com

//test etcd watch key
//
//使用客户端进行测试时,需要增加etcd api 版本，以及api 地址环境变量
//env:
//ETCDCTL_API=3
//ENDPOINTS=http://127.0.0.1:2379
//etcdctl put  name "eeee"
//
//Event received! PUT executed on "name" with value "eeee"

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})

	if err != nil {
		fmt.Printf("conn failed, %s", err.Error())
	}

	defer cli.Close()

	etcdWatchKey := "name"

	//go func() {
	//	fmt.Println("started goroutine for PUT...")
	//	for {
	//		cli.Put(context.Background(), etcdWatchKey, time.Now().String())
	//		fmt.Println("populated " + etcdWatchKey + " with a value..")
	//		time.Sleep(2 * time.Second)
	//	}
	//
	//}()

	fmt.Printf("build watch chan start \n")
	watch_chan := cli.Watch(context.Background(), etcdWatchKey)
	/**
	/usr/local/gopath/src/go.etcd.io/etcd/clientv3/watch.go
	*/
	for wresp := range watch_chan {
		fmt.Printf("wresp %v \n", wresp)
		for _, event := range wresp.Events {
			fmt.Printf("Event received! %s executed on %q with value %q\n", event.Type, event.Kv.Key, event.Kv.Value)
		}
	}

}
