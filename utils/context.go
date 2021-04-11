package utils

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func Handler(c *gin.Context, response interface{}) {
	err, ok := response.(error)
	if ok {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, response)
}
