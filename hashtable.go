package main

import "fmt"

// Constants of offset and Prime taken from
// https://cs.opensource.google/go/go/+/refs/tags/go1.20.2:src/hash/fnv/fnv.go;l=30
const map_size = 16
const fnvOffSetBasis uint64 = 14695981039346656037
const fnvPrime = 1099511628211
const loadFactorThreshold = 0.5

type Node struct {
	key   string
	value string
}

type HashTable struct {
	length int
	data   [][]Node
}

func to_bytes(str string) []byte {
	return []byte(str)
}

func hash(key string, limit int) int {
	hash := fnvOffSetBasis
	for _, b := range to_bytes(key) {
		hash = hash ^ uint64(b)
		hash = hash * fnvPrime
	}
	return int(hash % uint64(limit))
}

// The new function - Create a new hashtable
func New() HashTable {
	return HashTable{
		data: make([][]Node, map_size),
	}
}

func (ht *HashTable) loadFactor() float32 {
	return float32(ht.length) / float32(len(ht.data))
}

// Function to expand the hash table to avoid collisions
func (ht *HashTable) expandTable() error {
	newTable := make([][]Node, len(ht.data)*2)
	for _, data := range ht.data {
		for _, e := range data {
			newHash := hash(e.key, len(newTable))
			newTable[newHash] = append(newTable[newHash], Node{e.key, e.value})
		}
	}
	ht.data = newTable
	return nil
}

func (ht *HashTable) Insert(key string, value string) {
	hash := hash(key, len(ht.data))
	// Check if key is already added. If the key is there
	// then overwrite the present key.
	for i, e := range ht.data[hash] {
		if e.key == key {
			ht.data[hash][i].value = value
			return
		}
	}

	ht.data[hash] = append(ht.data[hash], Node{key, value})
	ht.length += 1
	if ht.loadFactor() > loadFactorThreshold {
		// If the table begins to fill up, then we expand the table.
		// We get to know if the table is filling up if the load factor increases
		// more than 0.5.
		ht.expandTable()
	}
}

// Get function which prints the value if its found or else
// prints an error message
func (ht *HashTable) Get(key string) {
	hash := hash(key, len(ht.data))
	for _, v := range ht.data[hash] {
		if v.key == key {
			fmt.Println("Value found: ", v.value)
			return
		}
	}
	fmt.Println("error: Could not find key: ", key)
}

// Delete function which also resizes the table
func (ht *HashTable) Delete(key string) {
	hash := hash(key, len(ht.data))
	for i, v := range ht.data[hash] {
		if v.key == key {
			current := ht.data[hash]
			current[i] = current[len(current)-1]
			current = current[:len(current)-1]
			ht.length -= 1
			ht.data[hash] = current
		}
	}
}

