package vo

import "fmt"

type Money float64

func (m Money) String() string {
	return fmt.Sprintf("Rp%1.2f", m)
}

func MoneyZero() Money {
	// TODO : REMOVE COMMENT THEN
	return Money(0)
}

func MoneyUndefined() Money {
	return Money(-1)
}

func (m Money) ToFloat64() float64 {
	return float64(m)
}

func (m Money) IsUndefined() bool {
	return float64(m) < 0
}
