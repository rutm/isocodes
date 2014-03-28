package isocodes

import (
    "fmt"
)

type Numeric int

func (n Numeric) String() string {
    return fmt.Sprintf("%03d", n)
}


