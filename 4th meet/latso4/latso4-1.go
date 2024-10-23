package main

import "fmt"

func main() {
	var x, i, y int
	fmt.Scan(&x)
	for i = 1; i <= x; i++ {
		y = y + i
	}
	fmt.Print(y)
}
