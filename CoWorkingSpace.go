package main

import "fmt"

const NMAX int = 50

type coWorkSpace struct {
	nama                                        string
	fasilitas                                   [10]string
	lokasi                                      string
	ulasan                                      [NMAX]string
	harga_sewa                                  float64
	rating                                      float64
	lenNama, lenLokasi, lenFasilitas, lenUlasan int
}
type coWorkS [NMAX]coWorkSpace

type Ulasan struct {
	ID, coWorkingID int
	kategori        string
	username        string
	komentar        [15]string
	rating          float64
}
type arrUlasan [NMAX]Ulasan

func main() {
	var cws coWorkS
	var nData, inputUser int
	//munculkan kata selamat datang
	fmt.Println("Selamat Datang di Co Working Space")
	//munculkan opsi penggguna 1 pemilik, 2 pelanggan
	fmt.Println("Masukkan Pilihan Pengguna: ")
	fmt.Println("1. Pemilik\n2. Pelanggan ")
	fmt.Print("Masukkan Pilihan: ")
	//scan input pemilik atau pelangga
	fmt.Scan(&inputUser)
	switch inputUser {
	case 1:
		menuPemilik(&cws, &nData)
	case 2:
		menuPelanggan(&cws, &nData)
	default:
		fmt.Print("Pilihan tidak valid")
	}
}

func menuPemilik(cws *coWorkS, nData *int) {
	var pilihanPemilik int
	fmt.Println("Pilih Menu Pemilik: ")
	fmt.Println("1. Tambah data Co-working space")
	fmt.Println("2. Ubah data Co-Working space")
	fmt.Println("3. Hapus data Co-Working space")
	fmt.Println("0. Keluar")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihanPemilik)

	switch pilihanPemilik {
	case 1:
		tambahCWS(cws, nData)
	case 2:
		fmt.Print("data belum ada")
	case 3:
		fmt.Print("data belum ada")
	default:
		fmt.Print("Pilihan tidak valid")
	}
}

func tambahCWS(cws *coWorkS, nData *int) {
	var jmlData int
	fmt.Print("Masukkan jumlah data yang ditambahkan: ")
	fmt.Scan(&jmlData)
	idx := *nData
	for i := 0; i < jmlData; i++ {
		fmt.Printf("\nCoworking space ke-%d : \n", *nData+1)
		currentposition := &cws[*nData]
		currentposition.lenFasilitas = 0
		currentposition.lenUlasan = 0
		currentposition.rating = 0

		fmt.Print("Masukkan nama: ")
		fmt.Scan(&currentposition.nama)
		fmt.Print("Masukkan lokasi: ")
		fmt.Scan(&currentposition.lokasi)

		// for j:= 0; j < 10; j++{
		// 	var fasilitas string
		// 	fmt.Print("fasilitas ke %d: ")
		// 	fmt.Scan(&fasilitas)
		// 	if fasilitas == "#"{
		// 		break
		// 	}
		// }
		fmt.Println("Masukkan fasilitas maksimal 10, ketik # untuk berhenti")
		j := 0
		for j < 10 {
			var fasilitas string
			fmt.Printf("fasilitas ke %d: ", j+1)
			fmt.Scan(&fasilitas)
			if fasilitas == "#" {
				j = 10
			} else {
				currentposition.fasilitas[j] = fasilitas
				currentposition.lenFasilitas++
			}
			j++
		}
		fmt.Print("Masukkan harga sewa: ")
		fmt.Scan(&currentposition.harga_sewa)

		*nData++
	}
	fmt.Printf("Berhasil menambahkan %d Co-Working space \n", jmlData)

	//print semua Co-Working yang di tambahkan
	fmt.Println("Hasil semua data yang sudah ditambahkan")
	for i := idx; i < *nData; i++ {
		fmt.Println("Data Co-Working Space")
		fmt.Printf("\nCo-working space ke %d", i+1)
		fmt.Printf("\nNama: %s", cws[i].nama)
		fmt.Printf("\nLokasi: %s", cws[i].lokasi)
		fmt.Print("\nFasilitas: ")
		for j := 0; j < cws[i].lenFasilitas; j++ {
			fmt.Print(cws[i].fasilitas[j])
			if j < cws[i].lenFasilitas-1 {
				fmt.Print(", ")
			}
		}
		fmt.Printf("\nHarga sewa: %f", cws[i].harga_sewa)
		fmt.Printf("\nRating: %f", cws[i].rating)
	}

}

