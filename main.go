package main

import (
	"saputrakhotim-technical-test/soalDua"
	"saputrakhotim-technical-test/soalSatu"
)

func main() {
	// inputStr string
	inputStr := "We Always Mekar"
	soalSatu.Jawab(inputStr)
	inputStr = "Coding is fun"
	soalSatu.Jawab(inputStr)

	// inputStr arrayString
	inputArr := []string{"Abc", "bCd"}
	soalDua.Jawab(inputArr)

}
