package main

import "fmt"

func main() {

	fmt.Println("Pointer")

	j, i := 24, 1024
	p := &j
	fmt.Println(*p, i)

	*p = 21
	fmt.Println(j)
	fmt.Println(*p, j)
}
