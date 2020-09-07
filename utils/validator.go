package utils

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func InitValiadtor() {
	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	V, _ := binding.Validator.Engine().(*validator.Validate)
	V.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	zh_translations.RegisterDefaultTranslations(V, trans)
}

func Translate(errs interface{}) (str string) {
	switch t := errs.(type) {
	case validator.ValidationErrors:
		for _, e := range t {
			str = e.Translate(trans)
			return
		}
	case error:
		str = t.Error()
		return
	default:
		str = "unknown mistake"
	}
	return
}
