package main

import (
	"fmt"
	"os"
	"path/filepath"
	"ratel/config"
	taskP "ratel/task"
)

func main() {
	// 当前执行路径
	currentPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("current exec path is ", currentPath)

	// 读取conf文件夹下配置
	conf, err := config.LoadConfig(currentPath)
	if err != nil {
		fmt.Print("load config err ", err)
		os.Exit(1)
	}

	// 获取任务
	tasks := conf.Tasks
	for _, task := range tasks {
		taskP.DoTask(&task)
	}
}
