package main

import (
	"fmt"

	"github.com/KrzysztofHeinke/DataStructures/src/linkedlist"
)

func main() {
	l := linkedlist.List{}
	l.Insert(29)
	l.Insert(31)
	l.Insert(31)
	l.Insert(32)
	l.Print()
	l.Delete(32)
	l.Print()
	fmt.Println(l.Search(32))
}
