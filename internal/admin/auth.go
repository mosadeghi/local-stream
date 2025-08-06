package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mosadeghi/local-stream/internal/config"
)

func BasicAuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, hasAuth := c.Request.BasicAuth()
		if !hasAuth || user != cfg.AdminUsername || pass != cfg.AdminPassword {
			c.Header("WWW-Authenticate", `Basic realm="Admin Panel"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
