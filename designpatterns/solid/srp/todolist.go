package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

var totalEntries = 0

type ToDoList struct {
	entries []string
}

func (t *ToDoList) String() string{
	return  strings.Join(t.entries, " ")
}

func (t *ToDoList) addEntry(e string) {
	totalEntries++
	t.entries = append(t.entries, e)
}

func (t *ToDoList) printList() {
	for i, v := range t.entries {
		fmt.Printf("%3d. %s\n", i, v)
	}
}

type Persistence struct{
}

func (p *Persistence) saveToFile(t *ToDoList, filename string){
	ioutil.WriteFile(filename, []byte(strings.Join(t.entries, "\n")), 0644)
}

func main() {
	t := ToDoList{} 
	t.addEntry("Milk")
	t.addEntry("Soya")
	t.addEntry("Egg")
	t.addEntry("Bread")
	t.addEntry("Chocolate")

	t.printList()

	p := Persistence{}

	p.saveToFile(&t, "todo.dat")
}
