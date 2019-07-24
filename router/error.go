package router

import (
	"github.com/gin-gonic/gin"
	controllerError "im.v2/controller/error"
)

func error(e *gin.Engine)  {
	group := e.Group(prefixError)

	group.GET("/404", controllerError.NotFound)
}