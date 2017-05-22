package singleton

//Config 配置文件管理类
type Config struct {
	ID   int
	Name string
}

var config *Config

//单例模式文章
//http://blog.csdn.net/qibin0506/article/details/50733314
