package es_MX

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type es_MX struct {
	locale string
}

// New returns a new instance of translator for the 'es_MX' locale
func New() locales.Translator {
	return &es_MX{
		locale: "es_MX",
	}
}

// Locale returns the current translators string locale
func (l *es_MX) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *es_MX) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
