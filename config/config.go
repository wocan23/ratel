package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"ratel/common"
	"ratel/model"
	"strings"
)

func LoadConfig(path string) (*model.Config,error){
	var config = new(model.Config)
	var err error
	configByte,err := ioutil.ReadFile(path +"/conf/main.yml")
	if err != nil{
		fmt.Print("load config error",err)
		return config,err
	}

	err = yaml.Unmarshal(configByte,config)
	if err != nil{
		fmt.Print("parse config error",err)
		return config,err
	}
	flushConfig(config)
	return config,err
}

func flushConfig(config *model.Config){
	var vars = config.Vars
	var tasks = config.Tasks
	for _,task := range tasks{
		// 替换变量
		var taskVas = task.Vars
		var taskFinalVars = common.CombineMap(vars,taskVas)
		var steps = task.Steps
		for i,step := range steps{
			var str = replaceVars(&step,taskFinalVars)
			steps[i] = str
		}
	}
}

func replaceVars(step *string, vars map[string]string) string{
	var str = *step
	for k,v := range vars{
		str = strings.ReplaceAll(str,"$"+k,v)
	}
	return str
}
