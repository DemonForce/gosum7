package teo_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type teo_KE struct {
	locale string
}

// New returns a new instance of translator for the 'teo_KE' locale
func New() locales.Translator {
	return &teo_KE{
		locale: "teo_KE",
	}
}

// Locale returns the current translators string locale
func (l *teo_KE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *teo_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
