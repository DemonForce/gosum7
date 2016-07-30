package ga

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ga struct {
	locale string
}

// New returns a new instance of translator for the 'ga' locale
func New() locales.Translator {
	return &ga{
		locale: "ga",
	}
}

// Locale returns the current translators string locale
func (l *ga) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ga) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	} else if n == 2 {
		return locales.PluralRuleTwo, nil
	} else if n >= 3 && n <= 6 {
		return locales.PluralRuleFew, nil
	} else if n >= 7 && n <= 10 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
