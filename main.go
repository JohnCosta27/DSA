package main

import (
	"fmt"
	"hash/fnv"

	"johncosta.tech/DSA/datastructures"
)

func main() {
	fmt.Println("Data Structures and Algorithms")
	hashTable := datastructures.NewHashTable[string, int](17,
		func(Key string, Size int) int {
			h := fnv.New32()
			h.Write([]byte(Key))
			return int(h.Sum32()) % Size
		})
	fmt.Println(hashTable)
}
