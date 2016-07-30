package so_DJ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type so_DJ struct {
	locale string
}

// New returns a new instance of translator for the 'so_DJ' locale
func New() locales.Translator {
	return &so_DJ{
		locale: "so_DJ",
	}
}

// Locale returns the current translators string locale
func (l *so_DJ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *so_DJ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
