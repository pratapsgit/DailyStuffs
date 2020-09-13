package main

import (
	"fmt"
	"math"
)

func myatof(str string) float32 {
	sign := 1
	retVal := 0
	precision := -1

	if len(str) <= 0 {
		return float32(retVal)
	}

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || str[i] == '+' {
			continue
		}

		if str[i] == '-' {
			sign *= -1
			continue
		}

		if str[i] == '.' {
			precision = 0
			continue
		}

		if str[i] >= 48 && str[i] < 58 {
			retVal = retVal*10 + int(str[i]-'0')
			fmt.Println(retVal)
			if precision >= 0 {
				precision++
			}
		}

	}

	if precision<0{
		precision = 0
	}

	return float32(retVal * sign) * float32(1 / math.Pow10(precision))
}

func myatoi(str string) int {
	sign := 1
	retVal := 0
	precision := -1

	if len(str) <= 0 {
		return retVal
	}

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || str[i] == '+' {
			continue
		}

		if str[i] == '-' {
			sign *= -1
			continue
		}

		if str[i] == '.' {
			precision = 0
			continue
		}

		if str[i] >= 48 && str[i] < 58 {
			retVal = retVal*10 + int(str[i]-'0')
			if precision >= 0 {
				precision++
			}
		}

	}

	if precision<0{
		precision = 0
	}

	return retVal * sign / int(math.Pow10(precision))
}


func main() {
	//str := "    +5363.434"
	//str := "    -5363.434"
	//str := "4193 with words"
	//str:= "words and 987"
	//str := "-91283472332"
	//str := "-"
	//str := "+"
	//str := "."
	str := "-91."
	fmt.Println("String passed ", str)
	rFloat := myatof(str)
	fmt.Println("Float received ", rFloat)
	rInt := myatoi(str)
	fmt.Println("Integer received ", rInt)
}
