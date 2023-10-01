package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	hand "github.com/thisjustinf/togo-list/internal/handlers"
	mid "github.com/thisjustinf/togo-list/internal/middleware"
)

func Run(orm *gorm.DB) {
	r := gin.Default()
	r.Use(mid.GinContextToContextMiddleware())
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.POST("/query", hand.GraphqlHandler(orm))
	r.GET("/", hand.PlaygroundHandler())
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
