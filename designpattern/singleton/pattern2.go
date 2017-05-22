package singleton

import "sync"

var lock *sync.Mutex = &sync.Mutex{}

//GetInstance2 加锁方式实现单例模式
func GetInstance2() *Config {
	lock.Lock()
	defer lock.Unlock()
	if config == nil {
		config = &Config{}
	}
	return config
}
