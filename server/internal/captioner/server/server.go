package server

import (
	"log"

	handler "captioner.com.ng/internal/captioner/handler/healthcheck"
	"captioner.com.ng/internal/captioner/router"
	"captioner.com.ng/internal/captioner/store"
	"github.com/gin-gonic/gin"
)

func Start(addr *string) {
	g := gin.Default()
	db := store.ConnectDB()

	g.GET("/healthcheck", handler.Healthcheck)
	router.Connect(g, db)
	err := g.Run(*addr)
	if err != nil {
		log.Fatal(err, db)
	}
}
