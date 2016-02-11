package sl

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, dd. MMMM y", Long: "dd. MMMM y", Medium: "d. MMM y", Short: "d. MM. yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "maj", Jun: "jun", Jul: "jul", Aug: "avg", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "dec"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "j", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "j", Jul: "j", Aug: "a", Sep: "s", Oct: "o", Nov: "n", Dec: "d"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "marec", Apr: "april", May: "maj", Jun: "junij", Jul: "julij", Aug: "avgust", Sep: "september", Oct: "oktober", Nov: "november", Dec: "december"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "ned", Mon: "pon", Tue: "tor", Wed: "sre", Thu: "čet", Fri: "pet", Sat: "sob"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "t", Wed: "s", Thu: "č", Fri: "p", Sat: "s"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "ned.", Mon: "pon.", Tue: "tor.", Wed: "sre.", Thu: "čet.", Fri: "pet.", Sat: "sob."},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "nedelja", Mon: "ponedeljek", Tue: "torek", Wed: "sreda", Thu: "četrtek", Fri: "petek", Sat: "sobota"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "dop.", PM: "pop."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "dop.", PM: "pop."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "dopoldne", PM: "popoldne"},
			},
		},
	}
}
