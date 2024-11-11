package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct{}

func InitHone() *HomeController {
	return &HomeController{}
}

func (hc *HomeController) Home(c *gin.Context) {
	c.String(http.StatusOK, "HOME INDEX")
}
