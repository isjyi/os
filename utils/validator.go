package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

const (
	phoneRegexString = "^1[3|4|5|7|8|9][0-9]\\d{8}$"
)

func checkPhone(f validator.FieldLevel) bool {
	reg := regexp.MustCompile(phoneRegexString)
	return reg.MatchString(f.Field().String())
}

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	//修改gin框架中的Validator属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// v.RegisterStructValidation(SignUpParamStructLevelValidation, SignUpParam{})

		if err := v.RegisterValidation("phone", checkPhone); err != nil {
			return err
		}

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}

		// 添加额外翻译
		// 注意！因为要使用到trans实例
		// 所以这一步注册要放到trans初始化的后面
		if err := v.RegisterTranslation(
			"phone",
			trans,
			registerTranslator("phone", "{0}非法"),
			translate,
		); err != nil {
			return err
		}

		// 指定TagName 字段

		// v.RegisterTagNameFunc(func(field reflect.StructField) string {
		// 	label := field.Tag.Get("label")
		// 	if label == "" {
		// 		return field.Name
		// 	}
		// 	return label
		// })

		return
	}
	return
}

// SignUpParamStructLevelValidation 自定义SignUpParam结构体校验函数
// func SignUpParamStructLevelValidation(sl validator.StructLevel) {
// 	su := sl.Current().Interface().(SignUpParam)

// 	if su.Password != su.RePassword {
// 		// 输出错误提示信息，最后一个参数就是传递的param
// 		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
// 	}
// }

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

// 去掉结构体名称前缀
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.LastIndex(field, ".")+1:]] = err[strings.LastIndexAny(err, field[strings.LastIndex(field, ".")+1:])+1:]
	}
	return res
}

//handler中调用的错误翻译方法
func T(err validator.ValidationErrors) interface{} {
	return removeTopStruct(err.Translate(trans)) // 翻译校验错误提示
}
