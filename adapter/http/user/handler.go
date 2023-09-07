package user

import (
	"fmt"
	"go-crud/domain/user"
	"go-crud/pkg/exception"
	"go-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *user.Service
}

func (h *Handler) create(c *gin.Context) {
	dto := &user.CreateDTO{}
	err := c.BindJSON(dto)
	if err != nil {
		err = fmt.Errorf("bindJSON -> %w", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"code":    exception.ErrKeyInvalidRequest,
			"message": err.Error(),
		})
		return
	}

	createdUser, err := h.service.Create(c, dto)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, createdUser)
	return
}

func (h *Handler) list(c *gin.Context) {
	users, err := h.service.ListAll(c)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
	return
}
