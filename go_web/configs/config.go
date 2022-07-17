package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type AppConf struct {
	ServerConf ServerConf `yaml:"server"`
	MysqlConf  MysqlConf  `yaml:"mysql"`
}

type ServerConf struct {
	Port int `yaml:"port"`
}

type MysqlConf struct {
	Url       string `yaml:"url"`
	DB        string `yaml:"db"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	MaxActive int    `yaml:"maxActive"`
}

var appConf AppConf

func GetInstance() *AppConf {
	return &appConf
}

func init() {
	bytes, err := ioutil.ReadFile("./configs/application.yaml")
	if err != nil {
		log.Fatalln("read application.yaml err:", err)
	}
	if err = yaml.Unmarshal(bytes, &appConf); err != nil {
		log.Fatalln("parse application.yaml err:", err)
	}
}
