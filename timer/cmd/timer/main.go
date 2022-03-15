package main

import (
	"github.com/robfig/cron"
	"os/exec"
	"strings"
	"time"
	"timer/internal/util/log"
	"timer/internal/util/slack"
	"timer/internal/util/yaml"
)

func init() {
	//log.Config(log.Stdout, log.Stdout, log.Stdout|log.EnableFile, log.Stderr|log.EnableFile,"/Users/redhat/GolandProjects/winc-hook-ci/error.log")
	//yaml.Init("/Users/redhat/GolandProjects/winc-hook-ci/config.yaml")
	log.Config(log.Stdout, log.Stdout, log.Stdout|log.EnableFile, log.Stderr|log.EnableFile, "./error.log")
	yaml.Init("./config.yaml")
}

func main() {
	log.Info.Println(yaml.GetConfig())
	c := cron.New()
	//spec := "0 0 */1 * * ?"
	spec := "0 0 4 * * ?"
	err := c.AddFunc(spec, func() {
		log.Info.Println("start cron job")
		tag := "6.0." + time.Now().Format("20060102")
		cmd := exec.Command("sh", "-c", `docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -e TAG="`+tag+`" -e USER="`+yaml.GetConfig().Quay.User+`" -e PASSWORD="`+yaml.GetConfig().Quay.Password+`" quay.io/winc/builder`)
		out, err := cmd.CombinedOutput()
		trimmed := strings.TrimSpace(string(out))
		log.Info.Println(trimmed)
		if err != nil {
			log.Warning.Println(trimmed)
			slack.SendSlack("build image error:" + err.Error())
		}
		log.Info.Println("job finish")
		slack.SendSlack("build new winc image: quay.io/winc/wmco-index:" + tag)
	})
	if err != nil {
		log.Error.Printf("AddFunc error : %v\n", err)
		return
	}
	c.Start()

	defer c.Stop()
	select {}
}
