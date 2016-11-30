package uz_Latn_UZ

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type uz_Latn_UZ struct {
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
	currencyPositivePrefix string
	currencyNegativePrefix string
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

// New returns a new instance of translator for the 'uz_Latn_UZ' locale
func New() locales.Translator {
	return &uz_Latn_UZ{
		locale:                 "uz_Latn_UZ",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "yan", "fev", "mar", "apr", "may", "iyn", "iyl", "avg", "sen", "okt", "noy", "dek"},
		monthsNarrow:           []string{"", "Y", "F", "M", "A", "M", "I", "I", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "yanvar", "fevral", "mart", "aprel", "may", "iyun", "iyul", "avgust", "Sentabr", "Oktabr", "noyabr", "dekabr"},
		daysAbbreviated:        []string{"Ya", "Du", "Se", "Ch", "Pa", "Ju", "Sh"},
		daysNarrow:             []string{"Y", "D", "S", "C", "P", "J", "S"},
		daysShort:              []string{"Ya", "Du", "Se", "Ch", "Pa", "Ju", "Sh"},
		daysWide:               []string{"yakshanba", "dushanba", "seshanba", "chorshanba", "payshanba", "juma", "shanba"},
		periodsAbbreviated:     []string{"TO", "TK"},
		periodsNarrow:          []string{"TO", "TK"},
		periodsWide:            []string{"TO", "TK"},
		erasAbbreviated:        []string{"", ""},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"WAST": "Gʻarbiy Afrika yozgi vaqti", "EAT": "Sharqiy Afrika vaqti", "CLT": "Chili standart vaqti", "GMT": "Grinvich o‘rtacha vaqti", "WESZ": "G‘arbiy Yevropa yozgi vaqti", "CHAST": "Chatem standart vaqti", "OEZ": "Sharqiy Yevropa standart vaqti", "SAST": "Janubiy Afrika standart vaqti", "COST": "Kolumbiya yozgi vaqti", "ECT": "Ekvador vaqti", "MEZ": "Markaziy Yevropa standart vaqti", "MESZ": "Markaziy Yevropa yozgi vaqti", "OESZ": "Sharqiy Yevropa yozgi vaqti", "HNT": "Nyufaundlend standart vaqti", "COT": "Kolumbiya standart vaqti", "ACWST": "Markaziy Avstraliya g‘arbiy standart vaqti", "WITA": "Markaziy Indoneziya vaqti", "LHST": "Lord-Xau standart vaqti", "SGT": "Singapur vaqti", "ACWDT": "Markaziy Avstraliya g‘arbiy yozgi vaqti", "EDT": "Sharqiy Amerika yozgi vaqti", "VET": "Venesuela vaqti", "AEST": "Sharqiy Avstraliya standart vaqti", "WARST": "Gʻarbiy Argentina yozgi vaqti", "ChST": "Chamorro standart vaqti", "GYT": "Gayana vaqti", "NZDT": "Yangi Zelandiya yozgi vaqti", "HKST": "Gonkong yozgi vaqti", "PDT": "Shimoliy Amerika Tinch okeani yozgi vaqti", "NZST": "Yangi Zelandiya standart vaqti", "CLST": "Chili yozgi vaqti", "WEZ": "G‘arbiy Yevropa standart vaqti", "SRT": "Surinam vaqti", "AKST": "Alyaska standart vaqti", "MDT": "Shimoliy Amerika togʻ yozgi vaqti", "ADT": "Atlantika yozgi vaqti", "∅∅∅": "Azor orollari yozgi vaqti", "GFT": "Fransuz Gvianasi vaqti", "ACDT": "Markaziy Avstraliya yozgi vaqti", "HAST": "Gavayi-aleut standart vaqti", "EST": "Sharqiy Amerika standart vaqti", "AKDT": "Alyaska yozgi vaqti", "MST": "Shimoliy Amerika togʻ standart vaqti", "AEDT": "Sharqiy Avstraliya yozgi vaqti", "WAT": "Gʻarbiy Afrika standart vaqti", "HAT": "Nyufaundlend yozgi vaqti", "LHDT": "Lord-Xau yozgi vaqti", "ACST": "Markaziy Avstraliya standart vaqti", "WIB": "Gʻarbiy Indoneziya vaqti", "ARST": "Argentina yozgi vaqti", "AWST": "G‘arbiy Avstraliya standart vaqti", "MYT": "Malayziya vaqti", "BOT": "Boliviya vaqti", "AST": "Atlantika standart vaqti", "TMT": "Turkmaniston standart vaqti", "TMST": "Turkmaniston yozgi vaqti", "BT": "Butan vaqti", "WIT": "Sharqiy Indoneziya vaqti", "JST": "Yaponiya standart vaqti", "PST": "Shimoliy Amerika Tinch okeani standart vaqti", "UYT": "Urugvay standart vaqti", "UYST": "Urugvay yozgi vaqti", "AWDT": "G‘arbiy Avstraliya yozgi vaqti", "JDT": "Yaponiya yozgi vaqti", "CHADT": "Chatem yozgi vaqti", "ART": "Argentina standart vaqti", "CST": "Markaziy Amerika standart vaqti", "CAT": "Markaziy Afrika vaqti", "WART": "Gʻarbiy Argentina standart vaqti", "CDT": "Markaziy Amerika yozgi vaqti", "HKT": "Gonkong standart vaqti", "HADT": "Gavayi-aleut yozgi vaqti", "IST": "Hindiston vaqti"},
	}
}

