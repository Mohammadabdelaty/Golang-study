package main

import "fmt"

func main() {
	// Using var
	var ptr1 *int
	fmt.Println("1. Using var to declare a pointer:", ptr1)

	// Pointer to an Existing Variable
	var value int = 42
	ptr2 := &value
	fmt.Println("2. Pointer to an Existing Variable:", *ptr2)
	fmt.Println("3. Memory address of Variable:", ptr2)

	// Creating a Pointer with new
	ptr3 := new(int)
	*ptr3 = 100
	fmt.Println("4. Creating a Pointer with new:", *ptr3)
	fmt.Println("5. Memory address of new pointer:", ptr3)

	// Pointer to Pointer
	ptr4 := &ptr3
	fmt.Println("6. Pointer to Pointer:", **ptr4)
	fmt.Println("7. Memory address of Pointer:", *ptr4)
	fmt.Println("8. Memory address of Memory address of Pointer:", ptr4)
}