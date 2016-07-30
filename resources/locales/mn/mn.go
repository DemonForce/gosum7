package mn

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mn struct {
	locale string
}

// New returns a new instance of translator for the 'mn' locale
func New() locales.Translator {
	return &mn{
		locale: "mn",
	}
}

// Locale returns the current translators string locale
func (l *mn) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mn) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
