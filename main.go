package main

import (
	"RocketService/controllers"
	"RocketService/docs"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	var serverPort int
	var debugMode bool

	flag.IntVar(&serverPort, "port", 8088, "Specify RocketService port, Default is 8080")
	flag.BoolVar(&debugMode, "debug", false, "Specify if Gin is running in \"debug\" or \"release\"(false) mode, Default is \"release\"(false)")

	flag.Parse()

	hostBinding := fmt.Sprintf("localhost:%d", serverPort)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "RocketService API"
	docs.SwaggerInfo.Description = "Rocket message lister service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = hostBinding
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	if !debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	c := controllers.NewController()

	r.POST("messages", c.ReceiveMessage)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(fmt.Sprintf(":%d", serverPort))

}
