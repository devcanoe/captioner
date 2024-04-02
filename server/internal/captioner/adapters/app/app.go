package app

import (
	"captioner.com.ng/config"
	"captioner.com.ng/internal/captioner/adapters/app/router"
	"captioner.com.ng/internal/captioner/drivers/database/mongodb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	App struct {
		Cfg *config.Config
	}
)

func New(app App) {
	g := gin.Default()
	db := mongodb.NewMongoStore(app.Cfg).Client
	v := validator.New()
	r := router.NewRouterGroup(g, db, v)
	{
		r.AuthRouter()
		r.Merouter()
	}

	g.Run(app.Cfg.Port)

}
