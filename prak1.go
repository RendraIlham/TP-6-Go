package main

import "fmt"

type waktu struct {
	jam, menit, detik int
}

type kendaraan struct {
	jke, no string
	jm, jk waktu
	bea, durasi int
}

func main () {
	var k01, k02, k03 kendaraan
	var jenis, nopol string
	var j1, m1, d1, j2, m2, d2 int

	fmt.Scan(&jenis, &nopol, &j1, &m1, &d1, &j2, &m2, &d2)

	k01 =
}

func kendaraan_baru(j, n string, a, b, c, d, e, f int) kendaraan {
	var k kendaraan

	k.jke = j
	k.no = n 
	k.jm.jam = a
	k.jm.menit = b 
	k.jm.detik = c 
	k.jk.jam = d 
	k.jk.menit = e 
	k.jk.detik = f 

	return k 
}

func durasi_parkir(k *kendaraan) {
	var tot_detik int

	*k.durasi = tot_detik / 3600
}

func bea_parkir(k *kendaraan) {
	if *k.jke == "mobil" {
		*k.bea += 5000 + ((durasi_parkir(&k)) * 2500)
		
	}
}