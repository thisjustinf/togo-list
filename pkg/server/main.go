package server

import (
	"fmt"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	hand "github.com/thisjustinf/togo-list/internal/handlers"
	mid "github.com/thisjustinf/togo-list/internal/middleware"
)

func Run(orm *gorm.DB) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("")
	}
	r := gin.New()
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(mid.GinContextToContextMiddleware())
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.POST("/query", hand.GraphqlHandler(orm))
	r.GET("/", hand.PlaygroundHandler())
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
