package main

import (
	"fmt"

	"github.com/KrzysztofHeinke/DataStructures/src/hashmap"
)

func main() {
	hm := hashmap.NewHashMap(10)
	hm.Insert("ERCI", 10)
	hm.Insert("ERIC", 11)
	hm.Delete("ERIC")
	hm.Print()
	fmt.Println(hm)
}
