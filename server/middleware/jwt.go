package middleware

import (
	"errors"
	"gin-admin/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserClaim struct {
	Username           string `json:"username"`
	NickName           string `json:"NickName"`
	jwt.StandardClaims        //嵌套了这个结构体就实现了Claim接口
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "未登录或非法访问",
			})
			c.Abort()
			return
		}
		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}

}

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

func ParseToken(tokenString string) (claim jwt.Claims, err error) {
	token, err := jwt.Parse(tokenString, secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	claim = token.Claims
	if token.Valid {
		return claim, nil
	}
	return nil, TokenInvalid
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ApplicationConfig.JwtSecret), nil
	}
}
