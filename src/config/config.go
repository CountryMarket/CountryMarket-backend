package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	C *Config // 全局配置
)

type Config struct {
	App     app     `yaml:"app"`
	LogConf logConf `yaml:"log_conf"`
	MySQL   mysql   `yaml:"mysql"`
	Redis   redis   `yaml:"redis"`
	Jwt     jwt     `yaml:"jwt"`
	Debug   bool    `yaml:"debug"`
}
type app struct {
	Addr      string `yaml:"addr"`
	Root      string `yaml:"root"`
	Appid     string `yaml:"appid"`
	AppSecret string `yaml:"app_secret"`
}
type logConf struct {
	LogPath     string `yaml:"log_path"`
	LogFileName string `yaml:"log_file_name"`
}
type mysql struct {
	Addr     string `yaml:"addr"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type redis struct {
	Addr     string `yaml:"addr"`
	Db       int    `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type jwt struct {
	Secret string `yaml:"secret"`
}

func init() {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Println("Cannot read config!")
		log.Panic(err)
		return
	}
	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Println("Cannot convert config!")
		log.Panic(err)
		return
	}
	C = config
	log.Println("Config was loaded.")
	if C.Debug {
		log.Printf("%+v\n", C)
	}
}
