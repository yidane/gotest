package singleton

//GetInstance3 双重锁实现单例模式
func GetInstance3() *Config {
	if config == nil {
		lock.Lock()
		defer lock.Unlock()
		if config == nil {
			config = &Config{}
		}
	}

	return config
}
