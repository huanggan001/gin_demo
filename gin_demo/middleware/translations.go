package middleware

import (
	"gin_demo/public"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// TranslationMiddleware 设置Translation
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//参照：https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go

		//设置支持语言
		en := en.New()
		zh := zh.New()

		//设置国际化翻译器
		uni := ut.New(zh, zh, en)
		val := validator.New()

		//根据参数取翻译器实例
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)

		//翻译器注册到validator
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(val, trans)
			en_translations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
			break
		default:
			zh_translations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			//自定义验证方法
			//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
			val.RegisterValidation("valid_username", func(fl validator.FieldLevel) bool {
				return fl.Field().String() == "admin"
			})
			val.RegisterValidation("valid_age", func(fl validator.FieldLevel) bool {
				return fl.Field().Int() > 20
			})
			val.RegisterValidation("valid_password", func(fl validator.FieldLevel) bool {
				return len(fl.Field().String()) >= 3
			})
			//val.RegisterValidation("valid_rule", func(fl validator.FieldLevel) bool {
			//	matched, _ := regexp.Match(`^\S+$`, []byte(fl.Field().String()))
			//	return matched
			//})

			//自定义验证器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			val.RegisterTranslation("valid_username", trans, func(ut ut.Translator) error {
				return ut.Add("valid_username", "{0} 填写不正确哦", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_username", fe.Field())
				return t
			})
			val.RegisterTranslation("valid_age", trans, func(ut ut.Translator) error {
				return ut.Add("valid_age", "{0} 不满20岁", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_age", fe.Field())
				return t
			})
			val.RegisterTranslation("valid_password", trans, func(ut ut.Translator) error {
				return ut.Add("valid_password", "{0} 密码长度不超过3位", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_password", fe.Field())
				return t
			})
			//val.RegisterTranslation("valid_rule", trans, func(ut ut.Translator) error {
			//	return ut.Add("valid_rule", "{0} 必须是非空字符", true)
			//}, func(ut ut.Translator, fe validator.FieldError) string {
			//	t, _ := ut.T("valid_rule", fe.Field())
			//	return t
			//})

			break
		}
		c.Set(public.TranslatorKey, trans)
		c.Set(public.ValidatorKey, val)
		c.Next()
	}
}
