package main

import "fmt"

type pegawai struct {
	nama string
	masa_kerja, besar_bonus int
	gaji_pokok float64
}

func main () {
	var orang pegawai
	
	fmt.Scan(&orang.nama, &orang.gaji_pokok, &orang.masa_kerja)
	
	hitung_bonus(&orang)
	
	fmt.Printf("Pegawai dengan nama %v mendapatkan bonus sebesar Rp %v ", orang.nama, orang.besar_bonus)
}

func hitung_bonus(p *pegawai) {
	
	if p.masa_kerja >= 10 {
		p.besar_bonus = int(p.gaji_pokok * 1.5)
	} else if p.masa_kerja >= 5 && p.masa_kerja < 10 {
		p.besar_bonus = int(p.gaji_pokok * 0.75)
	} else if p.masa_kerja < 5 {
		p.besar_bonus = int(p.gaji_pokok * 0.5)
	}
}