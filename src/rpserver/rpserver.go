package main

import (
	"fmt"
	"flag"
	"runtime"
)

/*
 * 编译命令：go build -ldflags "-X main.BuildVersion=xxx -X main.BuildTime=`date +%Y-%m-%d_%H:%M:%S` -X main.BuildComment=yyy"
 * 关于"-ldfflags -X"的说明：https://studygolang.com/articles/2052
 */
var (
	BuildVersion string
	BuildTime    string
	BuildComment string
	gConfigFile  = flag.String("config_file", "./rpserver.conf", "input rpserver config file name")
)

func version() {
	fmt.Printf("main.BuildVersion  = %s\n", BuildVersion)
	fmt.Printf("main.BuildTime     = %s\n", BuildTime)
	fmt.Printf("main.BuildComment  = %s\n", BuildComment)
}
func init() {
	// 仅仅因为好奇
	//flag.Set("config_file", "none")
}

func main() {
	fmt.Println("Starting rpserver!")

	// 显示版本信息
	version()
	// 解析命令行参数
	flag.Parse()
	fmt.Printf("config_file        : %s\n", *gConfigFile)
	// 最大化使用cpu
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("the number of cpus : %d\n", runtime.NumCPU())
	// 启动服务
	StartHttpServer("localhost:9100")
}
