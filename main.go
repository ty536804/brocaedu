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
	//router := Router.InitRouter()
	//srv := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//go func() {
	//	// 服务连接
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("listen: %s\n", err)
	//	}
	//}()
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//
	//<-quit
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}
	//
	//log.Printf("Actual pid is %d", syscall.Getpid())

	crontab := Services.NewCrontab()
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
