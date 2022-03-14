package main

import (
	"fmt"
)

func init() {
	fmt.Println("Hello from init")
}

func Sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(Sum(5, 5))
}
