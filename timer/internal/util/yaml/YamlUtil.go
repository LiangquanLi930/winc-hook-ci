package yaml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"timer/internal/util/err"
)

var conf config

func GetConfig() config {
	return conf
}

type config struct {
	Quay struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"quay"`
	SlackMessageUrl string `yaml:"slack_message_url"`
}

func Init(filePath string) {
	var erro error
	yamlFile, erro := ioutil.ReadFile(filePath)
	err.GetErr("无法读取客户端文件", erro)
	erro = yaml.Unmarshal(yamlFile, &conf)
	err.GetErr("解析yaml错误", erro)
}
