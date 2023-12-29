package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"sync"
)

var (
	config = &Config{
		Http: &Http{
			Host: "0.0.0.0",
			Port: 10001,
		},
		Mysql: &Mysql{
			Host:     "0.0.0.0",
			Port:     3306,
			Username: "demo",
			Password: "demo",
			Database: "demo",
		},
	}
	lock sync.Mutex
)

func init() {
	lock.Lock()
	_, err := toml.DecodeFile("demo-curd/etc/conf.toml", config)
	if err != nil {
		fmt.Println("加载配置文件失败...")
		panic(err)
	}
	lock.Unlock()
}

func GetConfig() *Config {
	return config
}

type Config struct {
	Http  *Http  `toml:"http"`
	Mysql *Mysql `toml:"mysql"`
}

type Http struct {
	Host string `toml:"host" env:"HTTP_HOST"`
	Port int    `toml:"port" env:"HTTP_PORT"`
}

type Mysql struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     int    `toml:"port" env:"MYSQL_PORT"`
	Username string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
}
