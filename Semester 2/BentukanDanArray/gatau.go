package main

import "fmt"

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64
	fmt.Scan(&tujuan, &lama, &jumlah)
	perhitunganBiaya(jumlahMhs, lamaPerjalanan, tujuan, totalBiaya)
	if tujuan == "domestik" {
		fmt.Println()
	}
}

func tanggunganHari(jumlahHari int, tujuan string) integer {
	if tujuan == "domestik" {
		if jumlahHari < 3 {
			return jumlahHari
		} else {
			return 3
		}
	} else if tujuan == "mancanegara" {
		if jumlahHari < 8 {
			return jumlahHari
		} else {
			return 5
		}
	}
}
func biayaPerhari(jumlahMhs int) int {
	return tanggunganHari * (70.000 + 250.000 + 300.000) * jumlahMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *int) {
	var hariDitanggung, biaya int
	tanggunganHari(jumlahhari, tujuan)
	biayaPerhari(jumlahMhs)
	hariDitanggung = tanggunganHari(jumlahHari, tujuan)
	biaya = biayaPerhari(jumlahMhs)
}
