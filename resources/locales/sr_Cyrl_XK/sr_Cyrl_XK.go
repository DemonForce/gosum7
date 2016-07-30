package sr_Cyrl_XK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sr_Cyrl_XK struct {
	locale string
}

// New returns a new instance of translator for the 'sr_Cyrl_XK' locale
func New() locales.Translator {
	return &sr_Cyrl_XK{
		locale: "sr_Cyrl_XK",
	}
}

// Locale returns the current translators string locale
func (l *sr_Cyrl_XK) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sr_Cyrl_XK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if (v == 0 && i%10 == 1 && i%100 != 11) || (f%10 == 1 && f%100 != 11) {
		return locales.PluralRuleOne, nil
	} else if (v == 0 && i%10 >= 2 && i%10 <= 4 && i%100 < 12 && i%100 > 14) || (f%10 >= 2 && f%10 <= 4 && f%100 < 12 && f%100 > 14) {
		return locales.PluralRuleFew, nil
	}

	return locales.PluralRuleOther, nil
}
