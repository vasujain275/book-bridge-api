package app

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/cors"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr            string
	maxOpenConns    int
	maxIdleConns    int
	maxConnLifetime string
}

func (app *application) mount() *gin.Engine {
	r := gin.New()

	// Middleware Setup
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Cors Middleware Setup
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           300 * time.Second,
	}))

	// Routes
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", app.healthCheckHandler)
	}

}
