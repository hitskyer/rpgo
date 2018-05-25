package main

import (
	"fmt"
	"io"
//	"strings"
	"time"
	"log"
	"net/http"
)

/*
 * 启动服务
 */
func StartHttpServer(listenOn string) {
	// 设置ServrMux
	// 参考：https://segmentfault.com/a/1190000006812688
	// 替代：DefaultServeMux代替ServerMux
	mux   := http.NewServeMux()

	baidu := http.RedirectHandler("https://www.baidu.com/", 307)
	mux.Handle("/baidu", baidu)                    // 路由1：重定向至baidu，测试之用
	
	hw    := http.HandlerFunc(HelloWorld)
	mux.Handle("/HelloWorld", hw)                  // 路由2：通过函数提供打个招呼服务，测试之用

	th    := &TimeHandlerStruct{format : time.RFC1123}
	mux.Handle("/time/struct", th)                 // 路由3：通过接口提供时间查询服务，测试之用

	mux.HandleFunc("/time/func", TimeHandleFunc)   // 路由4：通过函数提供时间查询服务，测试之用
	//thf   := http.HandlerFunc(TimeHandleFunc)
	//mux.Handle("/time/func", thf)                // 路由4的替代方案

	thfc  := TimeHandleFuncClosure(time.RFC1123)
	mux.Handle("/time/closure", thfc)              // 路由5：通过闭包提供时间查询服务，测试之用

	// 启动监听
	log.Println("Listening...")
	http.ListenAndServe(listenOn, mux)
}
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("Hello, world! BuildVersion is %s, BuildComment is %s, BuildTime is %s", BuildVersion, BuildComment, BuildTime))
}
type TimeHandlerStruct struct {
	format string
}
func (th *TimeHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is "+tm))
}
func TimeHandleFunc(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is "+tm))
}
func TimeHandleFuncClosure(format string) http.Handler {
	thfc := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is : " + tm))
	}
	return http.HandlerFunc(thfc)
}
