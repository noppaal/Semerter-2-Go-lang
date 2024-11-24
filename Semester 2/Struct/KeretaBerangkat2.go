package main

import "fmt"

type Waktu struct {
	Jam   int
	Menit int
}

type JadwalKereta struct {
	NomorKA      string
	KotaAsal     string
	JamBerangkat Waktu
	KotaTujuan   string
	JamTiba      Waktu
}

func main() {
	var jadwal1, jadwal2, jadwal3 JadwalKereta
	fmt.Scanln(&jadwal1.NomorKA, &jadwal1.KotaAsal, &jadwal1.JamBerangkat.Jam, &jadwal1.JamBerangkat.Menit, &jadwal1.KotaTujuan, &jadwal1.JamTiba.Jam, &jadwal1.JamTiba.Menit)
	fmt.Scanln(&jadwal2.NomorKA, &jadwal2.KotaAsal, &jadwal2.JamBerangkat.Jam, &jadwal2.JamBerangkat.Menit, &jadwal2.KotaTujuan, &jadwal2.JamTiba.Jam, &jadwal2.JamTiba.Menit)
	fmt.Scanln(&jadwal3.NomorKA, &jadwal3.KotaAsal, &jadwal3.JamBerangkat.Jam, &jadwal3.JamBerangkat.Menit, &jadwal3.KotaTujuan, &jadwal3.JamTiba.Jam, &jadwal3.JamTiba.Menit)

	jadwalTerlama := jadwal1
	if jadwal2.JamTiba.Jam > jadwalTerlama.JamTiba.Jam ||
		(jadwal2.JamTiba.Jam == jadwalTerlama.JamTiba.Jam && jadwal2.JamTiba.Menit > jadwalTerlama.JamTiba.Menit) {
		jadwalTerlama = jadwal2
	}
	if jadwal3.JamTiba.Jam > jadwalTerlama.JamTiba.Jam ||
		(jadwal3.JamTiba.Jam == jadwalTerlama.JamTiba.Jam && jadwal3.JamTiba.Menit > jadwalTerlama.JamTiba.Menit) {
		jadwalTerlama = jadwal3
	}

	fmt.Print(jadwalTerlama.NomorKA)
}
