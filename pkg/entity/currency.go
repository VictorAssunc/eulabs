package entity

import (
	"fmt"
)

type Currency uint

func (c Currency) String() string {
	return fmt.Sprintf("%d.%02d", c/100, c%100)
}
