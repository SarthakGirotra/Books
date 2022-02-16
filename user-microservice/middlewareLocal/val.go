package middlewareLocal

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

type ApiError struct {
	// example: password
	Param string
	// example: password should contain 1 uppercase letter, 1 lowercase letter, 1 number, 1 special char and min 10 chars long
	Message string
}

//main validator
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		errs := translateError(err, cv.trans)
		return echo.NewHTTPError(http.StatusBadRequest, errs)
	}
	return nil
}

// password regex
func validatePass(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	match, _ := regexp.MatchString("^(.{0,9}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$", password)
	return !match
}

func translateError(err error, trans ut.Translator) []ApiError {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	out := make([]ApiError, len(validatorErrs))
	for i, e := range validatorErrs {
		if e.Field() == "Password" {
			s := "password should contain 1 uppercase letter, 1 lowercase letter, 1 number, 1 special char and min 10 chars long"
			out[i] = ApiError{e.Field(), s}
		} else {
			translatedErr := fmt.Errorf(e.Translate(trans))
			out[i] = ApiError{e.Field(), translatedErr.Error()}
		}

	}
	return out
}

// to validate json fields
func Val(e *echo.Echo) {
	val := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(val, trans)
	val.RegisterValidation("pass", validatePass)
	e.Validator = &CustomValidator{validator: val, trans: trans}
}
