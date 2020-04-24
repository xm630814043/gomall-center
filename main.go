package main

import (
	"fmt"
	"gomall-center/pkg/settings"
	"gomall-center/pkg/storage"
	"gomall-center/pkg/token"
	"gomall-center/routers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	settings.Setup()
	storage.Setup()
	token.Setup()
}

func main() {
	gin.SetMode(settings.AppConfig.Server.RunMode)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.AppConfig.Server.HTTPPort),
		Handler:        routers.Init(),
		ReadTimeout:    settings.AppConfig.Server.ReadTimeout,
		WriteTimeout:   settings.AppConfig.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("[info] start http server listening: %s", server.Addr)

	_ = server.ListenAndServe()
}
