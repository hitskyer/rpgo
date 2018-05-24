package main

import (
	"fmt"
)

/*
 * 编译命令：go build -ldflags "-X main.BuildVersion=xxx -X main.BuildTime=`date +%Y-%m-%d_%H:%M:%S` -X main.BuildComment=yyy"
 * 关于"-ldfflags -X"的说明：https://studygolang.com/articles/2052
 */
var (
	BuildVersion string
	BuildTime    string
	BuildComment string
)

func version() {
	fmt.Printf("main.BuildVersion = %s\n", BuildVersion)
	fmt.Printf("main.BuildTime    = %s\n", BuildTime)
	fmt.Printf("main.BuildComment = %s\n", BuildComment)
}

func main() {
	fmt.Println("Starting rpserver!")

	// 显示版本信息
	version()
}
