package main

import(
	"fmt"
	"os"
)

func allUnique(s string) bool {
	ret := true
	charmap := map[byte]int{}

	strlen := len(s)
	for i:=0; i<strlen; i++{
		if _, ok := charmap[s[i]]; ok{
			ret = false
			break
		}else{
			charmap[s[i]] = i
		}
	}

	return ret
}

func findLongestSubString(s string) (string, int){
	strlen := len(s)
	maxlen := 0
	longestString := ""
	for i:=0; i<strlen; i++ {
		for j:=i+1; j<=strlen; j++ {
			isAllUnique := allUnique(s[i:j])
			//fmt.Printf("%15s : %v\n", s[i:j], isAllUnique)
			if isAllUnique == true && len(s[i:j]) > maxlen {
				maxlen = len(s[i:j])
				longestString = s[i:j]
			}
		}
	}

	return longestString, maxlen
}

func main(){
        inputString := ""

	if(len(os.Args)>1) {
		inputString = os.Args[1]
	}else{
		return;
	}

	fmt.Println("Input String ", inputString)
	s, l := findLongestSubString(inputString)

	fmt.Println("Longest string found is ", s, " has length ", l)
}