func fasilitas(cws *coWorkS, nData *int, keyWord string) {
	found := false
	for i := 0; i < *nData; i++ {
		for j := 0; j < 10; j++ {
			if keyWord == cws[i].fasilitas[j] {
				fmt.Println("Fasilitas: ", cws[i].fasilitas[j])
				found = true
			}
		}
	}
	if !found {
		fmt.Println("Fasilitas tidak ditemukan.")
	}
}

func menuPelanggan(cws *coWorkS, nData *int) {
	var pilihanPelanggan int
	fmt.Println("Pilih Menu Pelanggan: ")
	fmt.Println("1. Mencari Co-Working Space")
	fmt.Println("2. Mengurutkan Daftar Co-Working Space")
	fmt.Println("3. Memberikan Ulasan")
	fmt.Println("4. Memberikan Rating")
	fmt.Println("5. Mengubah Ulasan")
	fmt.Println("6. Mengubah Rating")
	fmt.Println("7. Menghapus Ulasan")
	fmt.Println("8. Menghapus Rating")
	fmt.Println("0. Exit")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihanPelanggan)

	// switch pilihanPelanggan {
	// case 1:
	// 	mencariCWS(cws, nData)
	// case 2:
	// 	fmt.Print("data belum ada")
	// case 3:
	// 	fmt.Print("data belum ada")
	// default:
	// 	fmt.Print("Pilihan tidak valid")
	// }
}

func hapusUlasan(arr *arrUlasan, n *int) {
	var user string
	fmt.Print("Masukkan username yang ingin dihapus ulasannya: ")
	fmt.Scan(&user)

	found := false
	for i := 0; i < *n; i++ {
		if arr[i].username == user {
			// Reset isi ulasan
			arr[i].username = ""
			for j := 0; j < arr[i].lenKomentar; j++ {
				arr[i].komentar[j] = ""
			}
			arr[i].rating = 0.0
			found = true
		}
	}
	if found {
		fmt.Println("Semua ulasan oleh", user, "berhasil dihapus.")
	} else {
		fmt.Println("Ulasan oleh", user, "tidak ditemukan.")
	}
}

func editUlasan(arr *arrUlasan, n int) {
	var user string
	fmt.Print("Masukkan username yang ingin diedit ulasannya: ")
	fmt.Scan(&user)

	edit := false
	for i := 0; i < n; i++ {
		if arr[i].username == user {
			edit = true

			// Tampilkan ulasan sebelumnya
			fmt.Println("\nUlasan Sebelumnya: ")
			fmt.Println("Username:", arr[i].username)
			fmt.Print("Komentar: ")
			for j := 0; j < 15; j++ {
				if arr[i].komentar[j] != "" {
					if j > 0 && arr[i].komentar[j-1] != "" {
						fmt.Print(", ")
					}
					fmt.Print(arr[i].komentar[j])
				}
			}
			fmt.Println()
			fmt.Println("Rating:", arr[i].rating)

			// Input komentar baru
			fmt.Println("\nMasukkan komentar baru (maks 15):")
			var komentarBaru [15]string
			var stopInput bool
			for j := 0; j < 15; j++ {
				if !stopInput {
					var komen string
					fmt.Printf("Komentar %d (kosongkan untuk selesai): ", j+1)
					fmt.Scanln(&komen)
					if komen == "" {
						stopInput = true
					} else {
						komentarBaru[j] = komen
					}
				}
			}
			// Simpan komentar baru
			for j := 0; j < 15; j++ {
				arr[i].komentar[j] = komentarBaru[j]
			}

			// Input rating baru
			fmt.Print("Masukkan rating baru: ")
			fmt.Scan(&arr[i].rating)

			// Tampilkan ulasan terbaru
			fmt.Println("\nUlasan Terbaru: ")
			fmt.Println("Username: ", arr[i].username)
			fmt.Print("Komentar: ")
			for j := 0; j < 15; j++ {
				if arr[i].komentar[j] != "" {
					if j > 0 && arr[i].komentar[j-1] != "" {
						fmt.Print(", ")
					}
					fmt.Print(arr[i].komentar[j])
				}
			}
			fmt.Println()
			fmt.Println("Rating: ", arr[i].rating)
			fmt.Println("Ulasan berhasil diperbarui\n")
		}
	}
	if !edit {
		fmt.Println("Ulasan oleh", user, "tidak ditemukan")
	}
}

