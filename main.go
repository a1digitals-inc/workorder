package main

import (
	"fmt"
	"ginapi/pkg/setting"
	"ginapi/routers"
	"net/http"
)

func main() {
	routers := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTP_PORT),
		Handler:        routers,
		ReadTimeout:    setting.READ_TIMEOUT,
		WriteTimeout:   setting.WRITE_TIMEOUT,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
