package server

import (
	"log"

	handler "captioner.com.ng/internal/captioner/handler/healthcheck"
	"github.com/gin-gonic/gin"
)

func Start(addr *string) {
	g := gin.Default()

	g.GET("/healthcheck", handler.Healthcheck)

	err := g.Run(*addr)
	if err != nil {
		log.Fatal(err)
	}
}
