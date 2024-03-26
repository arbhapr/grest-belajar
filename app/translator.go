package app

import (
	"grest.dev/grest"

	"grest-belajar/app/i18n"
)

// Translator returns a pointer to the translatorUtil instance (translator).
// If translator is not initialized, it creates a new translatorUtil instance, configures it, and assigns it to translator.
// It ensures that only one instance of translatorUtil is created and reused.
func Translator() *translatorUtil {
	if translator == nil {
		translator = &translatorUtil{}
		translator.configure()
	}
	return translator
}

// translator is a pointer to a translatorUtil instance.
// It is used to store and access the singleton instance of translatorUtil.
var translator *translatorUtil

// translatorUtil represents a translator utility.
// It embeds grest.translator, indicating that translatorUtil inherits from grest.Translator.
type translatorUtil struct {
	grest.Translator
}

// configure configures the translator utility instance.
func (t *translatorUtil) configure() {
	t.AddTranslation("en-US", i18n.EnUS())
	t.AddTranslation("id-ID", i18n.IdID())
}
