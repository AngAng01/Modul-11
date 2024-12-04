package main

import "fmt"

type SuaraArray [1000]int

type Kandidat struct {
	Nomor int
	Suara int
}

func hitungSuara(input SuaraArray, panjang int) (int, int, map[int]int) {
	totalSuara := 0
	suaraSah := 0
	jumlahSuara := make(map[int]int)

	for i := 0; i < panjang; i++ {
		if input[i] == 0 {
			break
		}
		totalSuara++
		if input[i] >= 1 && input[i] <= 20 {
			suaraSah++
			jumlahSuara[input[i]]++
		}
	}

	return totalSuara, suaraSah, jumlahSuara
}

func main() {
	var suara SuaraArray
	fmt.Println("Masukkan suara ( akhiri dengan 0 ) :")
	panjang := 0

	for i := 0; i < len(suara); i++ {
		var angka int
		fmt.Scan(&angka)
		if angka == 0 {
			break
		}
		suara[i] = angka
		panjang++
	}

	total, sah, jumlah := hitungSuara(suara, panjang)

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)

	var ketua, wakil Kandidat
	for calon, suara := range jumlah {
		if suara > ketua.Suara || (suara == ketua.Suara && calon < ketua.Nomor) {
			wakil = ketua
			ketua = Kandidat{Nomor: calon, Suara: suara}
		} else if suara > wakil.Suara || (suara == wakil.Suara && calon < wakil.Nomor) {
			wakil = Kandidat{Nomor: calon, Suara: suara}
		}
	}

	fmt.Println("Ketua RT:", ketua.Nomor)
	fmt.Println("Wakil Ketua RT:", wakil.Nomor)
}
