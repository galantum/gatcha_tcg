package main

import (
	"fmt"
	"math"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	var total int
	var perN, perR, perSR, perUR, perSUR, perSSR float64
	var n, r, sr, ur, sur, ssr, open int

	// Mengatur seed random agar hasilnya berbeda setiap kali dijalankan
	rand.Seed(uint64(time.Now().UnixNano()))

	// Meminta total item
	fmt.Print("Type a Total Item (min 100): ")
	fmt.Scanln(&total)
	if total < 100 {
		fmt.Println("Value must be at least 100")
		return
	}

	// Memasukkan presentase masing-masing kategori
	fmt.Print("Your presentase item Normal is (e.g., 70.8 for 70.8%): ")
	fmt.Scanln(&perN)
	n = int(math.Round(float64(total) * perN / 100))

	fmt.Print("Your presentase item Rare is (e.g., 10.9 for 10.9%): ")
	fmt.Scanln(&perR)
	r = int(math.Round(float64(total) * perR / 100))

	fmt.Print("Your presentase item Super Rare is (e.g., 8.3 for 8.3%): ")
	fmt.Scanln(&perSR)
	sr = int(math.Round(float64(total) * perSR / 100))

	fmt.Print("Your presentase item Ultra Rare is (e.g., 5.1 for 5.1%): ")
	fmt.Scanln(&perUR)
	ur = int(math.Round(float64(total) * perUR / 100))

	fmt.Print("Your presentase item Super Ultra Rare is (e.g., 3.5 for 3.5%): ")
	fmt.Scanln(&perSUR)
	sur = int(math.Round(float64(total) * perSUR / 100))

	fmt.Print("Your presentase item Super Special Rare is (e.g., 1.4 for 1.4%): ")
	fmt.Scanln(&perSSR)
	ssr = int(math.Round(float64(total) * perSSR / 100))

	// Sisanya untuk kategori terakhir
	//ssr2 = total - (n + r + sr + ur + sur)

	// Meminta input jumlah pack yang akan dibuka
	fmt.Print("Open pack: ")
	fmt.Scanln(&open)

	// Buat slice dengan kapasitas total
	list := make([]string, int(total))

	// Isi slice langsung tanpa append
	index := 0
	index = fillSlice(list, index, "n", n)
	index = fillSlice(list, index, "r", r)
	index = fillSlice(list, index, "sr", sr)
	index = fillSlice(list, index, "ur", ur)
	index = fillSlice(list, index, "sur", sur)
	fillSlice(list, index, "ssr", ssr)

	// fmt.Println("total perhitungan presentase ssr:", ssr1)
	// fmt.Println("total perhitungan pengurangan ssr:", ssr2)

	// Shuffle slice untuk variasi
	shuffleSlice(list)

	// Ambil nilai acak beberapa kali
	for i := 0; i < open; i++ {
		value, updatedResult := popRandomValue(list)
		fmt.Printf("Random Value: %s, Remaining Length: %d\n", value, open-i)
		list = updatedResult
	}
}

// Fungsi untuk mengisi slice dengan nilai tertentu
func fillSlice(slice []string, startIndex int, value string, count int) int {
	for i := 0; i < count; i++ {
		slice[startIndex+i] = value
	}
	return startIndex + count
}

// Fungsi untuk mengacak urutan slice
func shuffleSlice(slice []string) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// Fungsi untuk mengambil nilai acak dari slice dan menghapusnya
func popRandomValue(slice []string) (string, []string) {
	if len(slice) == 0 {
		return "", slice // Jika slice kosong
	}

	// Pilih indeks acak
	randomIndex := rand.Intn(len(slice))

	// Ambil nilai acak
	value := slice[randomIndex]

	// Hapus elemen dengan menggabungkan slice sebelum dan sesudah indeks
	updatedSlice := append(slice[:randomIndex], slice[randomIndex+1:]...)

	return value, updatedSlice
}
