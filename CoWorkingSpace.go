package main

import "fmt"

const NMAX int = 50

type coWorkSpace struct {
	nama, fasilitas                             [10]string
	lokasi, ulasan                              [NMAX]string
	harga_sewa                                  float64
	rating                                      float64
	lenNama, lenLokasi, lenFasilitas, lenUlasan int
}
type coWorkS [NMAX]coWorkSpace

func main() {
	var cws coWorkS
	var nData int
	fmt.Scan(&nData)
	bacaData(&cws, nData)
	cetakData(cws, nData)
}

func bacaData(cws *coWorkS, n int) {
	var i, j int
	for i = 0; i < n; i++ {
		cws[i].lenNama = 1
		cws[i].lenLokasi = 1
		cws[i].lenFasilitas = 1
		cws[i].lenUlasan = 1
		j = 0
		fmt.Printf("Co-Working Space %d\n", i+1)
		fmt.Print("Nama Co-Working Space: ")
		fmt.Scan(&cws[i].nama[j])
		for cws[i].nama[j] != "#" {
			j++
			cws[i].lenNama++
			fmt.Scan(&cws[i].nama[j])
		}
		j = 0
		fmt.Print("Lokasi Co-Working Space: ")
		fmt.Scan(&cws[i].lokasi[j])
		for cws[i].lokasi[j] != "#" {
			j++
			cws[i].lenLokasi++
			fmt.Scan(&cws[i].lokasi[j])
		}
		j = 0
		fmt.Print("Fasilitas Co-Working Space: ")
		fmt.Scan(&cws[i].fasilitas[j])
		for cws[i].fasilitas[j] != "#" {
			j++
			cws[i].lenFasilitas++
			fmt.Scan(&cws[i].fasilitas[j])
		}
		j = 0
		fmt.Print("Harga Sewa Co-Working Space: ")
		fmt.Scan(&cws[i].harga_sewa)

		fmt.Print("Ulasan: ")
		fmt.Scan(&cws[i].ulasan[j])
		for cws[i].ulasan[j] != "#" {
			j++
			cws[i].lenUlasan++
			fmt.Scan(&cws[i].ulasan[j])
		}
		fmt.Print("Rating: ")
		fmt.Scan(&cws[i].rating)
	}
}

func cetakData(cws coWorkS, n int) {
	var i int
	var sum float64 = 0
	for i = 0; i < n; i++ {
		fmt.Printf("\nData Co-Working Space %d", i+1)
		fmt.Print("\nNama tempat: ")
		for j := 0; j < cws[i].lenNama-1; j++ {
			fmt.Print(cws[i].nama[j], " ")
		}
		fmt.Print("\nLokasi: ")
		for j := 0; j < cws[i].lenLokasi-1; j++ {
			fmt.Print(cws[i].lokasi[j], " ")
		}
		fmt.Print("\nFasilitas: ")
		for j := 0; j < cws[i].lenFasilitas-1; j++ {
			fmt.Print(cws[i].fasilitas[j], " ")
		}
		fmt.Printf("\nHarga sewa: Rp.%.0f", cws[i].harga_sewa)
		fmt.Print("\nUlasan: ")
		for j := 0; j < cws[i].lenUlasan-1; j++ {
			fmt.Print(cws[i].ulasan[j], " ")
		}
		fmt.Printf("\nRating: %.1f/5\n", cws[i].rating)
		sum = sum + cws[i].rating
	}
	//Kalkulasi rating
	fmt.Printf("\nRating Akhir: %.1f/5\n", sum/float64(n))
}
