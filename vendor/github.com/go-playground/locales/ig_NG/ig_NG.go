package ig_NG

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ig_NG struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'ig_NG' locale
func New() locales.Translator {
	return &ig_NG{
		locale:                 "ig_NG",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Jen", "Feb", "Maa", "Epr", "Mee", "Juu", "Jul", "Ọgọ", "Sep", "Ọkt", "Nov", "Dis"},
		monthsWide:             []string{"", "Jenụwarị", "Febrụwarị", "Maachị", "Eprel", "Mee", "Juun", "Julaị", "Ọgọọst", "Septemba", "Ọktoba", "Novemba", "Disemba"},
		daysAbbreviated:        []string{"Ụka", "Mọn", "Tiu", "Wen", "Tọọ", "Fraị", "Sat"},
		daysWide:               []string{"Mbọsị Ụka", "Mọnde", "Tiuzdee", "Wenezdee", "Tọọzdee", "Fraịdee", "Satọdee"},
		periodsAbbreviated:     []string{"A.M.", "P.M."},
		periodsWide:            []string{"A.M.", "P.M."},
		erasAbbreviated:        []string{"T.K.", "A.K."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Tupu Kristi", "Afọ Kristi"},
		timezones:              map[string]string{"EDT": "EDT", "BT": "BT", "AWDT": "AWDT", "HKT": "HKT", "MEZ": "MEZ", "COT": "COT", "COST": "COST", "ART": "ART", "AKDT": "AKDT", "NZST": "NZST", "CHADT": "CHADT", "ACST": "ACST", "WAT": "WAT", "TMST": "TMST", "CDT": "CDT", "NZDT": "NZDT", "UYT": "UYT", "OESZ": "OESZ", "MESZ": "MESZ", "WARST": "WARST", "PST": "PST", "WITA": "WITA", "MYT": "MYT", "AWST": "AWST", "JDT": "JDT", "WESZ": "WESZ", "SGT": "SGT", "PDT": "PDT", "AST": "AST", "IST": "IST", "GYT": "GYT", "MDT": "MDT", "BOT": "BOT", "ACWDT": "ACWDT", "WAST": "WAST", "EST": "EST", "CAT": "CAT", "HAT": "HAT", "GMT": "GMT", "LHST": "LHST", "ACWST": "ACWST", "JST": "JST", "∅∅∅": "∅∅∅", "HADT": "HADT", "WART": "WART", "TMT": "TMT", "OEZ": "OEZ", "WEZ": "WEZ", "ECT": "ECT", "AEST": "AEST", "EAT": "EAT", "CST": "CST", "CLST": "CLST", "WIT": "WIT", "UYST": "UYST", "ADT": "ADT", "AKST": "AKST", "CLT": "CLT", "SRT": "SRT", "ACDT": "ACDT", "VET": "VET", "SAST": "SAST", "ChST": "ChST", "LHDT": "LHDT", "HKST": "HKST", "GFT": "GFT", "HAST": "HAST", "AEDT": "AEDT", "ARST": "ARST", "CHAST": "CHAST", "WIB": "WIB", "MST": "MST", "HNT": "HNT"},
	}
}

// Locale returns the current translators string locale
func (ig *ig_NG) Locale() string {
	return ig.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ig_NG'
func (ig *ig_NG) PluralsCardinal() []locales.PluralRule {
	return ig.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ig_NG'
func (ig *ig_NG) PluralsOrdinal() []locales.PluralRule {
	return ig.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ig_NG'
func (ig *ig_NG) PluralsRange() []locales.PluralRule {
	return ig.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ig_NG'
func (ig *ig_NG) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ig_NG'
func (ig *ig_NG) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ig_NG'
func (ig *ig_NG) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ig *ig_NG) MonthAbbreviated(month time.Month) string {
	return ig.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ig *ig_NG) MonthsAbbreviated() []string {
	return ig.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ig *ig_NG) MonthNarrow(month time.Month) string {
	return ig.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ig *ig_NG) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (ig *ig_NG) MonthWide(month time.Month) string {
	return ig.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ig *ig_NG) MonthsWide() []string {
	return ig.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ig *ig_NG) WeekdayAbbreviated(weekday time.Weekday) string {
	return ig.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ig *ig_NG) WeekdaysAbbreviated() []string {
	return ig.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ig *ig_NG) WeekdayNarrow(weekday time.Weekday) string {
	return ig.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ig *ig_NG) WeekdaysNarrow() []string {
	return ig.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ig *ig_NG) WeekdayShort(weekday time.Weekday) string {
	return ig.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ig *ig_NG) WeekdaysShort() []string {
	return ig.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ig *ig_NG) WeekdayWide(weekday time.Weekday) string {
	return ig.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ig *ig_NG) WeekdaysWide() []string {
	return ig.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ig_NG' and handles both Whole and Real numbers based on 'v'
func (ig *ig_NG) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ig_NG' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ig *ig_NG) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ig_NG'
func (ig *ig_NG) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ig.currencies[currency]
	l := len(s) + len(symbol) + 0 + 0*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ig.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ig.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, ig.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ig.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ig_NG'
// in accounting notation.
func (ig *ig_NG) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ig.currencies[currency]
	l := len(s) + len(symbol) + 2 + 0*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ig.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ig.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, ig.currencyNegativePrefix[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ig.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ig.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ig.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ig.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ig.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ig.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ig.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ig.periodsAbbreviated[0]...)
	} else {
		b = append(b, ig.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ig.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ig.periodsAbbreviated[0]...)
	} else {
		b = append(b, ig.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ig.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ig.periodsAbbreviated[0]...)
	} else {
		b = append(b, ig.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ig.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ig.periodsAbbreviated[0]...)
	} else {
		b = append(b, ig.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ig.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
