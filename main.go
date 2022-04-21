package main

import (
	"RocketService/controllers"
	"RocketService/docs"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var serverPort int
	var debugMode bool

	flag.IntVar(&serverPort, "port", 8088, "Specify RocketService port, Default is 8080")
	flag.BoolVar(&debugMode, "debug", false, "Specify if Gin is running in \"debug\" or \"release\"(false) mode, Default is \"release\"(false)")

	flag.Parse()

	hostBinding := fmt.Sprintf("localhost:%d", serverPort)
	hostPort := fmt.Sprintf(":%d", serverPort)

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

	srv := &http.Server{
		Addr:    hostPort,
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
