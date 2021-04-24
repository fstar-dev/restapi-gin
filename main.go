package main

import (
	"log"
	"os"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	config "github.com/restuwahyu13/gin-rest-api/configs"
	route "github.com/restuwahyu13/gin-rest-api/routes"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

func main() {
	/**
	@description Setup Database Connection
	*/
	db := config.Connection()
	/**
	@description Setup Router
	*/
	router := gin.Default()
	/**
	@description Setup Mode Application
	*/
	gin.SetMode(gin.ReleaseMode)
	/**
	@description Setup Middleware
	*/
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowCredentials: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))
	/**
	@description Init All Route
	*/
	route.InitAuthRoutes(db, router)
	route.InitStudentRoutes(db, router)
	/**
	@description Setup Server
	*/
	port := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		port <- util.GodotEnv("GO_PORT")
	} else {
		port <- os.Getenv("GO_PORT")
	}

	log.Fatal(router.Run(":" + <-port))
}
