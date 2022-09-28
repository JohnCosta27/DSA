package main

import (
	"fmt"
	"hash/fnv"

	"johncosta.tech/DSA/algorithms"
	"johncosta.tech/DSA/datastructures"
)

func main() {
	fmt.Println("Data Structures and Algorithms")
  myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  algorithms.MergeSort(myArray[:])

}

func HashTable() {
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
