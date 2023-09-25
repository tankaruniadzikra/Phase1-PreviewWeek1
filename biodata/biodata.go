package main

import (
	"fmt"
	"os"
)

// Mendeklarasikan sebuah struktur
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var dataTeman = []Teman{
	{
		Nama:      "Marcus Rashford",
		Alamat:    "Jl. Manchester No. 12",
		Pekerjaan: "Pemain Sepakbola",
		Alasan:    "Switch karir",
	},
	{
		Nama:      "Mason Mount",
		Alamat:    "Jl. Ayani 2 No. 6",
		Pekerjaan: "Artis",
		Alasan:    "Go untuk analisis data.",
	},
	{
		Nama:      "Andre Onana",
		Alamat:    "Jl. Kebangkitan No. 9",
		Pekerjaan: "Sistem Administrator",
		Alasan:    "Ingin membangun sistem yang efisien.",
	},
}

func main() {
	// untuk mendapatkan argumen yang diberikan saat menjalankan program melalui command line (CLI).
	args := os.Args
	// args variabel yang merupakan slice (daftar) dari string.
	// Slice ini berisi semua argumen yang diberikan saat program dijalankan dari command line.
	// Indeks pertama (indeks 0) biasanya berisi nama program itu sendiri.

	if len(args) != 2 {
		fmt.Println("Cara penggunaan: go run biodata.go [nomor_absen]")
		return
	}

	// untuk mengambil nomor absen dari argumen yang diberikan saat menjalankan program, dan kemudian memeriksa apakah nomor absen tersebut valid.
	nomorAbsen := args[1]
	absen, err := parseAbsen(nomorAbsen)
	if err != nil {
		fmt.Println(err)
		return
	}

	// untuk memastikan bahwa nomor absen yang diberikan saat menjalankan program berada dalam rentang yang valid (antara 1 dan jumlah teman yang ada).
	if absen < 1 || absen > len(dataTeman) {
		fmt.Println("Nomor absen tidak valid")
		return
	}

	// absen-1 digunakan karena indeks dalam slice dimulai dari 0, sedangkan nomor absen biasanya dimulai dari 1.
	// absen-1 akan menghasilkan indeks yang sesuai dengan nomor absen yang diberikan.
	teman := dataTeman[absen-1]
	displayDataTeman(teman)
}

// mengonversi string nomorAbsen menjadi integer. Jika konversi berhasil, maka fungsi akan mengembalikan integer tersebut beserta nil untuk nilai error
func parseAbsen(nomorAbsen string) (int, error) {
	var absen int
	_, err := fmt.Sscanf(nomorAbsen, "%d", &absen)
	return absen, err
}

// mencetak informasi teman ke konsol
func displayDataTeman(teman Teman) {
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}
