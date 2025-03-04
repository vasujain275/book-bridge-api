package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vasujain275/book-bridge-api/internal/app"
)

func (app *application) healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
