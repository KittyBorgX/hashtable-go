package main

import (
	"math/rand"
	"testing"
)

func TestInsertAndGet(t *testing.T) {
	ht := New()

	// insert values
	ht.Insert("apple")
	ht.Insert("banana")
	ht.Insert("cherry")
	ht.Insert("date")

	// check if all values can be retrieved correctly
	value, err := ht.Get("apple")
	if err != nil || value != "apple" {
		t.Errorf("failed to retrieve the correct value for 'apple'")
	}
	value, err = ht.Get("banana")
	if err != nil || value != "banana" {
		t.Errorf("failed to retrieve the correct value for 'banana'")
	}
	value, err = ht.Get("cherry")
	if err != nil || value != "cherry" {
		t.Errorf("failed to retrieve the correct value for 'cherry'")
	}
	value, err = ht.Get("date")
	if err != nil || value != "date" {
		t.Errorf("failed to retrieve the correct value for 'date'")
	}
}

func TestDelete(t *testing.T) {
	ht := New()

	// insert values
	ht.Insert("apple")
	ht.Insert("banana")
	ht.Insert("cherry")
	ht.Insert("date")

	// delete value
	err := ht.Delete("banana")
	if err != nil {
		t.Errorf("failed with error 'value not found'")
	}

	// check if value was deleted
	_, err = ht.Get("banana")
	if err == nil {
		t.Errorf("value 'banana' was not deleted")
	}
}

func TestHashCollisions(t *testing.T) {
	ht := New()

	// generate 1000 random strings of length 10 characters
	keys := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		keys[i] = generateRandomString(10)
		ht.Insert(keys[i])
	}

	// check for hash collisions
	for i := 0; i < 999; i++ {
		for j := i + 1; j < 1000; j++ {
			if ht.hash(keys[i]) == ht.hash(keys[j]) {
				t.Errorf("hash collision detected: %s and %s have the same hash value", keys[i], keys[j])
			}
		}
	}
}

func Test8k(t *testing.T) {
	ht := New()

	// generate 8000 random strings of length 8192 characters
	keys := make([]string, 8000)
	for i := 0; i < 8000; i++ {
		keys[i] = generateRandomString(8192)
		ht.Insert(keys[i])
	}

	// check if all values can be retrieved correctly
	for i := 0; i < 8000; i++ {
		value, err := ht.Get(keys[i])
		if err != nil {
			t.Errorf("failed with error 'value not found'")
		}
		if value != keys[i] {
			t.Errorf("failed to retrieve the correct value")
		}
	}
}

func TestStress(t *testing.T) {
	ht := New()

	// generate 8000 random strings of length 8192 characters
	keys := make([]string, 8000)
	for i := 0; i < 8000; i++ {
		keys[i] = generateRandomString(8192)
		ht.Insert(keys[i])
	}

	// check if all values can be retrieved correctly
	for i := 0; i < 8000; i++ {
		value, err := ht.Get(keys[i])
		if err != nil {
			t.Errorf("failed with error 'value not found'")
		}
		if value != keys[i] {
			t.Errorf("failed to retrieve the correct value")
		}
	}

	// delete half of the values and check load factor
	for i := 0; i < 4000; i++ {
		ht.Delete(keys[i])
	}
	maxLoadFactor := 0.75
	loadFactor := ht.loadFactor
	if loadFactor > maxLoadFactor {
		t.Errorf("load factor (%f) is greater than maximum allowed value (%f)", loadFactor, maxLoadFactor)
	}

	// add 4000 more values and check if they can be retrieved correctly
	for i := 4000; i < 8000; i++ {
		ht.Insert(keys[i])
	}
	for i := 4000; i < 8000; i++ {
		value, err := ht.Get(keys[i])
		if err != nil {
			t.Errorf("failed with error 'value not found'")
		}
		if value != keys[i] {
			t.Errorf("failed to retrieve the correct value")
		}
	}

	// delete all values and check load factor
	for i := 0; i < 8000; i++ {
		ht.Delete(keys[i])
	}
	// loadFactor = ht.lo`adFactor

}

func generateRandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
