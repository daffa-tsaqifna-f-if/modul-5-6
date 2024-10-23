package main

import "fmt"

func main() {
	var x, y, i int
	fmt.Scan(&x)
	y = 1
	for i = 1; i <= x; i++ {
		y = y * i
	}
	fmt.Print(y)
}
