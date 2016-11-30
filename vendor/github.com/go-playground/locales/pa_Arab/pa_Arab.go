package pa_Arab

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type pa_Arab struct {
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
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'pa_Arab' locale
func New() locales.Translator {
	return &pa_Arab{
		locale:                 "pa_Arab",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "‎-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "ر", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: "K",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: "K",
		monthsAbbreviated:      []string{"", "ਜਨ", "ਫ਼ਰ", "ਮਾਰਚ", "ਅਪ੍ਰੈ", "ਮਈ", "ਜੂਨ", "ਜੁਲਾ", "ਅਗ", "ਸਤੰ", "ਅਕਤੂ", "ਨਵੰ", "ਦਸੰ"},
		monthsNarrow:           []string{"", "ਜ", "ਫ਼", "ਮਾ", "ਅ", "ਮ", "ਜੂ", "ਜੁ", "ਅ", "ਸ", "ਅ", "ਨ", "ਦ"},
		monthsWide:             []string{"", "جنوری", "فروری", "مارچ", "اپریل", "مئ", "جون", "جولائی", "اگست", "ستمبر", "اکتوبر", "نومبر", "دسمبر"},
		daysAbbreviated:        []string{"ਐਤ", "ਸੋਮ", "ਮੰਗਲ", "ਬੁੱਧ", "ਵੀਰ", "ਸ਼ੁੱਕਰ", "ਸ਼ਨਿੱਚਰ"},
		daysNarrow:             []string{"ਐ", "ਸੋ", "ਮੰ", "ਬੁੱ", "ਵੀ", "ਸ਼ੁੱ", "ਸ਼"},
		daysShort:              []string{"ਐਤ", "ਸੋਮ", "ਮੰਗ", "ਬੁੱਧ", "ਵੀਰ", "ਸ਼ੁੱਕ", "ਸ਼ਨਿੱ"},
		daysWide:               []string{"اتوار", "پیر", "منگل", "بُدھ", "جمعرات", "جمعہ", "ہفتہ"},
		periodsAbbreviated:     []string{"ਪੂ.ਦੁ.", "ਬਾ.ਦੁ."},
		periodsNarrow:          []string{"ਸ.", "ਸ਼."},
		periodsWide:            []string{"ਪੂ.ਦੁ.", "ਬਾ.ਦੁ."},
		erasAbbreviated:        []string{"ايساپورو", "سں"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"ايساپورو", "سں"},
		timezones:              map[string]string{"COST": "ਕੋਲੰਬੀਆ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "WAT": "ਪੱਛਮੀ ਅਫਰੀਕਾ ਮਿਆਰੀ ਸਮਾਂ", "CST": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਕੇਂਦਰੀ ਮਿਆਰੀ ਸਮਾਂ", "OEZ": "ਪੂਰਬੀ ਯੂਰਪੀ ਮਿਆਰੀ ਸਮਾਂ", "WART": "ਪੱਛਮੀ ਅਰਜਨਟੀਨਾ ਮਿਆਰੀ ਸਮਾਂ", "EDT": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਪੂਰਬੀ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "UYT": "ਉਰੂਗਵੇ ਮਿਆਰੀ ਸਮਾਂ", "COT": "ਕੋਲੰਬੀਆ ਮਿਆਰੀ ਸਮਾਂ", "WIB": "ਪੱਛਮੀ ਇੰਡੋਨੇਸ਼ੀਆ ਸਮਾਂ", "AKST": "ਅਲਾਸਕਾ ਮਿਆਰੀ ਸਮਾਂ", "PDT": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਪੈਸਿਫਿਕ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "AEDT": "ਆਸਟ੍ਰੇਲੀਆਈ ਪੂਰਬੀ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "IST": "ਭਾਰਤੀ ਮਿਆਰੀ ਸਮਾਂ", "HNT": "ਨਿਊਫਾਉਂਡਲੈਂਡ ਮਿਆਰੀ ਸਮਾਂ", "OESZ": "ਪੂਰਬੀ ਯੂਰਪੀ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "MESZ": "ਮੱਧ ਯੂਰਪੀ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "ACST": "ਆਸਟ੍ਰੇਲੀਆਈ ਕੇਂਦਰੀ ਮਿਆਰੀ ਸਮਾਂ", "ECT": "ਇਕਵੇਡੋਰ ਸਮਾਂ", "ACWDT": "ਆਸਟ੍ਰੇਲੀਆਈ ਕੇਂਦਰੀ ਪੱਛਮੀ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "PST": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਪੈਸਿਫਿਕ ਮਿਆਰੀ ਸਮਾਂ", "BOT": "ਬੋਲੀਵੀਆ ਸਮਾਂ", "WIT": "ਪੂਰਬੀ ਇੰਡੋਨੇਸ਼ੀਆ ਸਮਾਂ", "LHDT": "ਲੌਰਡ ਹੋਵੇ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "ART": "ਅਰਜਨਟੀਨਾ ਮਿਆਰੀ ਸਮਾਂ", "WARST": "ਪੱਛਮੀ ਅਰਜਨਟੀਨਾ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "EAT": "ਪੂਰਬੀ ਅਫਰੀਕਾ ਸਮਾਂ", "WITA": "ਮੱਧ ਇੰਡੋਨੇਸ਼ੀਆਈ ਸਮਾਂ", "WESZ": "ਪੱਛਮੀ ਯੂਰਪੀ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "CHADT": "ਚੈਥਮ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "ChST": "ਚਾਮੋਰੋ ਮਿਆਰੀ ਸਮਾਂ", "GMT": "ਗ੍ਰੀਨਵਿਚ ਮੀਨ ਟਾਈਮ", "WEZ": "ਪੱਛਮੀ ਯੂਰਪੀ ਮਿਆਰੀ ਸਮਾਂ", "TMT": "ਤੁਰਕਮੇਨਿਸਤਾਨ ਮਿਆਰੀ ਸਮਾਂ", "GYT": "ਗੁਯਾਨਾ ਸਮਾਂ", "LHST": "ਲੌਰਡ ਹੋਵੇ ਮਿਆਰੀ ਸਮਾਂ", "JST": "ਜਪਾਨ ਮਿਆਰੀ ਸਮਾਂ", "∅∅∅": "ਬ੍ਰਾਜ਼ੀਲੀਆ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "MST": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਮਾਉਂਟੇਨ ਮਿਆਰੀ ਸਮਾਂ", "AWDT": "ਆਸਟ੍ਰੇਲੀਆਈ ਪੱਛਮੀ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "SAST": "ਦੱਖਣੀ ਅਫ਼ਰੀਕਾ ਮਿਆਰੀ ਸਮਾਂ", "WAST": "ਪੱਛਮੀ ਅਫਰੀਕਾ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "EST": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਪੂਰਬੀ ਮਿਆਰੀ ਸਮਾਂ", "ADT": "ਅਟਲਾਂਟਿਕ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "MDT": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਮਾਉਂਟੇਨ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "CHAST": "ਚੈਥਮ ਮਿਆਰੀ ਸਮਾਂ", "GFT": "ਫ੍ਰੈਂਚ ਗੁਏਨਾ ਸਮਾਂ", "MYT": "ਮਲੇਸ਼ੀਆ ਸਮਾਂ", "HKST": "ਹਾਂਗ ਕਾਂਗ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "HAST": "ਹਵਾਈ-ਅਲੇਯੂਸ਼ਿਅਨ ਮਿਆਰੀ ਸਮਾਂ", "ACWST": "ਆਸਟ੍ਰੇਲੀਆਈ ਕੇਂਦਰੀ ਪੱਛਮੀ ਮਿਆਰੀ ਸਮਾਂ", "CDT": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਕੇਂਦਰੀ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "NZST": "ਨਿਊਜ਼ੀਲੈਂਡ ਮਿਆਰੀ ਸਮਾਂ", "JDT": "ਜਪਾਨ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "NZDT": "ਨਿਊਜ਼ੀਲੈਂਡ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "BT": "ਭੂਟਾਨ ਸਮਾਂ", "CLST": "ਚਿਲੀ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "SRT": "ਸੂਰੀਨਾਮ ਸਮਾਂ", "ACDT": "ਆਸਟ੍ਰੇਲੀਆਈ ਕੇਂਦਰੀ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "ARST": "ਅਰਜਨਟੀਨਾ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "TMST": "ਤੁਰਕਮੇਨਿਸਤਾਨ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "CAT": "ਕੇਂਦਰੀ ਅਫਰੀਕਾ ਸਮਾਂ", "MEZ": "ਮੱਧ ਯੂਰਪੀ ਮਿਆਰੀ ਸਮਾਂ", "SGT": "ਸਿੰਗਾਪੁਰ ਮਿਆਰੀ ਸਮਾਂ", "VET": "ਵੈਨੇਜ਼ੂਏਲਾ ਸਮਾਂ", "HADT": "ਹਵਾਈ-ਅਲੇਯੂਸ਼ਿਅਨ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "HAT": "ਨਿਊਫਾਉਂਡਲੈਂਡ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "HKT": "ਹਾਂਗ ਕਾਂਗ ਮਿਆਰੀ ਸਮਾਂ", "AEST": "ਆਸਟ੍ਰੇਲੀਆਈ ਪੂਰਬੀ ਮਿਆਰੀ ਸਮਾਂ", "AWST": "ਆਸਟ੍ਰੇਲੀਆਈ ਪੱਛਮੀ ਮਿਆਰੀ ਸਮਾਂ", "CLT": "ਚਿਲੀ ਮਿਆਰੀ ਸਮਾਂ", "UYST": "ਉਰੂਗਵੇ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "AKDT": "ਅਲਾਸਕਾ ਪ੍ਰਕਾਸ਼ ਸਮਾਂ", "AST": "ਅਟਲਾਂਟਿਕ ਮਿਆਰੀ ਸਮਾਂ"},
	}
}

// Locale returns the current translators string locale
func (pa *pa_Arab) Locale() string {
	return pa.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'pa_Arab'
func (pa *pa_Arab) PluralsCardinal() []locales.PluralRule {
	return pa.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'pa_Arab'
func (pa *pa_Arab) PluralsOrdinal() []locales.PluralRule {
	return pa.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'pa_Arab'
func (pa *pa_Arab) PluralsRange() []locales.PluralRule {
	return pa.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'pa_Arab'
func (pa *pa_Arab) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'pa_Arab'
func (pa *pa_Arab) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'pa_Arab'
func (pa *pa_Arab) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := pa.CardinalPluralRule(num1, v1)
	end := pa.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (pa *pa_Arab) MonthAbbreviated(month time.Month) string {
	return pa.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (pa *pa_Arab) MonthsAbbreviated() []string {
	return pa.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (pa *pa_Arab) MonthNarrow(month time.Month) string {
	return pa.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (pa *pa_Arab) MonthsNarrow() []string {
	return pa.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (pa *pa_Arab) MonthWide(month time.Month) string {
	return pa.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (pa *pa_Arab) MonthsWide() []string {
	return pa.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (pa *pa_Arab) WeekdayAbbreviated(weekday time.Weekday) string {
	return pa.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (pa *pa_Arab) WeekdaysAbbreviated() []string {
	return pa.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (pa *pa_Arab) WeekdayNarrow(weekday time.Weekday) string {
	return pa.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (pa *pa_Arab) WeekdaysNarrow() []string {
	return pa.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (pa *pa_Arab) WeekdayShort(weekday time.Weekday) string {
	return pa.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (pa *pa_Arab) WeekdaysShort() []string {
	return pa.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (pa *pa_Arab) WeekdayWide(weekday time.Weekday) string {
	return pa.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (pa *pa_Arab) WeekdaysWide() []string {
	return pa.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'pa_Arab' and handles both Whole and Real numbers based on 'v'
func (pa *pa_Arab) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pa.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, pa.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(pa.minus) - 1; j >= 0; j-- {
			b = append(b, pa.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'pa_Arab' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (pa *pa_Arab) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 6
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pa.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(pa.minus) - 1; j >= 0; j-- {
			b = append(b, pa.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, pa.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'pa_Arab'
func (pa *pa_Arab) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pa.currencies[currency]
	l := len(s) + len(symbol) + 8

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pa.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(pa.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, pa.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(pa.minus) - 1; j >= 0; j-- {
			b = append(b, pa.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, pa.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'pa_Arab'
// in accounting notation.
func (pa *pa_Arab) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pa.currencies[currency]
	l := len(s) + len(symbol) + 8

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pa.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(pa.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, pa.currencyNegativePrefix[j])
		}

		for j := len(pa.minus) - 1; j >= 0; j-- {
			b = append(b, pa.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(pa.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, pa.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, pa.currencyNegativeSuffix...)
	} else {

		b = append(b, pa.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'pa_Arab'
func (pa *pa_Arab) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'pa_Arab'
func (pa *pa_Arab) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pa.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'pa_Arab'
func (pa *pa_Arab) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pa.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'pa_Arab'
func (pa *pa_Arab) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, pa.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pa.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'pa_Arab'
func (pa *pa_Arab) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, pa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, pa.periodsAbbreviated[0]...)
	} else {
		b = append(b, pa.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'pa_Arab'
func (pa *pa_Arab) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, pa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, pa.periodsAbbreviated[0]...)
	} else {
		b = append(b, pa.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'pa_Arab'
func (pa *pa_Arab) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, pa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, pa.periodsAbbreviated[0]...)
	} else {
		b = append(b, pa.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'pa_Arab'
func (pa *pa_Arab) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, pa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, pa.periodsAbbreviated[0]...)
	} else {
		b = append(b, pa.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := pa.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
