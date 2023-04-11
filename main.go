package main

import (
	"HexStringToInt64/converter"
	"HexStringToInt64/converter/storage/vectorStorage"
)

func main() {
	conv, err := converter.GetConverter(vectorStorage.NewVectorStorage(0))
	if err != nil {
		panic(err)
	}
	println(conv.GetHexStep(), conv.GetHexShift(), conv.GetDecStep(), conv.GetDecShift())
	str, _ := conv.Convert("363642789D3492527F")
	println(str)
	println(str == "1000034000023414002303")
}
