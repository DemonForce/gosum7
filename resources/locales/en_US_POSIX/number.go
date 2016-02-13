package en_US_POSIX

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "0/00"}
	formats = ut.NumberFormats{Decimal: "#0.######", Currency: "¤\u00a0#0.00", Percent: "#0%"}
)
