package singleton

import (
	"sync"
)

var once sync.Once

//GetInstance4 使用syne.Once只实现一次，实现单例模式
func GetInstance4() *Config {
	once.Do(func() {
		config = &Config{}
	})

	return config
}
