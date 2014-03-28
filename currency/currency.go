package currency

import (
    "strconv"
    . "github.com/rutm/isocodes"
)


const (
    _ = iota // skip zero
    ByCode
    ByNum
    ByName
)

type currency struct {
    code string
    name string
    num Numeric
}

var numeric map[Numeric]*currency

var codes map[string]*currency

var names map[string]*currency



func (c *currency) Code() string {
    return c.code
}

func (c *currency) Num() Numeric {
    return c.num
}

func (c *currency) Name() string {
    return c.name
}


// TODO: avoid copy/paste from country
func init() {
	numeric = make(map[Numeric]*currency)
	codes = make(map[string]*currency)
	names = make(map[string]*currency)

	for i := range currencies {
		numeric[currencies[i].num] = currencies[i]
		codes[currencies[i].code] = currencies[i]
		names[currencies[i].name] = currencies[i]
	}
}

func Get(what string, by int) (*currency, bool) {
	var c *currency
	var err bool
	switch by {
	case ByCode:
		c, err = codes[what]
	case ByName:
		c, err = names[what]
	case ByNum:
		if code, e := strconv.Atoi(what); e == nil {
			c, err = numeric[Numeric(code)]
		}
	}

	return c, err
}
