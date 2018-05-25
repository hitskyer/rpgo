#目标
* 用go语言实现一个demo性质的推荐引擎

#技能
* 通过“-ldfflags -X”建立编译时生成的变量
    * 参考：https://studygolang.com/articles/2052
    * 实例：关注BuildVersion、BuildTime、BuildComment
* 通过flag模块解析命令行参数
    * 参考：https://studygolang.com/articles/5608
    * 实例：关注config_file
* 关于init函数
    * 参考：https://blog.csdn.net/zhuxinquan61/article/details/73712251
* 搭建http服务
    * 参考：https://segmentfault.com/a/1190000006812688
    * 实例：关注./src/rpserver/http.go中的StartHttpServer函数
