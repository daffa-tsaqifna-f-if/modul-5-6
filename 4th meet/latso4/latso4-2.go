package main

import "fmt"

func main() {
	var r, phi, t, l, i, x float64
	phi = 3.14159265359
	fmt.Scan(&x)
	for i = 0; i < x; i++ {
		fmt.Scan(&r, &t)
		l = (phi * r * r * t) / 3
		fmt.Print(l)
	}
}
