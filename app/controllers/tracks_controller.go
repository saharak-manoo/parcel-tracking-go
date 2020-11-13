package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/saharak/parcel-tracking-go/app/models"
)

func (server *Server) KerryExpress(c *gin.Context) {
	trackingNumber := c.Query("trackingNumber")
	etracking := models.ETracking{}
	data, err := etracking.Track("kerry_express", trackingNumber)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data": data,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
}

func (server *Server) JTExpress(c *gin.Context) {
	trackingNumber := c.Query("trackingNumber")
	etracking := models.ETracking{}
	data, err := etracking.Track("jt_express", trackingNumber)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data": data,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
}

func (server *Server) FlashExpress(c *gin.Context) {
	trackingNumber := c.Query("trackingNumber")
	etracking := models.ETracking{}
	data, err := etracking.Track("flash_express", trackingNumber)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data": data,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
}