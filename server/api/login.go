package api

import (
	"gin-admin/config"
	"gin-admin/middleware"
	"gin-admin/model"
	"gin-admin/service"
	"gin-admin/utils"
	"gin-admin/utils/response"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// @Tags User
// @Summary 用户登录
// accept	json
// produce	json
// @Param   user_name      query    string     true        "用户名"
// @Param   password      query    string     true        "密码"
// @Success 200 {string} string	"ok"
// @Router /login/ [post]
func Login(c *gin.Context) {
	var user model.User
	var user2 model.User
	var err error
	err = c.ShouldBindJSON(&user)
	utils.Initvalidate()
	err = utils.Validate.Struct(user)
	if err != nil {
		utils.HandleError(err, c)
		return
	}
	user2, err = service.GetUserByName(user.UserName)

	if user2.Password != utils.MD5([]byte(user.Password)) {
		response.FailMessage("用户名或密码不对", c)
		return
	}
	token, ExpiresAt := CreateToken(c, &user2)
	response.OkData(gin.H{
		"token":      token,
		"user":       user2,
		"expires_at": ExpiresAt,
	}, c)

}

func CreateToken(c *gin.Context, user *model.User) (string, int64) {
	//自定义claim
	claim := middleware.UserClaim{
		Id:       user.Id,
		NickName: user.NickName,
		Username: user.UserName,
		RoleId:   user.RoleId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString([]byte(config.ApplicationConfig.JwtSecret))
	return tokenString, claim.StandardClaims.ExpiresAt
}
