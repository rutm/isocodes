package currency

import (
	"testing"
)

var tests = []struct {
	what   string
	by     int
	isReal bool
}{
	{"ARS", ByCode, true},
	{"032", ByNum, true},
}

func TestGet(t *testing.T) {
	var c *currency
	var ok bool
	for i, test := range tests {
		c, ok = Get(test.what, test.by)
		if !ok && test.isReal {
			t.Errorf("%d. %s is not exists %v", i, test.what, c)
		}
	}
}


