package middleware

import (
	helper "framework-auto-aprov-go/pkg/helpers"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//var jwtKey = []byte("my_secret_key")

func AuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if _, err := helper.ValidateToken(token); err != "" {
		log.Println("error:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

// func validateToken(token string) error {
// 	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
// 		}

// 		return []byte("secret"), nil
// 	})

// 	return err
// }
