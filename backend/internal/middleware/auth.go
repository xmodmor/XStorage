package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/xmodmor/XStorage/backend/internal/repository"
	"github.com/xmodmor/XStorage/backend/internal/response"
)

func JWTAuth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "missing or invalid authorization header")
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token claims")
			c.Abort()
			return
		}

		sub, _ := claims["sub"].(float64)
		c.Set("userID", uint(sub))

		c.Next()
	}
}

func APIKeyAuth(apiKeyRepo repository.APIKeyRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessKey := c.GetHeader("X-Access-Key")
		secretKey := c.GetHeader("X-Secret-Key")

		if accessKey == "" || secretKey == "" {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "missing api key headers")
			c.Abort()
			return
		}

		key, err := apiKeyRepo.FindByAccessKey(c.Request.Context(), accessKey)
		if err != nil || key.SecretKey != secretKey {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid api key")
			c.Abort()
			return
		}

		c.Set("appID", key.AppID)
		c.Set("permissions", key.Permissions)

		c.Next()
	}
}
