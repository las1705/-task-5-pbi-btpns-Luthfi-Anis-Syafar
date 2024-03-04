package main

import (
	"fmt"
	"time"
)

func tes() {
	currentTime := time.Now().UTC()

	// Format tahun, bulan, tanggal tanpa spasi
	formattedTime := currentTime.Format("20060102")

	fmt.Println("Waktu yang diformat:", formattedTime)
}
