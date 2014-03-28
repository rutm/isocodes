package country

import (
	"strconv"
    . "pkg.appto.mobi/v1/isocodes"
)

const (
	_ = iota // skip zero
	ByAlpha2
	ByAlpha3
	ByNum
)

type country struct {
	alpha2   string
	alpha3   string
	num      Numeric
 	currency Numeric
	language string
}

var numeric map[Numeric]*country

var alpha2 map[string]*country

var alpha3 map[string]*country

func (c *country) Alpha2() string {
	return c.alpha2
}

func (c *country) Alpha3() string {
	return c.alpha3
}

func (c *country) Num() Numeric {
	return c.num
}

func (c *country) Currency() Numeric {
	return c.currency
}

func (c *country) Language() string {
	return c.language
}


func init() {
	numeric = make(map[Numeric]*country)
	alpha2 = make(map[string]*country)
	alpha3 = make(map[string]*country)

	for i := range countries {
		numeric[countries[i].num] = countries[i]
		alpha2[countries[i].alpha2] = countries[i]
		alpha3[countries[i].alpha3] = countries[i]
	}
}

func Get(what string, by int) (*country, bool) {
	var c *country
	var err bool
	switch by {
	case ByAlpha2:
		c, err = alpha2[what]
	case ByAlpha3:
		c, err = alpha3[what]
	case ByNum:
		if code, e := strconv.Atoi(what); e == nil {
			c, err = numeric[Numeric(code)]
		}
	}

	return c, err
}
