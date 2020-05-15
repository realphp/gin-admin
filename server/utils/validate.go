package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Validate *validator.Validate
	Trans    ut.Translator
)

func Initvalidate() {
	zh := zh.New()
	uni := ut.New(zh, zh)
	Trans, _ = uni.GetTranslator("zh")
	Validate = validator.New()
	zh_translations.RegisterDefaultTranslations(Validate, Trans)
}

func HandleError(err error, c *gin.Context) {
	if err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  errs[0].Translate(Trans),
		})
	}
	return
}
