package converter

import (
	"strconv"
	"strings"
)

func (c *Converter) Convert(hexString string) (string, error) {
	//TODO: по идее можно заранее спрогнозировать длину с +- норм точностью
	var preDecString []int

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
		for i := range preDecString {
			preDecString[i] = preDecString[i]*c.hexShift + carry
			if preDecString[i] >= c.decShift {
				carry = preDecString[i] / c.decShift
				preDecString[i] = preDecString[i] % c.decShift
			} else {
				carry = 0
			}
		}
		if carry != 0 {
			preDecString = append(preDecString, carry)
		}
	}

	var sb strings.Builder
	for i := len(preDecString) - 1; i >= 0; i-- {
		tmpString := strconv.Itoa(preDecString[i])
		//проверяем надо ли добавить лидирующие нули
		if (len(tmpString) < c.decStep) && (i != (len(preDecString) - 1)) {
			sb.WriteString(strings.Repeat("0", c.decStep-len(tmpString)))
		}
		sb.WriteString(tmpString)
		//sb.WriteString(".")
	}
	return sb.String(), nil
}
