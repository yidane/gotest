package singleton

//GetInstance 获取Config的单例
func GetInstance() *Config {
	if config == nil {
		config = &Config{}
	}

	return config
}
