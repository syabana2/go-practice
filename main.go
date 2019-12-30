package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("Hello World. Sqrt(2) = %v\n", math.Sqrt(2))

	z := 0
	for i := 0; i < 10; i++ {
		z += i
	}

	fmt.Println(z)

}
