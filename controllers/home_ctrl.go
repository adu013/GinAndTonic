package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	// Pass data directly into the HTML renderer
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Title":   "Gin & Tonic",
		"Message": "Welcome to your beautifully organized Gin & Tonic homepage!",
		"Status":  "Online",
	})
}
