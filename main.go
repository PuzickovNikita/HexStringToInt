package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Шаг при разборе шестнадцатеричного числа
const hexStep = 7

// Число, при умножении сдвигающее на hexStep
const hexShift = 1 << (4 * hexStep)

// Десятичный шаг
var decStep int = len(fmt.Sprintf("%d", hexShift))

// Число, при умножении сдвигающее на decStep
var decShift int = int(math.Pow10(decStep))

func convert(hexString string) (string, error) {
	//TODO: по идее можно заранее спрогнозировать длинну с +- норм точностью
	var preDecString []int

	step := len(hexString) % hexStep
	for i := 0; i < len(hexString); i += step {
		if (step != hexStep) && (i == step) {
			step = hexStep
		}
		//преобразуем следующие
		val, err := strconv.ParseInt(hexString[i:i+step], 16, 64)
		if err != nil {
			return "", err
		}

		//умножаем "число" на 16^hexStep + int(val)
		carry := int(val)
		for i := range preDecString {
			preDecString[i] = preDecString[i]*hexShift + carry
			if preDecString[i] >= decShift {
				carry = preDecString[i] / decShift
				preDecString[i] = preDecString[i] % decShift
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
		if (len(tmpString) < decStep) && (i != (len(preDecString) - 1)) {
			sb.WriteString(strings.Repeat("0", decStep-len(tmpString)))
		}
		sb.WriteString(tmpString)
		//sb.WriteString(".")
	}
	return sb.String(), nil
}

func main() {
	println(hexStep, hexShift, hexShift*hexShift, decStep, decShift)
	str, _ := convert("2d8499a18139ceab5178cb039c029f9f008cde8656b914a0be9b76a7e9232dad3228d9cd2e954b28d10af418c9a95")
	println(str)
	println(str == "1710414223956606143409440747890222080251214130433827271746408001889495047732082895143275182701331674684724124309")
}
