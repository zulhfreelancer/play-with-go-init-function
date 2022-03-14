package main

import (
	"fmt"

	"init-with-testmain.io/controller"
)

func init() {
	fmt.Println("Hello from init")
}

func main() {
	fmt.Println(controller.Sum(5, 5))
}
