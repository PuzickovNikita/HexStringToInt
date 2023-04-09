package converter

import (
	"fmt"
	"math"
	"math/bits"
)

type Converter struct {
	hexStep      int   // Шаг при разборе шестнадцатеричного числа
	hexShift     int   // Число, при умножении сдвигающее на hexStep
	decStep      int   // Десятичный шаг
	decShift     int   // Число, при умножении сдвигающее на decSte
	preDecString []int // Разбиение переведенного числа в слайсы по decStep цифр
}

func GetConverter() (*Converter, error) {
	var c Converter
	switch bits.UintSize {
	case 64:
		c.hexStep = 7
	case 32:
		c.hexStep = 3
	case 16:
		c.hexStep = 1
	default:
		return nil, fmt.Errorf("strange system int format")
	}
	c.hexShift = 1 << (4 * c.hexStep)
	c.decStep = len(fmt.Sprintf("%d", c.hexShift))
	c.decShift = int(math.Pow10(c.decStep))
	return &c, nil
}

func (c *Converter) Clean() {
	c.preDecString = nil
}
