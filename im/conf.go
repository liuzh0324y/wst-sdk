package im

import (
	"log"

	"github.com/wst-libs/wst-sdk/conf"
)

// IMConf is a struct
type IMConf struct {
	Server struct {
		Appname   string `yaml:"appname"`
		Httpport  string `yaml:"httpport"`
		Runmodel  string `yaml:"runmode"`
		Copy      bool   `yaml:"copyrequestbody"`
		Endpoint  string `yaml:"endpoint"`
		AccessKey string `yaml:"accesskey"`
		SecretKey string `yaml:"secretkey"`
		Bucket    string `yaml:"bucket"`
		FilePath  string `yaml:"filepath"`
		PutHost   string `yaml:"puthost"`
		PutPort   string `yaml:"putport"`
		PutPath   string `yaml:"putpath"`
	} `yaml:"server"`
}

const (
	url = "http://39.105.53.16:48888/im-dev.yml"
)

var config = IMConf{}

func getconf() {
	err := conf.GetConf(url, config)
	if err != nil {
		log.Println(err.Error())
	}
}
