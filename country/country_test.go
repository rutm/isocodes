package country

import (
	"fmt"
	"testing"
)

var tests = []struct {
	what   string
	by     int
	alpha2 string
	alpha3 string
	num    int
	isReal bool
}{
	{"AQ", ByAlpha2, "AQ", "ATA", 10, true},
	{"10", ByNum, "AQ", "ATA", 10, true},
	{"ATA", ByAlpha3, "AQ", "ATA", 10, true},
	{"A1", ByAlpha2, "AQ", "ATA", 10, false},
	{"15", ByNum, "AQ", "ATA", 10, false},
	{"ZUZ", ByAlpha3, "AQ", "ATA", 10, false},
}

func TestGet(t *testing.T) {
	var c *country
	var ok bool
	for i, test := range tests {
		c, ok = Get(test.what, test.by)
		if !ok && test.isReal {
			t.Errorf("%d. %s is not exists %v", i, test.what, c)
		}
	}
}

func TestNumStr(t *testing.T) {
	c, _ := Get(tests[0].what, ByAlpha2)
	numStr := fmt.Sprintf("%03d", tests[0].num)

	if c.Num().String() != numStr {
		t.Errorf("numeric string %s representation of %v is wrong", numStr, c)
	}
}

func TestLanguage(t *testing.T) {
	c, _ := Get("USA", ByAlpha3)

	if c.Language() != "EN" {
		t.Errorf("wrong language code %s", c.language)
	}
}

func TestCurrency(t *testing.T) {
	c, _ := Get("USA", ByAlpha3)

	if c.Currency() != 840 {
		t.Errorf("wrong currency code %s", c.currency)
	}
}
