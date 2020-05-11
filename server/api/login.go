package api

import (
	"gin-admin/config"
	"gin-admin/initialize"
	"gin-admin/middleware"
	"gin-admin/model"
	"gin-admin/service"
	"gin-admin/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	var user2 model.User
	var err error
	err = c.ShouldBindJSON(&user)
	err = initialize.Validate.Struct(user)
	if err != nil {
		initialize.HandleError(err, c)
		return
	}
	user2, err = service.GetUserByName(user.Username)

	if user2.Password != utils.MD5([]byte(user.Password)) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "用户名或密码不对",
		})
		return
	}
	token := CreateToken(c, &user2)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": token,
	})
}

func CreateToken(c *gin.Context, user *model.User) string {
	//自定义claim
	//claim := jwt.MapClaims{
	//	"nickname": user.NickName,
	//	"username": user.Username,
	//	"nbf":      time.Now().Unix(),
	//	"iat":      time.Now().Unix(),
	//}
	claim := middleware.UserClaim{
		Id:       user.Id,
		NickName: user.NickName,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString([]byte(config.ApplicationConfig.JwtSecret))
	return tokenString
}
