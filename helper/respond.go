package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound = "record not found"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ok

// func Ok(c *gin.Context, data interface{}) {
// 	c.JSON(http.StatusOK, data)
// }

func OkWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func NoContent(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

// not ok

func BadRequest(c *gin.Context, message string, data ...interface{}) {
	obj := Response{
		Status:  http.StatusBadRequest,
		Message: message,
	}
	if len(data) > 0 {
		obj.Data = data[0]
	}
	c.JSON(http.StatusBadRequest, obj)
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Status:  http.StatusNotFound,
		Message: message,
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{
		Status:  http.StatusInternalServerError,
		Message: message,
	})
}

func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Response{
		Status:  http.StatusUnauthorized,
		Message: message,
	})
}
