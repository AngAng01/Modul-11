# MODUL SEBELAS
  ## SOAL 1
  Program di atas adalah program untuk menghitung suara dalam sebuah pemungutan suara dengan beberapa aturan. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Pendefinisian Tipe Data Suara, Tipe data array tetap dengan panjang maksimum 1000 elemen, digunakan untuk menyimpan suara yang diterima dalam pemungutan suara.
      - Fungsi hitungSuara, Fungsi utama yang melakukan perhitungan suara. Fungsi ini menghitung total suara masuk, jumlah suara sah, jumlah suara per kandidat, serta menghasilkan daftar kandidat yang diurutkan secara ascending menggunakan algoritma bubble sort.
      - Validasi Input, Logika di dalam fungsi hitungSuara memeriksa apakah suara termasuk dalam rentang valid (1–20). Suara di luar rentang ini dianggap tidak sah dan tidak dihitung.

      
   ## Code Explanation
   ```go
   type Suara [1000]int
   ```
   Kode di atas adalah kode untuk mendefinisikan tipe data baru yang disebut Suara, yang merupakan sebuah array dengan panjang tetap 1000 elemen bertipe int. 
  
   ```go
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
   ```
   Kode di atas adalah sebuah fungsi bernama hitungSuara yang menerima dua parameter: input berupa array Suara dan panjang yang menunjukkan jumlah elemen suara yang dimasukkan. Fungsi ini berfungsi untuk menghitung total suara, jumlah suara sah, serta menghitung jumlah suara yang diterima oleh masing-masing kandidat, dengan membatasi suara yang sah dalam rentang 1 hingga 20. Jika suara yang dimasukkan adalah 0, proses penghitungan akan dihentikan. Fungsi ini juga mengurutkan kandidat berdasarkan nomor mereka menggunakan algoritma bubble sort. Setelah itu, fungsi mengembalikan empat nilai: total suara masuk, total suara sah, peta jumlah suara per kandidat (map), dan daftar kandidat yang sudah diurutkan berdasarkan nomor secara ascending.

   ```go
	var suara Suara
   ```
   Kode di atas adalah deklarasi variabel suara yang bertipe Suara, sebuah array dengan kapasitas maksimal 1000 elemen bertipe int. 
   
   ```go
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
   ```
   Kode di atas adalah kode yang meminta input suara dari pengguna. Program mencetak pesan "Masukkan suara (akhiri dengan 0) :" dan menyimpan input suara dalam array suara dengan kapasitas maksimal 1000 elemen. Variabel panjang digunakan untuk menghitung jumlah suara yang dimasukkan. Dalam loop, program menerima input suara melalui fmt.Scan(&angka), menyimpannya dalam array, dan berhenti jika angka 0 dimasukkan.

   ```go
   total, sah, jumlah, urutan := hitungSuara(suara, panjang)
   ```
   Kode di atas adalah pemanggilan fungsi hitungSuara dengan parameter suara dan panjang, yang digunakan untuk menghitung total suara, jumlah suara sah, jumlah suara per kandidat, dan daftar kandidat yang diurutkan. 

   ```go
	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)
   ```
   Kode di atas adalah Kode yang mencetak total suara masuk dan jumlah suara sah menggunakan variabel total dan sah.
   
   ```go
	for i := 0; i < len(urutan); i++ {
		calon := urutan[i]
		fmt.Println(calon, ":", jumlah[calon])
	}
   ```
   Kode di atas adalah kode yang digunakan untuk mencetak daftar kandidat beserta jumlah suara yang diterima. Program mengiterasi array urutan dan untuk setiap kandidat (calon), mencetak nomor kandidat beserta jumlah suara yang diterima, yang diambil dari map jumlah.


   ## SOAL 2
  Program di atas adalah program untuk menghitung suara dalam pemungutan suara dan menentukan ketua dan wakil ketua berdasarkan suara yang diterima. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Tipe Data SuaraArray, Sebuah array dengan kapasitas maksimum 1000 elemen bertipe int untuk menyimpan suara yang diterima dalam pemungutan suara.
      - Tipe Data Kandidat, Struktur data yang digunakan untuk menyimpan informasi tentang kandidat, yaitu nomor kandidat (Nomor) dan jumlah suara yang diterima (Suara).
      - Fungsi hitungSuara, Fungsi yang digunakan untuk menghitung total suara, jumlah suara sah, serta jumlah suara yang diterima oleh masing-masing kandidat, dengan memvalidasi suara yang berada dalam rentang 1 hingga 20.

      
   ## Code Explanation
   ```go
   type SuaraArray [1000]int
   ```
   Kode di atas adalah kode untuk mendefinisikan tipe data baru yang disebut SuaraArray, yang merupakan sebuah array dengan panjang tetap 1000 elemen bertipe int. 
  
   ```go
	type Kandidat struct {
		Nomor int
		Suara int
	}
   ```
   Kode di atas adalah kode yang mendefinisikan struktur data Kandidat yang memiliki dua elemen: Nomor untuk menyimpan nomor kandidat dan Suara untuk menyimpan jumlah suara yang diterima. Struktur ini digunakan untuk menyimpan informasi kandidat dalam pemilihan.

   ```go
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
   ```
   Kode di atas adalah fungsi hitungSuara yang digunakan untuk menghitung total suara, jumlah suara sah, dan jumlah suara yang diterima oleh setiap kandidat. Fungsi ini menerima dua parameter: input yang merupakan array suara dan panjang yang menunjukkan jumlah suara yang dimasukkan. Fungsi ini memvalidasi suara, hanya menghitung suara yang berada dalam rentang 1 hingga 20 sebagai suara sah, dan menghitung jumlah suara untuk setiap kandidat menggunakan map. Fungsi ini kemudian mengembalikan total suara, jumlah suara sah, dan map jumlah suara untuk setiap kandidat.

   ```go
	var suara SuaraArray
   ```
   Kode di atas adalah deklarasi variabel suara yang bertipe SuaraArray, sebuah array dengan kapasitas maksimal 1000 elemen bertipe int. 
   
   ```go
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
   ```
   Kode di atas adalah kode yang meminta input suara dari pengguna dan menyimpannya dalam array suara. Program akan berhenti meminta input ketika angka 0 dimasukkan. Setiap suara selain 0 disimpan dalam array, dan variabel panjang menghitung jumlah suara yang diterima.

   ```go
	total, sah, jumlah := hitungSuara(suara, panjang)
   ```
   Kode di atas adalah Kode pemanggilan fungsi hitungSuara dengan parameter suara (array suara yang dimasukkan) dan panjang (jumlah suara yang diterima).
   
   ```go
	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)
   ```
   Kode di atas adalah Kode yang mencetak total suara masuk dan jumlah suara sah menggunakan variabel total dan sah.

   ```go
	var ketua, wakil Kandidat
   ```
   Kode di atas adalah Kode deklarasi dua variabel bertipe Kandidat, yaitu ketua dan wakil.

   ```go
	for calon, suara := range jumlah {
		if suara > ketua.Suara || (suara == ketua.Suara && calon < ketua.Nomor) {
			wakil = ketua
			ketua = Kandidat{Nomor: calon, Suara: suara}
		} else if suara > wakil.Suara || (suara == wakil.Suara && calon < wakil.Nomor) {
			wakil = Kandidat{Nomor: calon, Suara: suara}
		}
	}
   ```
   Kode di atas adalah kode yang digunakan menentukan ketua dan wakil ketua berdasarkan suara terbanyak. Program mengiterasi map jumlah untuk membandingkan suara setiap kandidat. Jika suara kandidat lebih banyak atau sama dengan ketua dengan nomor lebih kecil, kandidat tersebut menjadi ketua baru dan yang lama menjadi wakil. Proses serupa dilakukan untuk menentukan wakil ketua. Ketua dan wakil ketua yang terpilih memiliki suara terbanyak, dengan prioritas nomor lebih kecil jika ada suara yang sama.

   ```go
	fmt.Println("Ketua RT:", ketua.Nomor)
	fmt.Println("Wakil Ketua RT:", wakil.Nomor)
   ```
   Kode di atas adalah Kode yang digunakan untuk mencetak hasil pemilihan ketua dan wakil ketua. 


   ## SOAL 3
  Program di atas adalah program untuk menghitung suara dalam sebuah pemungutan suara dengan beberapa aturan. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Pendefinisian Tipe Data Suara, Tipe data array tetap dengan panjang maksimum 1000 elemen, digunakan untuk menyimpan suara yang diterima dalam pemungutan suara.
      - Fungsi hitungSuara, Fungsi utama yang melakukan perhitungan suara. Fungsi ini menghitung total suara masuk, jumlah suara sah, jumlah suara per kandidat, serta menghasilkan daftar kandidat yang diurutkan secara ascending menggunakan algoritma bubble sort.
      - Validasi Input, Logika di dalam fungsi hitungSuara memeriksa apakah suara termasuk dalam rentang valid (1–20). Suara di luar rentang ini dianggap tidak sah dan tidak dihitung.

      
   ## Code Explanation
   ```go
   type Suara [1000]int
   ```
   Kode di atas adalah kode untuk mendefinisikan tipe data baru yang disebut Suara, yang merupakan sebuah array dengan panjang tetap 1000 elemen bertipe int. 
  
   ```go
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
   ```
   Kode di atas adalah sebuah fungsi bernama hitungSuara yang menerima dua parameter: input berupa array Suara dan panjang yang menunjukkan jumlah elemen suara yang dimasukkan. Fungsi ini berfungsi untuk menghitung total suara, jumlah suara sah, serta menghitung jumlah suara yang diterima oleh masing-masing kandidat, dengan membatasi suara yang sah dalam rentang 1 hingga 20. Jika suara yang dimasukkan adalah 0, proses penghitungan akan dihentikan. Fungsi ini juga mengurutkan kandidat berdasarkan nomor mereka menggunakan algoritma bubble sort. Setelah itu, fungsi mengembalikan empat nilai: total suara masuk, total suara sah, peta jumlah suara per kandidat (map), dan daftar kandidat yang sudah diurutkan berdasarkan nomor secara ascending.

   ```go
	var suara Suara
   ```
   Kode di atas adalah deklarasi variabel suara yang bertipe Suara, sebuah array dengan kapasitas maksimal 1000 elemen bertipe int. 
   
   ```go
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
   ```
   Kode di atas adalah kode yang meminta input suara dari pengguna. Program mencetak pesan "Masukkan suara (akhiri dengan 0) :" dan menyimpan input suara dalam array suara dengan kapasitas maksimal 1000 elemen. Variabel panjang digunakan untuk menghitung jumlah suara yang dimasukkan. Dalam loop, program menerima input suara melalui fmt.Scan(&angka), menyimpannya dalam array, dan berhenti jika angka 0 dimasukkan.

   ```go
   total, sah, jumlah, urutan := hitungSuara(suara, panjang)
   ```
   Kode di atas adalah pemanggilan fungsi hitungSuara dengan parameter suara dan panjang, yang digunakan untuk menghitung total suara, jumlah suara sah, jumlah suara per kandidat, dan daftar kandidat yang diurutkan. 

   ```go
	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)
   ```
   Kode di atas adalah Kode yang mencetak total suara masuk dan jumlah suara sah menggunakan variabel total dan sah.
   
   ```go
	for i := 0; i < len(urutan); i++ {
		calon := urutan[i]
		fmt.Println(calon, ":", jumlah[calon])
	}
   ```
   Kode di atas adalah kode yang digunakan untuk mencetak daftar kandidat beserta jumlah suara yang diterima. Program mengiterasi array urutan dan untuk setiap kandidat (calon), mencetak nomor kandidat beserta jumlah suara yang diterima, yang diambil dari map jumlah.

   

   
