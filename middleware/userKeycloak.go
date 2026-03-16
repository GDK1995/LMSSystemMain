package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	issuer := os.Getenv("KEYCLOAK_URL")
	provider, err := oidc.NewProvider(context.Background(), issuer)
	if err != nil {
		panic("Failed to connect to Keycloak: " + err.Error())
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID: os.Getenv("KEYCLOAK_CLIENT_ID"),
	})

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		idToken, errTwo := verifier.Verify(c.Request.Context(), tokenStr)
		if errTwo != nil {
			fmt.Printf("Verification failed: %v", errTwo)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		var claims map[string]interface{}
		if errThree := idToken.Claims(&claims); errThree != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims"})
			return
		}

		c.Set("userID", claims["sub"])
		c.Next()
	}
}
