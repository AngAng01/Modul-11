package main

import "fmt"

type Suara [1000]int

func hitungSuara(input Suara, panjang int) (int, int, map[int]int, []int) {
	totalSuara := 0
	suaraSah := 0
	jumlahSuara := make(map[int]int)
	var urutan []int

	for i := 0; i < panjang; i++ {
		if input[i] == 0 {
			break
		}
		totalSuara++
		if input[i] >= 1 && input[i] <= 20 {
			suaraSah++
			if jumlahSuara[input[i]] == 0 {
				urutan = append(urutan, input[i])
			}
			jumlahSuara[input[i]]++
		}
	}

	for i := 0; i < len(urutan)-1; i++ {
		for j := i + 1; j < len(urutan); j++ {
			if urutan[i] > urutan[j] {
				temp := urutan[i]
				urutan[i] = urutan[j]
				urutan[j] = temp
			}
		}
	}

	return totalSuara, suaraSah, jumlahSuara, urutan
}

func main() {
	fmt.Println("Masukkan suara (akhiri dengan 0) :")
	var suara Suara
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

	total, sah, jumlah, urutan := hitungSuara(suara, panjang)

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)

	for i := 0; i < len(urutan); i++ {
		calon := urutan[i]
		fmt.Println(calon, ":", jumlah[calon])
	}
}
