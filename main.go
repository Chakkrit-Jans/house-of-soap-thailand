package main

import (
	"hst-api/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()
	db.Migrate()

	os.MkdirAll("uploads/products", 0755)
	r := gin.Default()
	serveRoutes(r)
	port := os.Getenv("PORT")
	if port != "" {
		port = "5000"
	}
	r.Run(":" + port)
}
