package http

import (
	"fmt"
	"gotest/internal/core/port/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationType       = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(token service.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(authorizationHeaderKey)

		isEmpty := len(authHeader) == 0

		fmt.Println(authHeader)

		if isEmpty {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})

			return

		}

		fields := strings.Fields(authHeader)
		isValid := len(fields) == 2

		if !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})

			return
		}

		currentAuthorizationType := strings.ToLower(fields[0])

		if currentAuthorizationType != authorizationType {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})

			return
		}

		accesToken := fields[1]
		payload, err := token.ValidateToken(accesToken)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})

			return
		}

		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}
