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
	fmt.Println("blah blah")

}
