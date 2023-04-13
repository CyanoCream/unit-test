package middleware

import (
	"fmt"
	"os"
	"sesi_8/helper"

	"github.com/gin-gonic/gin"
)

func TestMiddleware(c *gin.Context) {
	fmt.Println("Ini test middleware")

	c.Next()
}

func BasicAuth(c *gin.Context) {
	fmt.Println("Ini basic authentication")

	username, password, ok := c.Request.BasicAuth()
	if !ok {
		helper.Unauthorized(c, "something went wrong")
		c.Abort()
	}

	isValid := (username == os.Getenv("BASIC_USERNAME")) && (password == os.Getenv("BASIC_PASSWORD"))
	if !isValid {
		helper.Unauthorized(c, "wrong username or password")
		c.Abort()
	}
}
