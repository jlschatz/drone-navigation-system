package main

import (
	"os"

	"github.com/drone-navigation-system/controller"
	router "github.com/drone-navigation-system/http"
	"github.com/drone-navigation-system/service"
)

func main() {

	httpRouter := router.NewChiRouter()
	srv := service.NewDNSService()
	c := controller.NewDNSController(srv)

	baseURL := os.Getenv("BASE_URL")

	httpRouter.GET("/", c.Swagger)

	httpRouter.GET("/swagger", c.Swagger)

	httpRouter.GET("/ping", c.Ping)

	httpRouter.POST(baseURL+"/loc", c.GetLocation)

	httpRouter.SERVE(":8080")

}
