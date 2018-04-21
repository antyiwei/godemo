package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-a", "-l")
	// 重定向标准输出到文件
	stdout, err := os.OpenFile("stdout.log", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	defer stdout.Close()
	cmd.Stdout = stdout
	// 执行命令
	if err := cmd.Start(); err != nil {
		log.Println(err)
	}
}
