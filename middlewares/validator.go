package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"gitlab.zmwk.cn/open/open-golang-tools/logs"
	"go.uber.org/zap"
	"reflect"
	"strings"
	"zhima_chat_ai/internal/app/zhima_chat_ai/define"
)

var (
	uni *ut.UniversalTranslator
	//trans ut.Translator
)

func Init() {
	//register editor
	zhT := zh.New()
	enT := en.New()
	uni = ut.New(zhT, enT)
}
func switchTrans(lang string) ut.Translator {
	var err error
	trans, _ := uni.GetTranslator(lang)

	//get gin validator
	validate := binding.Validator.Engine().(*validator.Validate)
	// check trans
	switch lang {
	case "en":
		err = en_translations.RegisterDefaultTranslations(validate, trans)
	case define.LangZhCn:
		err = zh_translations.RegisterDefaultTranslations(validate, trans)
	default:
		err = zh_translations.RegisterDefaultTranslations(validate, trans)
	}
	if err != nil {
		logs.Error("validator register error", zap.Error(err))
	}
	return trans
}

func GetValidateErr(obj any, rawErr error, lang string) error {
	validationErrs, ok := rawErr.(validator.ValidationErrors)
	if !ok {
		return rawErr
	}
	trans := switchTrans(lang)
	var errString []string
	for _, validationErr := range validationErrs {
		field, ok := reflect.TypeOf(obj).FieldByName(validationErr.Field())
		if ok {
			if e := field.Tag.Get("msg"); e != "" {
				errString = append(errString, fmt.Sprintf("%s: %s", field.Name, e))
				continue
			} else {
				errString = append(errString, fmt.Sprintf("%s", validationErr.Translate(trans)))
			}
		} else {
			errString = append(errString, validationErr.Error())
		}
	}
	return errors.New(strings.Join(errString, ";"))
}
