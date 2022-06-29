package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/takadev15/mygram-api/utils"
)

var secret = "abcdefghijklmnopq"

func GenerateToken(id uint, email string) (string, error) {
  token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)

  claims["id"] = id
  claims["email"] = email
  claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

  signedToken, err := token.SignedString([]byte(secret)) 
  if err != nil {
    return "", err
  }
  return signedToken, nil
}

func Authentication() gin.HandlerFunc {
  return func(c *gin.Context) {
    verifiedToken, err := utils.VerifyToken(c)
    if err != nil {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
        "err" : "unauthorized",
        "message" : err.Error(),
      })
    }
    data := verifiedToken.(jwt.MapClaims)
    c.Set("id", data["id"])
    c.Set("email", data["email"])
    c.Set("user_data", verifiedToken)
    c.Next()
  }
}
