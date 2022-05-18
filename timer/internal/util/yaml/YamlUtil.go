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
	BuildConfig     []struct {
		OcpRelease   string `yaml:"ocp_release"`
		ContainerTag string `yaml:"container_tag"`
	} `yaml:"build_config"`
}

func Init(filePath string) {
	var erro error
	yamlFile, erro := ioutil.ReadFile(filePath)
	err.GetErr("can't read config yaml", erro)
	erro = yaml.Unmarshal(yamlFile, &conf)
	err.GetErr("yaml error", erro)
}
