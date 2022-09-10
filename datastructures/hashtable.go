package datastructures

type HashTable[K comparable, V any] struct {
	table     []V
	tableSize int
	hashFunc  func(Key K, Size int) int
}

func (table *HashTable[K, V]) Set(Key K, Value V) {
	index := table.hashFunc(Key, table.tableSize)
	table.table[index] = Value
}

func (table *HashTable[K, V]) Get(Key K) V {
  index := table.hashFunc(Key, table.tableSize)
  return table.table[index]
}

func NewHashTable[Key comparable, Value any](Size int, hashFunc func(Key Key, Size int) int) *HashTable[Key, Value] {
	table := &HashTable[Key, Value]{
		table:     make([]Value, Size),
		tableSize: Size,
    hashFunc: hashFunc,
	}
	return table
}
