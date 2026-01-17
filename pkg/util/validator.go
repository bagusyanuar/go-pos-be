package util

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strings"

	"github.com/bagusyanuar/go-pos-be/pkg/constant"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var trans ut.Translator

func GetTranslator() ut.Translator {
	return trans
}

func Validate(v *validator.Validate, req any) (map[string][]string, error) {
	err := v.Struct(req)
	messages := make(map[string][]string)
	if err != nil {
		translator := GetTranslator()

		val := reflect.ValueOf(req)
		var structName string
		if val.Kind() == reflect.Ptr {
			structName = val.Elem().Type().Name()
		} else {
			structName = val.Type().Name()
		}

		for _, e := range err.(validator.ValidationErrors) {
			fullNameSpace := e.Namespace()
			customKey := strings.TrimPrefix(fullNameSpace, structName+".")
			translated := strings.ToLower(e.Translate(translator))
			messages[customKey] = append(messages[customKey], translated)
		}
	}
	return messages, err
}

func RegisterValidatorTag(v *validator.Validate) {
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		tag := field.Tag.Get("json")
		if tag == "-" || tag == "" {
			return field.Name
		}

		name := strings.Split(tag, ",")[0]

		if name == "" {
			return field.Name
		}

		return name
	})
}

func RegisterValidatorRule(v *validator.Validate) {

	/* === register symbol validation === */
	if err := symbolValidationRule(v); err != nil {
		panic("failed to register default EN translations: " + err.Error())
	}

	/* === register array validation === */
	if err := arrayValidationRule(v); err != nil {
		panic("failed to register array validation: " + err.Error())
	}

	/* === register contact type validation === */
	if err := contactValidationRule(v); err != nil {
		panic("failed to register contact type validation: " + err.Error())
	}

}

func RegisterValidatorTranslation(v *validator.Validate) {
	locale := en.New()
	uni := ut.New(locale, locale)
	t, _ := uni.GetTranslator("en")
	trans = t

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		panic("failed to register default EN translations: " + err.Error())
	}

	if err := symbolTranslation(v, trans); err != nil {
		panic("failed to register symbol translation: " + err.Error())
	}

	if err := arrayTranslation(v, trans); err != nil {
		panic("failed to register array translation: " + err.Error())
	}

	if err := contactTypeTranslation(v, trans); err != nil {
		panic("failed to register contact type translation: " + err.Error())
	}
}

func symbolValidationRule(v *validator.Validate) error {
	return v.RegisterValidation("symbol", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()

		symbolRegex := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]{};':"\\|,.<>\/?]`)
		return symbolRegex.MatchString(value)
	})
}

func arrayValidationRule(v *validator.Validate) error {
	return v.RegisterValidation("array", func(fl validator.FieldLevel) bool {
		raw, ok := fl.Field().Interface().(json.RawMessage)
		if !ok {
			return false
		}

		var tmp any
		if err := json.Unmarshal(raw, &tmp); err != nil {
			return false
		}

		_, isArray := tmp.([]any) // cek apakah hasil decode adalah array JSON
		return isArray
	})
}

var ValidContactTypes = []constant.ContactType{
	constant.Phone,
	constant.Whatsapp,
	constant.Email,
	constant.Telegram,
	constant.OnlineShop,
	constant.Instagram,
	constant.Facebook,
	constant.Tiktok,
	constant.Youtube,
}

func contactValidationRule(v *validator.Validate) error {
	return v.RegisterValidation("contact_type", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		for _, t := range ValidContactTypes {
			if string(t) == value {
				return true
			}
		}
		return false
	})
}

func symbolTranslation(v *validator.Validate, trans ut.Translator) error {
	return v.RegisterTranslation("symbol", trans,
		func(ut ut.Translator) error {
			return ut.Add("symbol", "{0} must contain at least one symbol", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("symbol", fe.Field())
			return t
		},
	)
}

func arrayTranslation(v *validator.Validate, trans ut.Translator) error {
	return v.RegisterTranslation("array", trans,
		func(ut ut.Translator) error {
			return ut.Add("array", "{0} must be an array", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("array", fe.Field())
			return t
		},
	)
}

func contactTypeTranslation(v *validator.Validate, trans ut.Translator) error {
	return v.RegisterTranslation("contact_type", trans,
		func(ut ut.Translator) error {
			var types []string
			for _, t := range ValidContactTypes {
				types = append(types, string(t))
			}

			joinedEnums := strings.Join(types, ", ")
			msg := "{0} must be one of: " + joinedEnums

			return ut.Add("contact_type", msg, true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("contact_type", fe.Field())
			return t
		},
	)
}
