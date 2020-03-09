package main

import (
	"example.com/m/pkg/setting"
	"example.com/m/routers"
	"fmt"
	"net/http"
)

func main() {
	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPonint := fmt.Sprintf(":%d",setting.HTTPPort)
	//
	//server := endless.NewServer(endPonint,routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d",syscall.Getpid())
	//}
	//err := server.ListenAndServe()
	//if err!=nil {
	//	log.Printf("server err:%v",err)
	//}
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
