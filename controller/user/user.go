package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"im.v2/config"
	"im.v2/enum"
	"im.v2/utils"
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
		logrus.Info("invalid code")
		c.AbortWithStatusJSON(http.StatusBadRequest, &enum.ErrorGithubInvalidCOde)
		return
	}
	logrus.Debug("code:", code)

	var req struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
	}
	req.ClientID = config.Github().ClientID
	req.ClientSecret = config.Github().ClientSecret
	req.Code = code
	logrus.Debug("request body:", req)

	var resp struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	err := utils.HTTPPostJsonParse(config.Github().AccessTokenURL, req, resp)
	if err != nil {
		logrus.Warn("fetch access token from github error:", err)
		c.AbortWithStatusJSON(http.StatusBadGateway, &enum.ErrorGithubAccessToken)
		return
	}
	logrus.Debugf("response body: %#v", resp)
}
