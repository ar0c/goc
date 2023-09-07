package tests

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func diff(target string) {
	// 创建一个命令对象
	cmd := exec.Command("git", "diff", target)

	// 创建一个字节缓冲区来存储输出结果
	var output bytes.Buffer

	// 将命令的输出连接到字节缓冲区
	cmd.Stdout = &output

	// 执行命令
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// 打印输出结果
	fmt.Println(output.String())
}
