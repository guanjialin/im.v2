package error

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFound(c *gin.Context)  {
	c.JSON(http.StatusNotFound, "<h1>404</h1>")
}