package main

import "fmt"

func main() {
	var x, y, z, i int
	z = 1
	fmt.Scan(&x, &y)
	for i = 0; i < y; i++ {
		z = z * x
	}
	fmt.Print(z)
}
