package main

import (
	"fmt"
	"strconv"
)

func main() {
	// meminta pengguna untuk memasukkan angka, kemudian memeriksa apakah input valid atau tidak
	fmt.Print("Masukkan angka: ")
	input, err := getInput()
	if err != nil {
		fmt.Println("Input tidak valid. Masukkan angka positif.")
		return
	}

	// Cek apakah angka tersebut genap atau ganjil
	if input%2 == 0 {
		fmt.Println("Bilangan ini genap")
	} else {
		fmt.Println("Bilangan ini ganjil")
	}

	// Cetak semua angka genap dari 1 hingga angka yang diinput
	fmt.Println("Bilangan genap dari 1 hingga masukan Anda adalah: ")
	for i := 1; i <= input; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Cek apakah angka tersebut adalah bilangan prima
	if isPrime(input) {
		fmt.Println("Angka adalah bilangan prima.")
	} else {
		fmt.Println("Angka bukan bilangan prima.")
	}
}

// untuk mendapatkan input dari pengguna dalam bentuk integer, dan memberikan nilai integer tersebut bersama dengan nilai error (jika ada) sebagai hasil dari fungsi.
func getInput() (int, error) {
	var inputStr string
	fmt.Scanln(&inputStr)
	input, err := strconv.Atoi(inputStr)
	return input, err
}

// Fungsi untuk memeriksa apakah sebuah angka adalah bilangan Prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
