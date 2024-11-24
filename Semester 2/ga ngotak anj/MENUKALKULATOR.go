package main

import "fmt"

func menu() {
	fmt.Println("-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------")
}

func hitungJumlah() {
	var b1, b2 int
	fmt.Print("Masukan dua bilangan yang akan dijumlahkan : ")
	fmt.Scanln(&b1, &b2)
	fmt.Printf("Hasil penjumlahan : %d\n", b1+b2)
}

func hitungKali() {
	var b1, b2 int
	fmt.Print("Masukan dua bilangan yang akan dikalikan : ")
	fmt.Scanln(&b1, &b2)
	fmt.Printf("Hasil perkalian : %d\n", b1*b2)
}

func hitungBagi() {
	var b1, b2 float64
	fmt.Print("Masukan dua bilangan yang akan dibagikan : ")
	fmt.Scanln(&b1, &b2)
	fmt.Printf("Hasil pembagian : %.1f\n", b1/b2)
}

func main() {
	var pilih int
	for pilih != 4 {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scanln(&pilih)
		if pilih == 1 {
			hitungJumlah()
		} else if pilih == 2 {
			hitungKali()
		} else if pilih == 3 {
			hitungBagi()
		}
	}
}
