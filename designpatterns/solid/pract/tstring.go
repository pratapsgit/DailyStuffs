package main

import (
	"fmt"
	"strings"
)

type WordBook struct {
	words []string
}

func (w *WordBook) addWord(text string) {
	w.words = append(w.words, text)
}

func (w *WordBook) printWords() {
	fmt.Println(strings.Join(w.words, " "))
}

func main() {
	wb := WordBook{}

	wb.addWord("Hello")
	wb.addWord("Teacher")
	wb.addWord("Talk")
	wb.addWord("To")
	wb.addWord("My")
	wb.addWord("Father")

	wb.printWords()
}
