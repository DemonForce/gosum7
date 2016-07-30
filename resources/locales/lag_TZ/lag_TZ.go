package lag_TZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lag_TZ struct {
	locale string
}

// New returns a new instance of translator for the 'lag_TZ' locale
func New() locales.Translator {
	return &lag_TZ{
		locale: "lag_TZ",
	}
}

// Locale returns the current translators string locale
func (l *lag_TZ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lag_TZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 0 {
		return locales.PluralRuleZero, nil
	} else if (i == 0 || i == 1) && n != 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
