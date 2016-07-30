package fr_BE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_BE struct {
	locale string
}

// New returns a new instance of translator for the 'fr_BE' locale
func New() locales.Translator {
	return &fr_BE{
		locale: "fr_BE",
	}
}

// Locale returns the current translators string locale
func (l *fr_BE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *fr_BE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
