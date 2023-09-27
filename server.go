package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/JustinTimeToCode/togo-list/graph/database"
	gen "github.com/JustinTimeToCode/togo-list/graph/generated"
	mid "github.com/JustinTimeToCode/togo-list/graph/middleware"
	res "github.com/JustinTimeToCode/togo-list/graph/resolvers"
)

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{Resolvers: &res.Resolver{}}))
	h.Use(extension.FixedComplexityLimit(5))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	database.InitDB()
	r := gin.Default()
	r.Use(mid.GinContextToContextMiddleware())
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
