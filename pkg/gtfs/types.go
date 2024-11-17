package gtfs

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Color string

var validColor = regexp.MustCompile(`(?i)^[a-f\d]{6}$`)

func (c *Color) parseColor(v string) error {
	f := strings.TrimSpace(v)
	if !validColor.MatchString(f) {
		return fmt.Errorf("invalid color: %s", v)
	}

	*c = Color(strings.ToUpper(f))
	return nil
}

type currencyCode string

var validCurrencyCodes = map[string]int{
	"AED": 2,
	"AFN": 2,
	"ALL": 2,
	"AMD": 2,
	"ANG": 2,
	"AOA": 2,
	"ARS": 2,
	"AUD": 2,
	"AWG": 2,
	"AZN": 2,
	"BAM": 2,
	"BBD": 2,
	"BDT": 2,
	"BGN": 2,
	"BHD": 3,
	"BIF": 0,
	"BMD": 2,
	"BND": 2,
	"BOB": 2,
	"BOV": 2,
	"BRL": 2,
	"BSD": 2,
	"BTN": 2,
	"BWP": 2,
	"BYN": 2,
	"BZD": 2,
	"CAD": 2,
	"CDF": 2,
	"CHE": 2,
	"CHF": 2,
	"CHW": 2,
	"CLF": 4,
	"CLP": 0,
	"CNY": 2,
	"COP": 2,
	"COU": 2,
	"CRC": 2,
	"CUP": 2,
	"CVE": 2,
	"CZK": 2,
	"DJF": 0,
	"DKK": 2,
	"DOP": 2,
	"DZD": 2,
	"EGP": 2,
	"ERN": 2,
	"ETB": 2,
	"EUR": 2,
	"FJD": 2,
	"FKP": 2,
	"GBP": 2,
	"GEL": 2,
	"GHS": 2,
	"GIP": 2,
	"GMD": 2,
	"GNF": 0,
	"GTQ": 2,
	"GYD": 2,
	"HKD": 2,
	"HNL": 2,
	"HTG": 2,
	"HUF": 2,
	"IDR": 2,
	"ILS": 2,
	"INR": 2,
	"IQD": 3,
	"IRR": 2,
	"ISK": 0,
	"JMD": 2,
	"JOD": 3,
	"JPY": 0,
	"KES": 2,
	"KGS": 2,
	"KHR": 2,
	"KMF": 0,
	"KPW": 2,
	"KRW": 0,
	"KWD": 3,
	"KYD": 2,
	"KZT": 2,
	"LAK": 2,
	"LBP": 2,
	"LKR": 2,
	"LRD": 2,
	"LSL": 2,
	"LYD": 3,
	"MAD": 2,
	"MDL": 2,
	"MGA": 2,
	"MKD": 2,
	"MMK": 2,
	"MNT": 2,
	"MOP": 2,
	"MRU": 2,
	"MUR": 2,
	"MVR": 2,
	"MWK": 2,
	"MXN": 2,
	"MXV": 2,
	"MYR": 2,
	"MZN": 2,
	"NAD": 2,
	"NGN": 2,
	"NIO": 2,
	"NOK": 2,
	"NPR": 2,
	"NZD": 2,
	"OMR": 3,
	"PAB": 2,
	"PEN": 2,
	"PGK": 2,
	"PHP": 2,
	"PKR": 2,
	"PLN": 2,
	"PYG": 0,
	"QAR": 2,
	"RON": 2,
	"RSD": 2,
	"RUB": 2,
	"RWF": 0,
	"SAR": 2,
	"SBD": 2,
	"SCR": 2,
	"SDG": 2,
	"SEK": 2,
	"SGD": 2,
	"SHP": 2,
	"SLE": 2,
	"SOS": 2,
	"SRD": 2,
	"SSP": 2,
	"STN": 2,
	"SVC": 2,
	"SYP": 2,
	"SZL": 2,
	"THB": 2,
	"TJS": 2,
	"TMT": 2,
	"TND": 3,
	"TOP": 2,
	"TRY": 2,
	"TTD": 2,
	"TWD": 2,
	"TZS": 2,
	"UAH": 2,
	"UGX": 0,
	"USD": 2,
	"USN": 2,
	"UYI": 0,
	"UYU": 2,
	"UYW": 4,
	"UZS": 2,
	"VED": 2,
	"VES": 2,
	"VND": 0,
	"VUV": 0,
	"WST": 2,
	"YER": 2,
	"ZAR": 2,
	"ZMW": 2,
	"ZWG": 2,
}

func (c *currencyCode) parseCurrencyCode(v string) error {
	f := strings.TrimSpace(v)
	f = strings.ToUpper(f)
	if _, ok := validCurrencyCodes[f]; !ok {
		return fmt.Errorf("invalid currency code: %s", v)
	}

	*c = currencyCode(f)
	return nil
}

type Time time.Time

var validDate = regexp.MustCompile(`^\d{8}$`)
var dateFormat = "20060102"

func (t *Time) parse(v string) error {
	f := strings.TrimSpace(v)
	if !validDate.MatchString(f) {
		return fmt.Errorf("invalid date format: %s", v)
	}

	p, err := time.Parse(dateFormat, f)
	if err != nil {
		return fmt.Errorf("invalid date value: %s", v)
	}

	*t = Time(p)
	return nil
}

type Enum int

var (
	Availability int  = 1
	Available    Enum = 0
	Unavailable  Enum = 1

	Accessibility          int  = 2
	UnknownAccessibility   Enum = 0
	AccessibeForAtLeastOne Enum = 1
	NotAccessible          Enum = 2

	ContinuousPickup   int  = 3
	ContinuousDropOff  int  = 3
	DropOffType        int  = 3
	PickupType         int  = 3
	RegularlyScheduled Enum = 0
	NoneAvailable      Enum = 1
	MustPhoneAgency    Enum = 2
	MustCoordinate     Enum = 3

	Timepoint       int  = 1
	ApproximateTime Enum = 0
	ExactTime       Enum = 1
)

func (e *Enum) Parse(v string, u int) error {
	f := strings.TrimSpace(v)
	i, err := strconv.Atoi(f)
	if err != nil {
		return fmt.Errorf("invalid enum value: %s", v)
	}

	if i < 0 || i > u {
		return fmt.Errorf("enum out of bounds: %d", i)
	}

	*e = Enum(i)
	return nil
}

type Int int

func (i *Int) Parse(v string) error {
	f := strings.TrimSpace(v)
	p, err := strconv.Atoi(f)
	if err != nil {
		return fmt.Errorf("invalid integer value: %s", v)
	}
	*i = Int(p)
	return nil
}

type Float64 float64

func (fl *Float64) Parse(v string) error {
	f := strings.TrimSpace(v)
	p, err := strconv.ParseFloat(f, 64)
	if err != nil {
		return fmt.Errorf("invalid float value: %s", v)
	}

	*fl = Float64(p)
	return nil
}

type errorList []error

func (e *errorList) add(err error) error {
	if err == nil {
		return err
	}
	*e = append(*e, err)
	return err
}

type String string

func (s *String) Parse(v string) error {
	f := strings.TrimSpace(v)
	*s = String(f)
	return nil
}
