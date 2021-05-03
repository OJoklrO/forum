package middleware

import (
	"forum/pkg/app"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT(strict bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			if strict {
				app.ResponseError(c, http.StatusUnauthorized, "the token is empty")
				c.Abort()
			}
			return
		}

		claims, err := app.ParseToken(token)
		if err != nil {
			if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
				app.ResponseError(c, http.StatusUnauthorized, "token expired")
			} else {
				app.ResponseError(c, http.StatusUnauthorized, "token validation error")
			}
			c.Abort()
			return
		}

		c.Set("token", token)
		c.Set("user_name", claims.Name)
		c.Set("user_id", claims.Subject)
		c.Next()
	}
}
