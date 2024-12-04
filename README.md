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
      - 

      
   ## Code Explanation
   ```go
   type arrBerat [1000]float64
   ```
   Kode di atas adalah deklarasi tipe data baru berupa array dengan kapasitas 1000 elemen bertipe float64.
  
   ```go
  func BeratTerkecil(arr arrBerat, n int) float64 {
  	min := arr[0]
  	for i := 1; i < n; i++ {
  		if arr[i] < min {
  			min = arr[i]
  		}
  	}
  	return min
  }
   ```
   Kode di atas adalah fungsi BeratTerkecil yang mencari nilai terkecil dalam array arrBerat. Fungsi ini menerima dua parameter, yaitu array arr dan integer n, yang menunjukkan jumlah elemen. Fungsi membandingkan elemen-elemen array dan mengembalikan nilai terkecil yang ditemukan.

   ```go
  func BeratTerbesar(arr arrBerat, n int) float64 {
  	max := arr[0]
  	for i := 1; i < n; i++ {
  		if arr[i] > max {
  			max = arr[i]
  		}
  	}
  	return max
  }
   ```
   Kode di atas adalah fungsi BeratTerbesar yang digunakan untuk mencari nilai terbesar dalam array arrBerat. Fungsi ini menerima dua parameter, yaitu array arr dan integer n yang menunjukkan jumlah elemen. Fungsi membandingkan elemen-elemen array dan mengembalikan nilai terbesar yang ditemukan.

   ```go
   var N int
   ```
   Kode diatas adalah deklarasi variabel yang menyatakan bahwa variabel N bertipe int (integer) dan digunakan untuk menyimpan nilai bilangan bulat. 

   ```go
   var berat arrBerat
   ```
   Kode di atas adalah deklarasi variabel yang menyatakan bahwa variabel berat bertipe arrBerat, yaitu array dengan kapasitas 1000 elemen bertipe float64

   ```go
	fmt.Print("Masukkan jumlah anak kelinci (N) : ")
	fmt.Scan(&N)
   ```
   Kode di atas adalah kode untuk memberi pesan dan meminta input dari pengguna dan menyimpannya dalam variabel N.

   ```go
	fmt.Printf("Masukkan berat %d anak kelinci :\n", N)
	for i := 1; i <= N; i++ {
		fmt.Printf("Kelinci ke %d : ", i)
		fmt.Scan(&berat[i])
	}
   ```
   Kode di atas adalah kode untuk meminta pengguna memasukkan berat anak kelinci. Pesan pertama menampilkan jumlah anak kelinci yang harus dimasukkan, dan kemudian menggunakan loop for untuk meminta input berat setiap kelinci satu per satu, menyimpannya dalam array berat.

   ```go
	min := BeratTerkecil(berat, N)
	max := BeratTerbesar(berat, N)
   ```
   Kode di atas memanggil dua fungsi, BeratTerkecil dan BeratTerbesar, untuk mencari nilai terkecil dan terbesar dalam array berat dengan jumlah elemen N. Hasilnya disimpan dalam variabel min dan max yang masing-masing berisi berat terkecil dan terbesar.

   ```go
	fmt.Printf("Berat terkecil : %.2f\n", min)
	fmt.Printf("Berat terbesar : %.2f\n", max)
   ```
   Kode di atas digunakan untuk menampilkan hasil berat terkecil dan terbesar yang sudah dihitung sebelumnya. fmt.Printf("Berat terkecil : %.2f\n", min) mencetak berat terkecil dengan dua angka desimal, sedangkan fmt.Printf("Berat terbesar : %.2f\n", max) mencetak berat terbesar dengan dua angka desimal.
