package main

import "HexStringToInt64/converter"

func main() {
	conv, err := converter.GetConverter()
	if err != nil {
		panic(err)
	}
	println(conv.GetHexStep(), conv.GetHexShift(), conv.GetDecStep(), conv.GetDecShift())
	str, _ := conv.Convert("2d8499a18139ceab5178cb039c029f9f008cde8656b914a0be9b76a7e9232dad3228d9cd2e954b28d10af418c9a95")
	println(str)
	println(str == "1710414223956606143409440747890222080251214130433827271746408001889495047732082895143275182701331674684724124309")
}
