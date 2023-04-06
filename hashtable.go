package main

import (
	"errors"
)

const (
	initialCap    = 10
	maxLoadFactor = 0.75
	minLoadFactor = 0.25
)

type HashTable struct {
	data        []*string
	cap         int
	size        int
	loadFactor  float64
	minCapacity int
}

func New() *HashTable {
	return &HashTable{
		data:        make([]*string, initialCap),
		cap:         initialCap,
		size:        0,
		loadFactor:  0,
		minCapacity: initialCap,
	}
}

func (ht *HashTable) hash(s string) uint64 {
	const prime = 1099511628211
	var hash uint64 = 14695981039346656037

	for i := 0; i < len(s); i++ {
		hash = hash ^ uint64(s[i])
		hash = hash * prime
	}

	return hash
}

func (ht *HashTable) Insert(s string) {
	if ht.loadFactor >= maxLoadFactor {
		ht.resize(true)
	}

	h := ht.hash(s)
	index := h % uint64(ht.cap)
	for ; ht.data[index] != nil && *ht.data[index] != s; index = (index + 1) % uint64(ht.cap) {
	}
	if ht.data[index] == nil {
		ht.size++
	}
	ht.data[index] = &s
	ht.loadFactor = float64(ht.size) / float64(ht.cap)

	if ht.size == ht.minCapacity && ht.cap > ht.minCapacity {
		ht.resize(false)
	}
}

func (ht *HashTable) Get(s string) (string, error) {
	h := ht.hash(s)
	index := h % uint64(ht.cap)
	for ; ht.data[index] != nil && *ht.data[index] != s; index = (index + 1) % uint64(ht.cap) {
	}
	if ht.data[index] == nil {
		return "", errors.New("value not found")
	}
	return *ht.data[index], nil
}

func (ht *HashTable) Delete(s string) error {
	h := ht.hash(s)
	index := h % uint64(ht.cap)
	for ; ht.data[index] != nil && *ht.data[index] != s; index = (index + 1) % uint64(ht.cap) {
	}
	if ht.data[index] != nil {
		ht.data[index] = nil
		ht.size--
		ht.loadFactor = float64(ht.size) / float64(ht.cap)

		if ht.loadFactor < minLoadFactor && ht.cap > ht.minCapacity {
			ht.resize(false)
		}
		return nil
	}
	return nil

}

func (ht *HashTable) resize(upscale bool) {
	oldData := ht.data
	// oldCap := ht.cap
	if upscale {
		ht.cap = ht.cap * 2
	} else {
		ht.cap = ht.cap / 2
	}
	ht.data = make([]*string, ht.cap)
	ht.size = 0
	for _, v := range oldData {
		if v != nil {
			h := ht.hash(*v)
			index := h % uint64(ht.cap)
			for ; ht.data[index] != nil && *ht.data[index] != *v; index = (index + 1) % uint64(ht.cap) {
			}
			if ht.data[index] == nil {
				ht.size++
			}
			ht.data[index] = v
		}
	}
	ht.loadFactor = float64(ht.size) / float64(ht.cap)
	ht.minCapacity = int(float64(ht.cap) * minLoadFactor)
}
