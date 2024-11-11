package controllers

import (
	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func InitHone() *HomeController {
	return &HomeController{}
}

func (hc *HomeController) Index(c *gin.Context) {
	c.File("views/home/index.html")
}
