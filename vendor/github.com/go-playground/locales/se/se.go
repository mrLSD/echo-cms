package se

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type se struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'se' locale
func New() locales.Translator {
	return &se{
		locale:                 "se",
		pluralsCardinal:        []locales.PluralRule{2, 3, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "−",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "Dkr", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HK$", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "₹", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JP¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "kr", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "Skr", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "ođđj", "guov", "njuk", "cuo", "mies", "geas", "suoi", "borg", "čakč", "golg", "skáb", "juov"},
		monthsNarrow:           []string{"", "O", "G", "N", "C", "M", "G", "S", "B", "Č", "G", "S", "J"},
		monthsWide:             []string{"", "ođđajagemánnu", "guovvamánnu", "njukčamánnu", "cuoŋománnu", "miessemánnu", "geassemánnu", "suoidnemánnu", "borgemánnu", "čakčamánnu", "golggotmánnu", "skábmamánnu", "juovlamánnu"},
		daysAbbreviated:        []string{"sotn", "vuos", "maŋ", "gask", "duor", "bear", "láv"},
		daysNarrow:             []string{"S", "V", "M", "G", "D", "B", "L"},
		daysWide:               []string{"sotnabeaivi", "vuossárga", "maŋŋebárga", "gaskavahkku", "duorasdat", "bearjadat", "lávvardat"},
		periodsAbbreviated:     []string{"i.b.", "e.b."},
		periodsWide:            []string{"iđitbeaivet", "eahketbeaivet"},
		erasAbbreviated:        []string{"o.Kr.", "m.Kr."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"ovdal Kristtusa", "maŋŋel Kristtusa"},
		timezones:              map[string]string{"UYST": "UYST", "ECT": "ECT", "GYT": "GYT", "UYT": "UYT", "SRT": "SRT", "AST": "AST", "WAT": "WAT", "CLT": "CLT", "WIT": "WIT", "CHADT": "CHADT", "COST": "COST", "SGT": "SGT", "SAST": "SAST", "CDT": "CDT", "MDT": "MDT", "CLST": "CLST", "∅∅∅": "∅∅∅", "HKT": "HKT", "ACST": "ACST", "AKDT": "AKDT", "MST": "MST", "NZST": "NZST", "LHDT": "LHDT", "HKST": "HKST", "COT": "COT", "WART": "WART", "HAST": "HAST", "AWDT": "AWDT", "OEZ": "nuorti-Eurohpá dábálašáigi", "ART": "ART", "EDT": "EDT", "ADT": "ADT", "EST": "EST", "PST": "PST", "WITA": "WITA", "OESZ": "nuorti-Eurohpá geassiáigi", "MEZ": "gaska-Eurohpá dábálašáigi", "BOT": "BOT", "AEDT": "AEDT", "IST": "IST", "HNT": "HNT", "CAT": "CAT", "JST": "JST", "CHAST": "CHAST", "ARST": "ARST", "MESZ": "gaska-Eurohpá geassiáigi", "ACDT": "ACDT", "TMST": "TMST", "HAT": "HAT", "AWST": "AWST", "GMT": "Greenwich gaskka áigi", "GFT": "GFT", "VET": "VET", "CST": "CST", "WARST": "WARST", "HADT": "HADT", "ACWST": "ACWST", "BT": "BT", "JDT": "JDT", "LHST": "LHST", "WEZ": "oarje-Eurohpá dábálašáigi", "WIB": "WIB", "EAT": "EAT", "ChST": "ChST", "AEST": "AEST", "AKST": "AKST", "TMT": "TMT", "WAST": "WAST", "WESZ": "oarje-Eurohpá geassiáigi", "ACWDT": "ACWDT", "PDT": "PDT", "MYT": "MYT", "NZDT": "NZDT"},
	}
}

// Locale returns the current translators string locale
func (se *se) Locale() string {
	return se.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'se'
func (se *se) PluralsCardinal() []locales.PluralRule {
	return se.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'se'
func (se *se) PluralsOrdinal() []locales.PluralRule {
	return se.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'se'
func (se *se) PluralsRange() []locales.PluralRule {
	return se.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'se'
func (se *se) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'se'
func (se *se) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'se'
func (se *se) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (se *se) MonthAbbreviated(month time.Month) string {
	return se.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (se *se) MonthsAbbreviated() []string {
	return se.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (se *se) MonthNarrow(month time.Month) string {
	return se.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (se *se) MonthsNarrow() []string {
	return se.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (se *se) MonthWide(month time.Month) string {
	return se.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (se *se) MonthsWide() []string {
	return se.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (se *se) WeekdayAbbreviated(weekday time.Weekday) string {
	return se.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (se *se) WeekdaysAbbreviated() []string {
	return se.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (se *se) WeekdayNarrow(weekday time.Weekday) string {
	return se.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (se *se) WeekdaysNarrow() []string {
	return se.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (se *se) WeekdayShort(weekday time.Weekday) string {
	return se.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (se *se) WeekdaysShort() []string {
	return se.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (se *se) WeekdayWide(weekday time.Weekday) string {
	return se.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (se *se) WeekdaysWide() []string {
	return se.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'se' and handles both Whole and Real numbers based on 'v'
func (se *se) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, se.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(se.group) - 1; j >= 0; j-- {
					b = append(b, se.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(se.minus) - 1; j >= 0; j-- {
			b = append(b, se.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'se' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (se *se) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 7
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, se.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(se.minus) - 1; j >= 0; j-- {
			b = append(b, se.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, se.percentSuffix...)

	b = append(b, se.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'se'
func (se *se) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := se.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, se.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(se.group) - 1; j >= 0; j-- {
					b = append(b, se.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(se.minus) - 1; j >= 0; j-- {
			b = append(b, se.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, se.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, se.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'se'
// in accounting notation.
func (se *se) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := se.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, se.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(se.group) - 1; j >= 0; j-- {
					b = append(b, se.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(se.minus) - 1; j >= 0; j-- {
			b = append(b, se.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, se.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, se.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, se.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'se'
func (se *se) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'se'
func (se *se) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'se'
func (se *se) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'se'
func (se *se) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'se'
func (se *se) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'se'
func (se *se) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'se'
func (se *se) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'se'
func (se *se) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}
