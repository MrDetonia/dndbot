package main

// check creates a runtime error when neccesary from an error variable.
func check(e error) {
	if e != nil {
		panic(e)
	}
}
