package main

func main() {
	ht := New()
	ht.Insert("Name", "KittyBorgX")
	ht.Insert("Age", "16")
	ht.Get("Name") // Returns "KittyBorgX"
	ht.Delete("Name")
	ht.Get("Name") // Gives en error
	ht.Get("Age2") // Returns an error
}