// Locale returns the current translators string locale
func (uz *uz_Latn_UZ) Locale() string {
	return uz.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) PluralsCardinal() []locales.PluralRule {
	return uz.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) PluralsOrdinal() []locales.PluralRule {
	return uz.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) PluralsRange() []locales.PluralRule {
	return uz.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := uz.CardinalPluralRule(num1, v1)
	end := uz.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (uz *uz_Latn_UZ) MonthAbbreviated(month time.Month) string {
	return uz.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (uz *uz_Latn_UZ) MonthsAbbreviated() []string {
	return uz.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (uz *uz_Latn_UZ) MonthNarrow(month time.Month) string {
	return uz.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (uz *uz_Latn_UZ) MonthsNarrow() []string {
	return uz.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (uz *uz_Latn_UZ) MonthWide(month time.Month) string {
	return uz.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (uz *uz_Latn_UZ) MonthsWide() []string {
	return uz.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (uz *uz_Latn_UZ) WeekdayAbbreviated(weekday time.Weekday) string {
	return uz.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (uz *uz_Latn_UZ) WeekdaysAbbreviated() []string {
	return uz.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (uz *uz_Latn_UZ) WeekdayNarrow(weekday time.Weekday) string {
	return uz.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (uz *uz_Latn_UZ) WeekdaysNarrow() []string {
	return uz.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (uz *uz_Latn_UZ) WeekdayShort(weekday time.Weekday) string {
	return uz.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (uz *uz_Latn_UZ) WeekdaysShort() []string {
	return uz.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (uz *uz_Latn_UZ) WeekdayWide(weekday time.Weekday) string {
	return uz.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (uz *uz_Latn_UZ) WeekdaysWide() []string {
	return uz.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'uz_Latn_UZ' and handles both Whole and Real numbers based on 'v'
func (uz *uz_Latn_UZ) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uz.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uz.group) - 1; j >= 0; j-- {
					b = append(b, uz.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'uz_Latn_UZ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (uz *uz_Latn_UZ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uz.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, uz.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uz.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uz.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uz.group) - 1; j >= 0; j-- {
					b = append(b, uz.group[j])
				}
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

	for j := len(uz.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, uz.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uz.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'uz_Latn_UZ'
// in accounting notation.
func (uz *uz_Latn_UZ) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uz.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uz.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uz.group) - 1; j >= 0; j-- {
					b = append(b, uz.group[j])
				}
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

		for j := len(uz.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, uz.currencyNegativePrefix[j])
		}

		b = append(b, uz.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(uz.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, uz.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uz.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) FmtDateShort(t time.Time) string {

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

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, uz.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, uz.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, uz.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uz.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x29}...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'uz_Latn_UZ'
func (uz *uz_Latn_UZ) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := uz.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
