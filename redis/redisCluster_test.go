package redis

import (
	"github.com/go-redis/redis"
	"testing"
)

var master = "172.22.0.2:6379"
var clusters = []string{
	"172.22.0.3:6379",
	"172.22.0.4:6379",
	"172.22.0.5:6379",
}

func mergeMasterAdnClusters() []string {
	var result []string
	result = append(result, master)
	result = append(result, clusters...)
	return result
}

//redis服务采取单点部署
//单体应用，不支持高可用，单点崩溃则整个redis服务不可用
func TestRedisSingleClient(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    master,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pong)
}

//redis服务采取主从节点部署
//主节点写，从节点只读
//若主节点崩溃，则整个服务不可用
//为了保证性能，主节点不持久化，从节点可以开启RDB和AOF持久化。缺点：会丢失部分数据，在系统能容忍情况下，可以满足要求。
func TestMaster_cluster(t *testing.T) {
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: mergeMasterAdnClusters(),
	})

	pong, err := clusterClient.Ping().Result()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pong)
}

//哨兵集群模式
//哨兵会监控所有redis节点，包括主节点和从节点
//当从节点不可用，哨兵会把从节点从redis服务中移除
//当主节点不可用，哨兵会从从节点中选取一个节点作为主节点，保证集群高可用
//缺点：当主节点不可用，哨兵正在选举从节点时候，会丢失部分数据
func TestSentinelClient(t *testing.T) {
	sentinelClient := redis.NewSentinelClient(&redis.Options{
		Addr: master,
	})

	pong, err := sentinelClient.Ping().Result()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pong)
}
