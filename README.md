# iso4217
[![Go Reference](https://pkg.go.dev/badge/github.com/ferdypruis/iso4217.svg)](https://pkg.go.dev/github.com/ferdypruis/iso4217)
[![Go build](https://github.com/ferdypruis/iso4217/actions/workflows/go.yml/badge.svg)](https://github.com/ferdypruis/iso4217/actions/workflows/go.yml)

A Go package providing all ISO 4217 currency codes as constants of type `iso4217.Currency`.

For each currency the three-letter alphabetic code, the three-digit numeric code, the currency 
exponent and the English name are available.

The exponent for historic denominations, for which the constants are annotated as `deprecated`, 
is currently inaccurate.

Currencies can either be hardcoded using the available constants or loaded from a string using
`FromAlpha()` and `FromNumeric()`.

## Examples
Use the constants to directly reference currencies.
```go
rates := []struct {
    currency iso4217.Currency
    amount   float32
}{
    {iso4217.CAD, 1.30488},
    {iso4217.JPY, 110.10},
    {iso4217.CLF, 0.028043},
}

fmt.Println("Today, one United States Dollar equals;")
for _, r := range rates {
    fmt.Printf("- %.*f %s\n", r.currency.Exponent(), r.amount, r.currency.Name())
}

// Output:
// Today, one United States Dollar equals;
// - 1.30 Canadian Dollar
// - 110 Yen
// - 0.0280 Unidad de Fomento
```

Use `FromAlpha()` and `FromNumeric()` to load a currency from a string.
```go
currency, _ := iso4217.FromNumeric("608") // Ignoring error for simplicity
fmt.Println(currency.Alpha(), "is the alphabetic currency code for the", currency.Name())

// Output:
// PHP is the alphabetic currency code for the Philippine Peso
```

## Source
- [currency-iso.org](https://www.currency-iso.org/en/home/tables.html)
