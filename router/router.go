package router

import (
	"github.com/gin-gonic/gin"
)

const (
	prefixUser = "/user"
	prefixError = "/error"
)

func Register(e *gin.Engine)  {
	login(e)
}