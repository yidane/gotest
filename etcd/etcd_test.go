package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func Test_Dail(t *testing.T) {
	cli, err := clientv3.NewFromURL("http://127.0.0.1:2379")

	if err != nil {
		t.Fatal(err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	//response, err := cli.Get(ctx, "/testdir/testkey")
	//cancel()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//fmt.Println(response)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	resp, err := cli.Put(ctx, "/testdir/sample_key", "sample_value")
	cancel()
	if err != nil {
		// handle error!
	}

	fmt.Println(resp)
}
