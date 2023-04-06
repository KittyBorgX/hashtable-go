package main

import "fmt"

func main() {
	ht := New()
	ht.Insert("apples")
	fmt.Println(ht.Get("apples")) // returns "apples"
	fmt.Println(ht.Get("sdf")) // returns an error
	fmt.Println(ht.Delete("apples")) // delete the value "apples"
	fmt.Println(ht.Get("apples")) // returns an error
}
