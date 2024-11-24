package main

import "fmt"

func CetakPola(n, i int) {
	if n > 0 {
		CetakPola(n-1, i)
		CetakPola(n, i+1)
		fmt.Print("*")
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	CetakPola(n, 1)
}