// func bacaData(cws *coWorkS, n int) {
// 	var i, j int
// 	for i = 0; i < n; i++ {
// 		cws[i].lenNama = 1
// 		cws[i].lenLokasi = 1
// 		cws[i].lenFasilitas = 1
// 		cws[i].lenUlasan = 1
// 		j = 0
// 		fmt.Printf("Co-Working Space %d\n", i+1)
// 		fmt.Print("Nama Co-Working Space: ")
// 		fmt.Scan(&cws[i].nama[j])
// 		for cws[i].nama[j] != "#" {
// 			j++
// 			cws[i].lenNama++
// 			fmt.Scan(&cws[i].nama[j])
// 		}
// 		j = 0
// 		fmt.Print("Lokasi Co-Working Space: ")
// 		fmt.Scan(&cws[i].lokasi[j])
// 		for cws[i].lokasi[j] != "#" {
// 			j++
// 			cws[i].lenLokasi++
// 			fmt.Scan(&cws[i].lokasi[j])
// 		}
// 		j = 0
// 		fmt.Print("Fasilitas Co-Working Space: ")
// 		fmt.Scan(&cws[i].fasilitas[j])
// 		for cws[i].fasilitas[j] != "#" {
// 			j++
// 			cws[i].lenFasilitas++
// 			fmt.Scan(&cws[i].fasilitas[j])
// 		}
// 		j = 0
// 		fmt.Print("Harga Sewa Co-Working Space: ")
// 		fmt.Scan(&cws[i].harga_sewa)

// 		fmt.Print("Ulasan: ")
// 		fmt.Scan(&cws[i].ulasan[j])
// 		for cws[i].ulasan[j] != "#" {
// 			j++
// 			cws[i].lenUlasan++
// 			fmt.Scan(&cws[i].ulasan[j])
// 		}
// 		fmt.Print("Rating: ")
// 		fmt.Scan(&cws[i].rating)
// 	}
// }

// func cetakData(cws coWorkS, n int) {
// 	var i int
// 	var sum float64 = 0
// 	for i = 0; i < n; i++ {
// 		fmt.Printf("\nData Co-Working Space %d", i+1)
// 		fmt.Print("\nNama tempat: ")
// 		for j := 0; j < cws[i].lenNama-1; j++ {
// 			fmt.Print(cws[i].nama[j], " ")
// 		}
// 		fmt.Print("\nLokasi: ")
// 		for j := 0; j < cws[i].lenLokasi-1; j++ {
// 			fmt.Print(cws[i].lokasi[j], " ")
// 		}
// 		fmt.Print("\nFasilitas: ")
// 		for j := 0; j < cws[i].lenFasilitas-1; j++ {
// 			fmt.Print(cws[i].fasilitas[j], " ")
// 		}
// 		fmt.Printf("\nHarga sewa: Rp.%.0f", cws[i].harga_sewa)
// 		fmt.Print("\nUlasan: ")
// 		for j := 0; j < cws[i].lenUlasan-1; j++ {
// 			fmt.Print(cws[i].ulasan[j], " ")
// 		}
// 		fmt.Printf("\nRating: %.1f/5\n", cws[i].rating)
// 		sum = sum + cws[i].rating
// 	}
// 	//Kalkulasi rating
// 	fmt.Printf("\nRating Akhir: %.1f/5\n", sum/float64(n))
// }
