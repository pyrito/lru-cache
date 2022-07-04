package main

import (
	lrucache "cache/lru"
	"fmt"
)

func main() {
	lru, err := lrucache.NewLRUCache(5)
	if err != nil {
		fmt.Println("Looks like we have a problem with NewLRUCache!")
	}

	lru.Add("lalala", 42)
	val, err := lru.Get("lalala")
	if err != nil {
		fmt.Println("Looks like the value didn't exist?")
	}
	fmt.Println(val)
	fmt.Println("blah blah")

}
