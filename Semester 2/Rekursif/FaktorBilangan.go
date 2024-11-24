package main

import "fmt"

func FaktorBilangan(n, Nawal int) {
	if n == 1 {
		fmt.Printf("%d ", n)
	} else {
		FaktorBilangan(n-1, Nawal)
		if Nawal%n == 0 {
			fmt.Printf("%d ", n)
		}
	}
}

func main() {
	var n, Nawal int
	fmt.Scan(&n)
	Nawal = n
	FaktorBilangan(n, Nawal)
}
