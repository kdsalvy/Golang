package basics

import "fmt"

const globalPI = 3.14

// Demo is exported function
func Demo() {
	fmt.Println("Hello From A Module!")

	// different ways to declare variables
	var i int
	i = 42
	fmt.Println(i)

	var f float32
	f = 3.14
	fmt.Println(f)

	firstName := "Saurabh"
	fmt.Println(firstName)

	// using pointer to reference variable
	var name = new(string)
	*name = "Saurabh"
	fmt.Println(*name)

	// getting pointer value from variable
	name2 := "Saurabh"
	ptr := &name2
	// print memory address and value
	fmt.Println(ptr, *ptr)

	// declaring constants
	const pi = 3.14
	fmt.Println("Constant ", pi)

	// global constant
	fmt.Println("Global Constant ", globalPI)
}
