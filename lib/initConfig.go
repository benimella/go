package lib

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Remote struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`
}
type SysConfigStruct struct {
	Remotes []*Remote `yaml:"remotes"` // 远程主机列表
}

func (this *SysConfigStruct) GetRemote(name string) *Remote {
	for _, remote := range this.Remotes {
		if remote.Name == name {
			return remote
		}
	}
	return nil
}

var SysConfig *SysConfigStruct

func init() {
	SysConfig = &SysConfigStruct{}
	f, err := ioutil.ReadFile("./config/app.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(f, SysConfig)
	if err != nil {
		log.Fatal(err)
	}
}
