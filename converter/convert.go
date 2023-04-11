package converter

import (
	"strconv"
	"strings"
)

func (c *Converter) Convert(hexString string) (string, error) {
	//TODO: по идее можно заранее спрогнозировать длину с +- норм точностью
	c.Clean()
	defer c.Clean()

	step := len(hexString) % c.hexStep
	for i := 0; i < len(hexString); i += step {
		if (step != c.hexStep) && (i == step) {
			step = c.hexStep
		}
		//преобразуем следующие
		val, err := strconv.ParseInt(hexString[i:i+step], 16, 64)
		if err != nil {
			return "", err
		}

		//умножаем "число" на 16^hexStep + int(val)
		carry := int(val)

		for it := c.Storage.Begin(); it.InRange(); it.Next() {
			tmp := it.Get()*c.hexShift + carry
			if tmp >= c.decShift {
				carry = tmp / c.decShift
				tmp = tmp % c.decShift
			} else {
				carry = 0
			}
			it.Set(tmp)
		}
		if carry != 0 {
			c.Storage.PushBack(carry)
		}
	}

	var sb strings.Builder
	for it := c.Storage.End(); it.InRange(); it.Prev() {
		tmpString := strconv.Itoa(it.Get())
		//проверяем надо ли добавить лидирующие нули
		if (len(tmpString) < c.decStep) && (!c.Storage.End().IsEqual(&it)) {
			sb.WriteString(strings.Repeat("0", c.decStep-len(tmpString)))
		}
		sb.WriteString(tmpString)
		//sb.WriteString(".")
	}
	return sb.String(), nil
}
