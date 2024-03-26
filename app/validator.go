package app

import (
	enLocale "github.com/go-playground/locales/en"
	idLocale "github.com/go-playground/locales/id"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	idTranslation "github.com/go-playground/validator/v10/translations/id"
	"grest.dev/grest"
)

// Validator returns a pointer to the validatorUtil instance (validator).
// If validator is not initialized, it creates a new validatorUtil instance, configures it, and assigns it to validator.
// It ensures that only one instance of validatorUtil is created and reused.
func Validator() *validatorUtil {
	if validator == nil {
		validator = &validatorUtil{}
		validator.configure()
	}
	return validator
}

// validator is a pointer to a validatorUtil instance.
// It is used to store and access the singleton instance of validatorUtil.
var validator *validatorUtil

// validatorUtil represents a validator utility.
// It embeds grest.Validator, indicating that validatorUtil inherits from grest.Validator.
type validatorUtil struct {
	grest.Validator
}

// configure configures the validator utility instance.
func (v *validatorUtil) configure() {
	v.New()
	v.RegisterCustomTypeFunc(v.ValidateValuer,
		NullBool{},
		NullInt64{},
		NullFloat64{},
		NullString{},
		NullDateTime{},
		NullDate{},
		NullTime{},
		NullText{},
		NullJSON{},
		NullUUID{},
	)
	v.RegisterTranslator("en", enLocale.New(), enTranslation.RegisterDefaultTranslations)
	v.RegisterTranslator("id", idLocale.New(), idTranslation.RegisterDefaultTranslations)
	v.RegisterTranslator("id-ID", idLocale.New(), idTranslation.RegisterDefaultTranslations)
}
