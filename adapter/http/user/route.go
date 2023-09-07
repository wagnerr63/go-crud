package user

import (
	"go-crud/domain/user"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup, service *user.Service) {
	handler := Handler{
		service: service,
	}

	router.POST("/users", handler.create)
	router.GET("/users", handler.list)
}
