package main

import (
	"fmt"
)

var romanMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func theroman_translator(str string) int {
	retVal := 0

	if len(str) <= 0 {
		return 0
	}

	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && romanMap[str[i]] >= romanMap[str[i+1]] {
			//Kind of MC, VI etc.
			retVal += romanMap[str[i]]
		}else if i+1 < len(str) && romanMap[str[i]] < romanMap[str[i+1]] {
			//Kind of IV, CM etc.
			retVal += (romanMap[str[i+1]] -  romanMap[str[i]])
			i++
		}else{
			//does not have anything onto it's right
			retVal += romanMap[str[i]]
		}
	}

	return retVal
}

func main() {
	//rstr := "III"
	//rstr := "IV"
	//rstr := "IX"
	//rstr := "LVIII"
	rstr := "MCMXCIV"
	fmt.Println("Input : ", rstr)
	iroman := theroman_translator(rstr)
	fmt.Println("Output : ", iroman)
}
