// Package iso4217 provides the ISO 4217 codes for the representation of currencies and funds.
package iso4217

// Currency is a representation of a currency or fund.
type Currency uint16

// Alpha returns the ISO 4217 three-letter alphabetic code.
func (c Currency) Alpha() string { return currencies[c].alpha }

// Numeric returns the ISO 4217 three-digit numeric code.
func (c Currency) Numeric() string { return currencies[c].numeric }

// Exponent returns the decimal point location.
func (c Currency) Exponent() int { return currencies[c].exponent }

// Name returns the English name.
func (c Currency) Name() string { return currencies[c].name }

// FromAlpha returns Currency for the three-letter alpha code.
// Or an error if it does not exist.
func FromAlpha(alpha string) (Currency, error) {
	if c, ok := fromAlpha[alpha]; ok {
		return c, nil
	}
	return Currency(0), Error("no currency exists with alphabetic code " + alpha)
}

// FromNumeric returns Currency for the three-digit numeric code.
// Or an error if it does not exist.
func FromNumeric(numeric string) (Currency, error) {
	if c, ok := fromNumeric[numeric]; ok {
		return c, nil
	}
	return Currency(0), Error("no currency exists with numeric code " + numeric)
}

// Must panics if err is non-nil and otherwise returns c.
// Could be used to return a single value from FromAlpha/FromNumeric.
func Must(c Currency, err error) Currency {
	if err != nil {
		panic(err)
	}
	return c
}

// Error is the type of error returned by this package
type Error string

func (e Error) Error() string { return "iso4217: " + string(e) }
