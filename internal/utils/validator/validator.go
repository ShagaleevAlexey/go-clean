package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"

	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate = validator.New()
	trans    ut.Translator
)

func SetupValidator() {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	//validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
	//	return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("required", fe.Field())
	//
	//	return t
	//})
}

func TranslateError(errs validator.ValidationErrors) string {
	translates := errs.Translate(trans)

	return translates[reflect.ValueOf(translates).MapKeys()[0].String()]
}

func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}
	//for _, err := range err.(validator.ValidationErrors) {
	//	fmt.Println(err.Namespace())
	//	fmt.Println(err.Field())
	//	fmt.Println(err.StructNamespace())
	//	fmt.Println(err.StructField())
	//	fmt.Println(err.Tag())
	//	fmt.Println(err.ActualTag())
	//	fmt.Println(err.Kind())
	//	fmt.Println(err.Type())
	//	fmt.Println(err.Value())
	//	fmt.Println(err.Param())
	//	fmt.Println()
	//}

	return err
}