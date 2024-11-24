package main

import "fmt"

func Pemangkatan(x, y int) int {
	if y == 0 {
		return 1
	} else if y == 1 {
		return x
	} else {
		return x * Pemangkatan(x, y-1)
	}
}
func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Print(Pemangkatan(x, y))
}
