package haw_US

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type haw_US struct {
	locale string
}

// New returns a new instance of translator for the 'haw_US' locale
func New() locales.Translator {
	return &haw_US{
		locale: "haw_US",
	}
}

// Locale returns the current translators string locale
func (l *haw_US) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *haw_US) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
