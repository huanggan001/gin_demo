package main

import (
	"flag"
	"gin_demo/common/lib"
	"gin_demo/router"
	"os"
	"os/signal"
	"syscall"
)

//endpoint dashboard后台管理 server代理服务器
//config ./conf/dev/ 对应配置文件夹

// go run main.go -config="./conf/dev/" -endpoint="dashboard"
// go run main.go -config="./conf/dev/" -endpoint="server"
var (
	//endpoint = flag.String("endpoint", "", "input dashboard or server")
	config = flag.String("config", "", "input config file like ./conf/dev/")
)

func main() {
	flag.Parse()
	lib.InitModule(*config)
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.HttpServerStop()
}
