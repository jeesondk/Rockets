package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"log"
	"net/http"
)

func main() {
	var serverPort int
	var debugMode bool

	flag.IntVar(&serverPort, "port", 8080, "Specify RocketService port, Default is 8080")
	flag.BoolVar(&debugMode, "debug", false, "Specify if Gin is running in \"debug\" or \"release\"(false) mode, Default is \"release\"(false)")

	flag.Parse()

	hostBinding := fmt.Sprintf("localhost:%d", serverPort)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "RocketService API"
	docs.SwaggerInfo.Description = "Rocket message lister service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = hostBinding
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	if !debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	//c := controllers.NewController()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	/*
		r.GET("/counter/increment", c.Increment)
		r.GET("/counter/decrement", c.Decrement)
		r.GET("/counter/atomicincrement", c.AtomicIncrement)
		r.GET("/counter/atomicdecrement", c.AtomicDecrement)
		r.GET("/counter/currentnumber", c.CurrentNumber)
		r.GET("/counter/reset", c.Reset)
	*/
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(r.Run(fmt.Sprintf(":%d", serverPort)))

}
