package main

import "fmt"

func BarisBilanganGanjil(n int) {
	if n == 1 {
		fmt.Printf("%d ", n)
	} else {
		BarisBilanganGanjil(n - 1)
		if n%2 != 0 {
			fmt.Printf("%d ", n)
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	BarisBilanganGanjil(n)
}
