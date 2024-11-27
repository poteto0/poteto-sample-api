package main

import (
	"net/http"

	"github.com/poteto0/poteto"
	"github.com/poteto0/poteto-sample-api/controller"
	"github.com/poteto0/poteto-sample-api/middles"
	"github.com/poteto0/poteto/middleware"
)

func main() {
	p := poteto.New()

	// logger
	logConfig := middleware.DefaultRequestLoggerConfig
	logConfig.LogHandleFunc = middles.MyHTTPLoggerFunc
	p.Register(middleware.RequestLoggerWithConfig(logConfig))

	// CORS
	p.Register(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		},
	))

	// User Router
	p.Leaf("/users", func(userApi poteto.Leaf) {
		userApi.Register(middleware.CamaraWithConfig(middleware.DefaultCamaraConfig))
		userApi.GET("/", controller.UserHandler)
		userApi.GET("/:name", controller.UserIdHandler)
	})

	p.GET("/users", controller.UserHandler)

	p.Run(":8080")
}

type User struct {
	Name any `json:"name"`
}
