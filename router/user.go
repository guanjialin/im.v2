package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"im.v2/controller/user"
)

func login(e *gin.Engine) {
	group := e.Group(prefixUser)

	group.Handle(http.MethodPost, "/signin", user.Signin)
	group.Handle(http.MethodPost, "/signup", user.Signup)
	group.Handle(http.MethodDelete, "/signout", user.Signout)
	group.Handle(http.MethodGet, "/oauth/redirect", user.OAuth)
}
