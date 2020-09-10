package main

import (
	"fmt"
)

var hashMap = map[int]int{}

type elements struct {
	first  int
	second int
}

func sumof2(arr []int, sum int) []elements {
	result := []elements{}
	if len(arr) <= 0 {
		return result
	}

	for i := 0; i < len(arr)-1; i++ {
		if val, ok := hashMap[arr[i]]; ok {
			e := elements{
				first: arr[val],
				second: arr[i],
			}
			result = append(result, e)
		} else {
			hashMap[sum-arr[i]] = i
		}
	}

	return result
}

func main() {
	sum := 2
	arr := []int{-1, 0, 1, 2, -1, 4}
	result := sumof2(arr, sum)
	for _, f := range result {
		fmt.Printf("(%4d,%4d) ", f.first, f.second)
	}
	fmt.Println()
}
