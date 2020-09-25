package router

import (
	"github.com/gin-gonic/gin"
	"go_gin_api/service"
)

func Route(router *gin.Engine) *gin.Engine {
	router.GET("/", service.HelloWorld)
	router.GET("/people", service.GetPeople)
	router.GET("/people/:id", service.GetPerson)
	router.POST("/people", service.CreatePerson)
	router.PUT("/people/:id", service.UpdatePerson)
	router.DELETE("/people/:id", service.DeletePerson)

	return router
}
