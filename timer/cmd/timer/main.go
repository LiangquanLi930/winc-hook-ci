package main

import (
	"github.com/robfig/cron"
	"os/exec"
	"strings"
	"timer/internal/util/log"
	"timer/internal/util/slack"
	"timer/internal/util/yaml"
)

func init() {
	log.Config(log.Stdout, log.Stdout, log.Stdout|log.EnableFile, log.Stderr|log.EnableFile, "./error.log")
	yaml.Init("./config.yaml")
}

func main() {
	log.Info.Println(yaml.GetConfig())
	c := cron.New()
	//spec := "0 0 */1 * * ?"
	spec := "0 0 3 * * ?"
	err := c.AddFunc(spec, func() {
		log.Info.Println("start cron job")
		for _, buildConfig := range yaml.GetConfig().BuildConfig {
			cmd := exec.Command("sh", "-c", `ansible localhost -m include_role -a name=wmco-build -e container_registry_login_user=`+yaml.GetConfig().Quay.User+` -e container_registry_login_token=`+yaml.GetConfig().Quay.Password+` -e ocp_release=`+buildConfig.OcpRelease+` -e container_tag=`+buildConfig.ContainerTag)
			out, err := cmd.CombinedOutput()
			trimmed := strings.TrimSpace(string(out))
			log.Info.Println(trimmed)
			if err != nil {
				log.Warning.Println(trimmed)
				slack.SendSlack("build image error:" + err.Error())
			}else {
				log.Info.Println("build new winc image: quay.io/winc/wmco-index:" + buildConfig.ContainerTag)
			}
			log.Info.Println("job finish")
		}
	})
	if err != nil {
		log.Error.Printf("AddFunc error : %v\n", err)
		return
	}
	c.Start()

	defer c.Stop()
	select {}
}
