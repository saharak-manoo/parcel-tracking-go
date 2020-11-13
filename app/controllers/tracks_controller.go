package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) KerryExpress(c *gin.Context) {
	errList = map[string]string{}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"data": nil,
	})
}

func (server *Server) JTExpress(c *gin.Context) {
	errList = map[string]string{}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"data": nil,
	})
}

func (server *Server) FlashExpress(c *gin.Context) {
	errList = map[string]string{}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"data": nil,
	})
}