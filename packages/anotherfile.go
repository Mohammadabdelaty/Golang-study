package main

import (
	"fmt"
	"os"
)

func another() {
	fmt.Println("This is another file.")
	fmt.Println("Current working directory:")
	dir, err := os.Getwd()
	fmt.Println(dir, err)
}