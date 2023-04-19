package middleware

import (
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/go-chatgpt-api/api"
	"github.com/linweiyuan/go-chatgpt-api/util/logger"
)

//goland:noinspection GoUnhandledErrorResult
func HeaderCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader(api.AuthorizationHeader) == "" {
			logger.Info("Missing access token")
			c.AbortWithStatusJSON(http.StatusForbidden, api.ReturnMessage("Missing accessToken."))
			return
		}
		c.Next()
	}
}

func CheckClientIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ids := c.Query("ids")
		if ids == "" {
			logger.Info("Missing ids parameter")
			c.AbortWithStatusJSON(http.StatusBadRequest, api.ReturnMessage("Missing ids parameter."))
			return
		}

		if ids != os.Getenv("CHAT_CLIENT_ID") {
			logger.Info("Invalid ids parameter!")
			c.AbortWithStatusJSON(http.StatusForbidden, api.ReturnMessage("Invalid ids parameter!"))
		}
	}
}
