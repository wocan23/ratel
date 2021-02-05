package task

import (
	"fmt"
	"os"
	"os/exec"
	"ratel/model"
	"runtime"
	"time"
)

func DoTask(task *model.Task) {
	// 获取操作系统
	fmt.Println("start exec task ", task.Name)
	var startTime = time.Now().Unix()
	for _, step := range task.Steps {
		fmt.Println("exec step ", step)
		var cmd = DoCmd(&step)
		//var
		var out []byte
		var err error
		if out, err = cmd.Output(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(out))
	}
	var endTime = time.Now().Unix()
	fmt.Println(fmt.Sprintf("执行任务[%s]耗费时间[%d]", task.Name, endTime-startTime))
}

func DoCmd(cmdStr *string) *exec.Cmd {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", *cmdStr)
		break
	case "linux":
		cmd = exec.Command("/bin/sh", "-C", *cmdStr)
		break
	case "mac":
	}
	return cmd
}
