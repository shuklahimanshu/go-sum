package main

import "fmt"

func main() {
	fmt.Println("Addition of 5 + 4:", add(4, 5))
}

func add(a, b int) int {
	return a + b
}
