package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Signin(c *gin.Context) {

}

func Signup(c *gin.Context) {

}

func Signout(c *gin.Context) {

}

func OAuth(c *gin.Context) {
	r := c.Request
	code := r.FormValue("code")
	if len(code) == 0 {
		logrus.Debug("invalid code")
		return
	}
}
