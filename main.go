package main

import (
	//_ "brocaedu/Commands"
	"brocaedu/Pkg/setting"
	"brocaedu/Router"
	"brocaedu/Services"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"os"
	"syscall"
)

func main() {
	crontab := Services.NewCronTab()
	// 实现接口的方式添加定时任务

	if err := crontab.AddByFunc("1", "30 21 * * *", func() {
		Services.GetArt()
	}); err != nil {
		fmt.Printf("error to add crontab task:%s", err)
		os.Exit(-1)
	}
	crontab.Start()

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	server := endless.NewServer(endPoint, Router.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
