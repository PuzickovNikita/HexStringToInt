package converter

import (
	"HexStringToInt64/converter"
	"math"
	"math/bits"
	"testing"
)

func TestMultiply(t *testing.T) {
	conv, _ := converter.GetConverter()
	//TODO: добавить в ридми(а сначала создать его) объяснение этих формул
	tmp := math.Log(float64(int(1)<<(bits.UintSize-4*conv.GetHexStep())-5)) / math.Log(10)
	if float64(conv.GetDecStep()) >= tmp {
		t.Errorf("Возможно переполнение: decStep %d >= %f", conv.GetDecStep(), tmp)
	}

	tmp = math.Log(float64(int(1)<<(4*conv.GetHexStep()))) / math.Log(10)
	if float64(conv.GetDecStep()) <= tmp {
		t.Errorf("Выбранный шаг в 10ой системе не вмещает в себя максимальное число с выбранным шагом в 16ой системе:\ndecStep %d <= %f", conv.GetDecStep(), tmp)
	}
}
