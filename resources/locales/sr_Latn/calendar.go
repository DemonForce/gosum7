package sr_Latn

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd.MM.y.", Short: "d.M.yy."},
			Time:     ut.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "maj", Jun: "jun", Jul: "jul", Aug: "avg", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "dec"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "j", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "j", Jul: "j", Aug: "a", Sep: "s", Oct: "o", Nov: "n", Dec: "d"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "mart", Apr: "april", May: "maj", Jun: "jun", Jul: "jul", Aug: "avgust", Sep: "septembar", Oct: "oktobar", Nov: "novembar", Dec: "decembar"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ned", Mon: "pon", Tue: "uto", Wed: "sre", Thu: "čet", Fri: "pet", Sat: "sub"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "u", Wed: "s", Thu: "č", Fri: "p", Sat: "s"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ne", Mon: "po", Tue: "ut", Wed: "sr", Thu: "če", Fri: "pe", Sat: "su"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "nedelja", Mon: "ponedeljak", Tue: "utorak", Wed: "sreda", Thu: "četvrtak", Fri: "petak", Sat: "subota"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "pre podne", PM: "po podne"},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "pre podne", PM: "po podne"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "pre podne", PM: "po podne"},
			},
		},
	}
}
