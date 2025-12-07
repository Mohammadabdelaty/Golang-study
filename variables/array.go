package main

import "fmt"

func main() {
	fmt.Println("Array:")
	var o [5]int = [5]int{1, 2, 3, 4, 5} // array of five integers
	var p [3]string = [3]string{"Hi",",", "You!"} 
	var s [4]bool = [4]bool{true, false}  // 1 true 3 flases                 // array of three strings (default zero values)
	// p[0] = "Hello"
	// p[1] = "World"
	// p[2] = "Go"
	fmt.Println("o=", o)
	fmt.Println("p=", p)
	fmt.Println("s=", s)
}