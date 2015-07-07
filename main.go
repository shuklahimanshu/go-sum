package main

import (
	"fmt"

	"github.com/6amedev/go_sum/sum"
)

func main() {
	fmt.Println("Result of 5 + 4:", sum.Add(4, 5))
	fmt.Println("Result of 4 - 5:", sum.Sub(4, 5))
}
