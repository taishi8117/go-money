package money

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestCurrency_Get(t *testing.T) {
	tcs := []struct {
		code     string
		expected string
	}{
		{"EUR", "EUR"},
		{"Eur", "EUR"},
	}

	for _, tc := range tcs {
		c := newCurrency(tc.code).get()

		if c.Code != tc.expected {
			t.Errorf("Expected %s got %s", tc.expected, c.Code)
		}
	}
}

func TestCurrency_Get1(t *testing.T) {
	code := "RANDOM"
	c := newCurrency(code).get()

	if c.Grapheme != code {
		t.Errorf("Expected %s got %s", c.Grapheme, code)
	}
}

func TestCurrency_Equals(t *testing.T) {
	tcs := []struct {
		code  string
		other string
	}{
		{"EUR", "EUR"},
		{"Eur", "EUR"},
		{"usd", "USD"},
	}

	for _, tc := range tcs {
		c := newCurrency(tc.code).get()
		oc := newCurrency(tc.other).get()

		if !c.equals(oc) {
			t.Errorf("Expected that %v is not equal %v", c, oc)
		}
	}
}

func TestCurrency_AddCurrency(t *testing.T) {
	tcs := []struct {
		code     string
		template string
	}{
		{"GOLD", "1$"},
	}

	for _, tc := range tcs {
		AddCurrency(tc.code, "", tc.template, "", "", 0)
		c := newCurrency(tc.code).get()

		if c.Template != tc.template {
			t.Errorf("Expected currency template %v got %v", tc.template, c.Template)
		}
	}
}

func TestCurrency_GetCurrency(t *testing.T) {
	code := "KLINGONDOLLAR"
	desired := Currency{Decimal: ".", Thousand: ",", Code: code, Fraction: 2, Grapheme: "$", Template: "$1"}
	AddCurrency(desired.Code, desired.Grapheme, desired.Template, desired.Decimal, desired.Thousand, desired.Fraction)
	currency := GetCurrency(code)
	if !reflect.DeepEqual(currency, &desired) {
		t.Errorf("Currencies do not match %+v got %+v", desired, currency)
	}
}

func TestCurrency_GetNonExistingCurrency(t *testing.T) {
	currency := GetCurrency("I*am*Not*a*Currency")
	if currency != nil {
		t.Errorf("Unexpected currency returned %+v", currency)
	}
}

func TestCurrency_Marshaling(t *testing.T) {
	given := GetCurrency("BTC")
	expected := `{"code":"BTC"}`

	b, err := json.Marshal(given)

	if err != nil {
		t.Error(err)
	}

	if string(b) != expected {
		t.Errorf("expected %s got %s", expected, string(b))
	}

}

func TestCurrency_Unmarshaling(t *testing.T) {
	given := `{"code": "JPY"}`
	expected := GetCurrency("JPY")

	var c Currency
	err := json.Unmarshal([]byte(given), &c)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(c)
	if !c.equals(expected) {
		t.Errorf("expected %s, got %s", expected.Code, c.Code)
	}

}
