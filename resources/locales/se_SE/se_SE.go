package se_SE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type se_SE struct {
	locale string
}

// New returns a new instance of translator for the 'se_SE' locale
func New() locales.Translator {
	return &se_SE{
		locale: "se_SE",
	}
}

// Locale returns the current translators string locale
func (l *se_SE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *se_SE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	} else if n == 2 {
		return locales.PluralRuleTwo, nil
	}

	return locales.PluralRuleOther, nil
}
