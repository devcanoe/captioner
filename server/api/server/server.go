package server

import (
	"log"

	"captioner.com.ng/api/server/router"
	handler "captioner.com.ng/internal/captioner/handler/healthcheck"
	"captioner.com.ng/internal/captioner/store"
	"github.com/gin-gonic/gin"
)

func Start(addr *string) {
	g := gin.Default()
	db := store.ConnectDB()
	g.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Methods", "HEAD, OPTIONS, GET, POST,PATCH,DELETE")
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("x-origin", "api.captioner.com.ng")
		ctx.Next()
	})
	g.GET("/healthcheck", handler.Healthcheck)
	router.Connect(g, db)
	err := g.Run(*addr)
	if err != nil {
		log.Fatal(err, db)
	}
}
