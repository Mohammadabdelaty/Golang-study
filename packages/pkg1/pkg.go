package pkg1

import (
	"fmt"
	"os"
)

/* 
This is public function will be used in the main.go file
because it starts with a capital letter, otherwise it won't
*/
func Another() {
	fmt.Println("This is another file.")
	fmt.Println("Current working directory:")
	dir, err := os.Getwd()
	fmt.Println(dir, err)
}


