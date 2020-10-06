package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var localToTrans = make(map[string]ut.Translator)
var defaultTrans ut.Translator

func ValidatorTransInit() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)
		// en
		{
			Trans, _ := uni.GetTranslator("en")
			_ = enTranslations.RegisterDefaultTranslations(v, Trans)
			localToTrans["en"] = Trans
			defaultTrans = Trans
		}
		// zh
		{
			Trans, _ := uni.GetTranslator("zh")
			_ = zhTranslations.RegisterDefaultTranslations(v, Trans)
			localToTrans["zh"] = Trans
		}
	}
}

func TranslateValidateErr(errs validator.ValidationErrors, context *gin.Context) validator.ValidationErrorsTranslations {
	local := context.GetHeader("Accept-Language")
	if trans, ok := localToTrans[local]; ok {
		return errs.Translate(trans)
	} else {
		return errs.Translate(defaultTrans)
	}
}
