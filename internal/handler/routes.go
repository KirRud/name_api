package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(handler handlerInterface) *gin.Engine {
	router := gin.Default()
	router.GET("/info/", handler.GetInfo)
	router.DELETE("/person/:id", handler.DeletePerson)
	router.PUT("/person/:id", handler.PutPerson)
	router.POST("/perosn", handler.PostPerson)
	return router
}
