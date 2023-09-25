package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// memastikan bahwa program dijalankan dengan jumlah argumen yang benar. jika tidak, memberikan petunjuk kepada pengguna tentang cara menggunakan program ini.
	args := os.Args // os.Args adalah slice dari string yang berisi semua argumen yang diberikan pada command line.

	if len(args) != 4 {
		fmt.Println("Usage: calculator <operation> <operand1> <operand2>")
		return
	}

	// mengambil operasi dari argumen kedua (indeks 1) dalam slice args
	operation := args[1]
	// untuk mengambil dan memastikan bahwa operand pertama adalah angka yang valid serta mengonversi ke float64
	operand1, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Println("Operand 1 harus berupa angka")
		return
	}

	// untuk mengambil dan memastikan bahwa operand kedua adalah angka yang valid serta mengonversi ke float64
	operand2, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		fmt.Println("Operand 2 harus berupa angka")
		return
	}

	// melakukan operasi sesuai dengan perintah
	result := 0.0 // untuk menyimpan hasil dari operasi aritmatika yang akan dilakukan
	switch operation {
	case "addition":
		result = operand1 + operand2
	case "subtraction":
		result = operand1 - operand2
	case "multiplication":
		result = operand1 * operand2
	case "division":
		if operand2 != 0 {
			result = operand1 / operand2
		} else {
			fmt.Println("Tidak dapat melakukan pembagian dengan nol")
			return
		}
	default:
		fmt.Println("Operasi tidak valid. Pilih dari: addition, subtraction, multiplication, division")
		return
	}

	// menampilkan hasil
	fmt.Printf("Hasil %s dari %.2f dan %.2f adalah %.2f\n", operation, operand1, operand2, result)
}
