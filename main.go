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

  hashTable.Set("Hello", 69)
  hashTable.Set("Hi", 420)
  fmt.Println(hashTable.Get("Hello"))
  fmt.Println(hashTable.Get("Hi"))
  fmt.Println(hashTable.Get("Not found"))
  fmt.Println(hashTable)
}
