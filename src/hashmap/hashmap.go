package hashmap

import (
	"github.com/KrzysztofHeinke/DataStructures/src/linkedlist"
)

type hashMap struct {
	hashArray []*linkedlist.List
	capacity  int
}

type hashData struct {
	Key   interface{}
	Value interface{}
}

// NewHashMap creats HashMap
func NewHashMap(capacity int) *hashMap {
	return &hashMap{
		hashArray: make([]*linkedlist.List, capacity),
		capacity:  capacity,
	}
}

// hash function to calculate index
func (hm *hashMap) hash(v interface{}) int {
	hash := 0
	switch v {
	case v.(string):
		for _, c := range v.(string) {
			hash += int(c)
		}
	}
	for {
		hash = hash / hm.capacity
		if hash < hm.capacity {
			break
		}
	}
	return hash
}

func (hm *hashMap) Insert(key, value interface{}) {
	pos := hm.hash(key)
	if hm.hashArray[pos] == nil {
		hm.hashArray[pos] = linkedlist.NewList()
	}
	n := hm.Search(key)
	if n == nil {
		hm.hashArray[pos].Insert(hashData{Key: key, Value: value})
	} else {
		n.Value = hashData{Key: key, Value: value}
	}
}

func (hm *hashMap) Delete(key interface{}) {
	pos := hm.hash(key)
	if hm.hashArray[pos] != nil {
		hm.hashArray[pos].Delete(key)
	}
}

func (hm *hashMap) Search(key interface{}) *linkedlist.Node {
	pos := hm.hash(key)
	node := hm.hashArray[pos].GetHead()
	if hm.hashArray[pos] != nil {
		for {
			if node == nil {
				return nil
			}
			if node.Value.(hashData).Key == key {
				break
			}
			node = node.Next
		}
	}
	return node
}

func (hm *hashMap) Print() {
	for _, elem := range hm.hashArray {
		if elem != nil {
			elem.Print()
		}
	}
}
