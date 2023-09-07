package utils

import (
	"errors"
	"go-crud/pkg/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var appErr *exception.Error
	if errors.As(err, &appErr) {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"code":    appErr.Code(),
			"message": appErr.Error(),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
		"code":    exception.ErrKeyInternalServerError,
		"message": err.Error(),
	})
	return
}
