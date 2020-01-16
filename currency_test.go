package iso4217_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/ferdypruis/iso4217"
)

func TestCurrency_Exponent(t *testing.T) {
	tests := []struct {
		name string
		c    iso4217.Currency
		want int
	}{
		{"AED", iso4217.AED, 2},
		{"XTS", iso4217.XTS, 0},
		{"CLF", iso4217.CLF, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Exponent(); got != tt.want {
				t.Errorf("Exponent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrency_Name(t *testing.T) {
	tests := []struct {
		name string
		c    iso4217.Currency
		want string
	}{
		{"ADP", iso4217.ADP, "Andorran Peseta"},
		{"XTS", iso4217.XTS, "Codes specifically reserved for testing purposes"},
		{"ZWR", iso4217.ZWR, "Zimbabwe Dollar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromAlpha(t *testing.T) {
	want := iso4217.EUR
	got, err := iso4217.FromAlpha("EUR")
	if err != nil {
		t.Errorf("FromAlpha() error = %v, wantErr nil", err)
	}
	if got != want {
		t.Errorf("FromAlpha() got = %v, want %v", got, want)
	}
}

func TestFromAlphaError(t *testing.T) {
	_, err := iso4217.FromAlpha("000")
	if _, ok := err.(iso4217.Error); !ok {
		t.Fatalf("FromAlpha() error %T, want iso4217.Error", err)
	}
}

func TestFromNumeric(t *testing.T) {
	want := iso4217.USD
	got, err := iso4217.FromNumeric("840")
	if err != nil {
		t.Errorf("FromNumeric() error = %v, wantErr nil", err)
	}
	if got != want {
		t.Errorf("FromNumeric() got = %v, want %v", got, want)
	}
}

func TestFromNumericError(t *testing.T) {
	_, err := iso4217.FromNumeric("AAA")
	if _, ok := err.(iso4217.Error); !ok {
		t.Fatalf("FromNumeric() error %T, want iso4217.Error", err)
	}
}

func TestMust(t *testing.T) {
	want := iso4217.XTS
	got := iso4217.Must(want, nil)

	if got != want {
		t.Errorf("Must() got = %v, want %v", got, want)
	}
}

func TestMustPanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Must() did not panic")
		}
	}()

	iso4217.Must(iso4217.Currency(0), errors.New("this should cause panic"))
}

func TestErrorString(t *testing.T) {
	err := iso4217.Error("test error")
	if !strings.HasPrefix(err.Error(), "iso4217:") {
		t.Fatalf("Error.String() %q, want prefix 'iso4217:'", err)
	}
}

func ExampleFromAlpha() {
	currency, _ := iso4217.FromAlpha("EUR") // Ignoring error for simplicity
	fmt.Println("The three-digit numeric code for the", currency.Name(), "is", currency.Numeric())

	// Output:
	// The three-digit numeric code for the Euro is 978
}

func ExampleFromNumeric() {
	currency, _ := iso4217.FromNumeric("608") // Ignoring error for simplicity
	fmt.Println(currency.Alpha(), "is the alphabetic currency code for the", currency.Name())

	// Output:
	// PHP is the alphabetic currency code for the Philippine Peso
}

func ExampleMust() {
	fmt.Println("The numeric code for euros is", iso4217.Must(iso4217.FromAlpha("EUR")).Numeric())

	// Output:
	// The numeric code for euros is 978
}

func ExampleCurrency_Exponent() {
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
}
