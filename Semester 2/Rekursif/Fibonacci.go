package main

import "fmt"

func fibo(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	var n, hasil int
	fmt.Scanln(&n)
	hasil = fibo(n)
	fmt.Print(hasil)
}
