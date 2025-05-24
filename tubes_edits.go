package main

import "fmt"

const NMAX int = 20

/*deklarasi tipe bentukan coWorkSpace yang akan menyimpan data-data co-working space
seperti nama(string), lokasi(string), array fasilitas, harga sewa(real),
rating total(real), serta panjang elemen array fasilitas*/
type coWorkSpace struct {
	ID                 int
	nama, lokasi       string
	fasilitas          [15]string
	harga_sewa, rating float64
	lenFasilitas       int
}
type coWorkS [NMAX]coWorkSpace

/*deklarasi tipe bentukan Ulasan utk menyimpan data-data ulasan
seperti ID, coWorkingID, username, komentar, rating dari pengguna*/
type Ulasan struct {
	ID, coWorkingID    int
	username string
	komentar [15]string
	rating             float64
}
type arrUlasan [NMAX]Ulasan

func main() {
	// deklarasi variabel yang dibutuhkan
	var cws coWorkS
	var ulas arrUlasan
	var nData, n int
	var pilih, pilPemilik, pilPekerja, pilTambah, pilCari int
	var aksesKe, aksesPilTambah int
	var indexCws int
	var keyNama, keyLokasi string
	var valid, back bool
	valid = false
	back = false
	aksesKe = 1
	aksesPilTambah = 1
	for !valid { // terlalu gpt
		mainMenu()
		fmt.Scan(&pilih)
		// pengecekan pilihan pengguna
		if pilih == 3 {
			valid = true
		} else if pilih == 1 {
			menuTambahTempat()
			fmt.Scan(&nData)
			bacaData(&cws, nData)
			cetakData()
			menuPemilik()
			fmt.Scan(&pilPemilik)
			if pilPemilik == 1 {
				ubahData(&cws)
			} else {
				hapusData(&cws, &nData)
			}
		} else { //pilihan 2
			if aksesKe == 1 {
				DataKosong()
			} else {
				for !back {
					menuPekerja()
					fmt.Scan(&pilPekerja)
					if pilPekerja == 7 {
						back = true
					} else if pilPekerja == 1 {
						menuTambahUlasan()
						bacaUlasan(&ulas)
					}
				}
			}
		}
		} if pilih == 2 {
			/*menu baru (menuTambah) akan dimunculkan kemudian pengguna bisa memilih
			akan melakukan aksi apa*/
			menuTambah()
			fmt.Scan(&pilTambah)
			switch pilTambah {
			case 1: // pengguna akan menambahkan data co-working space
				if aksesKe == 1 && aksesPilTambah == 1 {
					menuTambahTempat()
					fmt.Scan(&nData)
					bacaData(&cws, nData)
					aksesPilTambah++
				} else {
					/* jika pengguna ingin menambahkan data lagi padahal data telah tersedia,
					program akan menampilkan bahwa data telah tersedia dan pengguna tidak bisa
					menambahkan data lagi*/
					DataSudahAda()
				}
			case 2: // pengguna akan menambahkan ulasan & rating terhadap co-working space
				cetakData(cws, nData)
				menuUlasan()
				fmt.Scan(&n)
				bacaUlasan(&cws, n-1) //membaca ulasan untuk data co-working space index ke n-1
			case 3:
				//ubahData()

			case 4:
				//ubahUlasan()

			}
		} else if pilih == 3 {
			/* jika pengguna memilih 3 (mencari co-working space), akan dicek terlebih dulu apakah pengguna
			mengakses pertama kali atau tidak*/
			if aksesKe == 1 {
				/*jika pengguna ingin mencari co-working space pada akses pertama, maka
				pengguna akan diberitahu bahwa data co-working space belum ada*/
				DataKosong()
			} else {
				/*kondisi else adalah ketika pengguna ingin mencari co-working space bukan pada akses
				pertama sehingga pengguna akan diarahkan ke menuCariData dan diminta untuk memilih
				ingin mencari co-working space berdasarkan nama atau lokasi nya*/
				menuCariData()
				fmt.Scan(&pilCari)
				if pilCari == 1 {
					// cari co-working space berdasarkan nama (Sequential Search)
					fmt.Print("Masukkan nama co-working space yang anda cari: ")
					fmt.Scan(&keyNama)
					indexCws = cariNama(cws, nData, keyNama)
					if indexCws == -1 {
						// co-working space tidak ada
						DataKosong()
					} else {
						// co-working space ada
						fmt.Println("Data ditemukan!")
						printData(cws, indexCws)
					}
				} else {
					// cari co-working space berdasarkan lokasi (sequential search)
					fmt.Print("Masukkan lokasi co-working space yang anda cari: ")
					fmt.Scan(&keyLokasi)
					indexCws = cariLokasi(cws, nData, keyLokasi)
					if indexCws == -1 {
						// co-working space tidak ada
						DataKosong()
					} else {
						// co-working space ada
						fmt.Println("Data ditemukan!")
						printData(cws, indexCws)
					}
				}
			}
		} else {
			// kondisi else yaitu pengguna ingin mengurutkan data
		}
		aksesKe++
	}
}

func bacaUlasan(ulas *arrUlasan) {
	var IdCws int
	var i, j int
	var sum float64 = 0
	var idUlas int = 1
	fmt.Scan(&IdCws)
	i = IdCws-1

	fmt.Println("+======================================+")
	fmt.Printf("| %-36s |\n", "Komentar: isi dengan kategori dibawah ini:")
	fmt.Printf("| %-36s |\n", "1. Nyaman")
	fmt.Printf("| %-36s |\n", "2. Bersih")
	fmt.Printf("| %-36s |\n", "3. Tenang")
	fmt.Printf("| %-36s |\n", "4. Wifi/Internet_kencang")
	fmt.Printf("| %-36s |\n", "5. Harga_terjangkau")
	fmt.Printf("| %-36s |\n", "6. Fasilitas_lengkap")
	fmt.Printf("| %-36s |\n", "7. Dekat_kampus")
	fmt.Printf("| %-36s |\n", "8. Pelayanan_cepat")
	fmt.Printf("| %-36s |\n", "9. View_bagus")
	fmt.Printf("| %-36s |\n", "10. Dekat_kuliner")
	fmt.Printf("| %-36s |\n", "11. Akses_mudah")
	fmt.Printf("| %-36s |\n", "12. Tempat_rapi")
	fmt.Printf("| %-36s |\n", "13. AC_dingin")
	fmt.Printf("| %-36s |\n", "14. Toilet_bersih")
	fmt.Printf("| %-36s |\n", "15. Satpam_ramah")
	fmt.Printf("| %-36s |\n", "Maksimal 15 komentar, atau")
	fmt.Printf("| %-36s |\n", "Jika telah selesai memilih kategori")
	fmt.Printf("| %-36s |\n", "diakhiri dengan #")
	fmt.Printf("| %-36s |\n", "EX: Nyaman Satpam_ramah #")
	fmt.Println("+======================================+")

	fmt.Printf("Ulasan Co-Working Space dengan ID:%d\n", IdCws)
	fmt.Printf("\nUlasan ke-%d\n", i+1)
	ulas[i].ID = IdCws
	fmt.Printf("ID Ulasan: %d\n", ulas[i].ID)
	fmt.Print("Username: ")
	fmt.Scan(&ulas[i])
	fmt.Print("Komentar: ")
	fmt.Scan(&ulas[i].komentar[j])
	for ulas[i].komentar[j] != "#" {
		
	}
	cws[n].rating = sum / float64(j)
}

func hapusData(cws *coWorkS, n *int) {
	var dataKe, i int
	fmt.Println("+======================================================+")
	fmt.Printf("|%5s%s%5s|\n", "", "Menghapus data co-working space ke berapa?", "")
	fmt.Println("+======================================================+")
	fmt.Print("Co-working space: ")
	fmt.Scan(&dataKe)
	for i = dataKe; i < *n; i++ {
		cws[i] = cws[i+1]
	}
	*n--
}

func ubahData(cws *coWorkS) {
	var dataKe, i, j int
	fmt.Println("+======================================================+")
	fmt.Printf("|%5s%s%5s|\n", "", "Mengubah data co-working space ke berapa?", "")
	fmt.Println("+======================================================+")
	fmt.Print("Co-working space: ")
	fmt.Scan(&dataKe)
	i = dataKe - 1
	fmt.Printf("\nCo-Working Space %d\n", i+1)
	fmt.Printf("Nama Co-Working Space sebelumnya: %s\n", cws[i].nama)
	fmt.Print("Masukkan Nama Co-Working Space yang baru: ")
	fmt.Scan(&cws[i].nama)

	fmt.Printf("Lokasi Co-Working Space sebelumnya: %s\n", cws[i].lokasi)
	fmt.Print("Masukkan Nama Co-Working Space yang baru: ")
	fmt.Scan(&cws[i].lokasi)

	fmt.Print("Fasilitas Co-Working Space sebelumnya: ")
	for j := 0; j < cws[i].lenFasilitas; j++ {
		fmt.Print(cws[i].fasilitas[j])
		if j < cws[i].lenFasilitas-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("Masukkan Fasilitas Co-Working Space yang baru: ")
	fmt.Println("Masukkan fasilitas maksimal 10 / ketik # untuk berhenti")
	cws[i].lenFasilitas = 1
	j = 0
	for j < 10 {
		fmt.Printf("Fasilitas ke %d: ", j+1)
		fmt.Scan(&cws[i].fasilitas[j])
		if cws[i].fasilitas[j] == "#" {
			j = 10
			cws[i].lenFasilitas++
			j++
		}
	}

	fmt.Print("Harga Sewa Co-Working Space sebelumnya: Rp. %.0f\n", cws[i].harga_sewa)
	fmt.Print("Masukkan harga sewa Co-Working Space yang baru: ")
	fmt.Scan(&cws[i].harga_sewa)
}

func cariFasilitas(cws *coWorkS, nData *int, keyWord string) {
	var found bool = false
	var i, j int
	for i = 0; i < *nData; i++ {
		for j = 0; j < 10; j++ {
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

func printData(cws coWorkS, idx int) {
	var j int
	fmt.Printf("\nData Co-Working Space %d", idx+1)
	fmt.Printf("\nNama tempat: %s", cws[idx].nama)

	fmt.Printf("\nLokasi: %s", cws[idx].lokasi)
	fmt.Print("\nFasilitas: ")
	for j = 0; j < cws[idx].lenFasilitas-1; j++ {
		fmt.Print(cws[idx].fasilitas[j], " ")
	}
	fmt.Printf("\nHarga sewa: Rp.%.0f", cws[idx].harga_sewa)

	fmt.Print("\nUlasan: ")
	for j = 0; j < cws[idx].lenUlasan-1; j++ {
		fmt.Print(cws[idx].ulasan[j], " ")
	}
	fmt.Printf("\nRating: %.1f/5\n", cws[idx].rating)
	fmt.Println()
}

func cariNama(cws coWorkS, n int, key string) int { // binary search A-Z mencari nama
	var left, mid, right, idx int
	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if cws[mid].nama == key {
			idx = mid
		} else if cws[mid].nama < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

func cariLokasi(cws coWorkS, n int, key string) int {
	var idx, i int
	idx = -1
	i = 0
	for i < n && idx == -1 {
		if cws[i].lokasi == key {
			idx = i
		}
		i++
	}
	return idx
}

func DataSudahAda() {
	fmt.Printf("+===========================+\n")
	fmt.Printf("|%2s%s%2s|\n", "", "Tidak bisa menambah data!", "")
	fmt.Printf("|%4s%s%4s|\n", "", "Data sudah tersedia", "")
	fmt.Printf("+===========================+\n")
}

func DataKosong() {
	fmt.Printf("+===========================+\n")
	fmt.Printf("|%4s%s%4s|\n", "", "Data tidak tersedia", "")
	fmt.Printf("+===========================+\n")
}

func mainMenu() {
	fmt.Println("+========== Coseum [Platform Co-Working Space] ==========+")
	fmt.Printf("|  %s  |\n", "Selamat datang di Coseum [Penyedia Co-Working Space]")
	fmt.Printf("|%21s%s%21s|\n", "", "Masuk Sebagai:", "")
	fmt.Printf("| %-54s |\n", "1. Pemilik Co-Working Space")
	fmt.Printf("| %-54s |\n", "2. Pengguna Co-Working Space (Pekerja)")
	fmt.Printf("| %-54s |\n", "3. exit")
	fmt.Println("+========================================================+")
	fmt.Print("Pilih (1/2/3)? ")
}

func menuFasilitas() {
	fmt.Println("+======================================================+")
	fmt.Printf("|%5s%s%5s|\n", "", "Fasilitas apa yang anda cari?", "")
	fmt.Println("+======================================================+")
	fmt.Print("Fasilitas: ")
}

func menuUlasan() {
	fmt.Println("+========== Coseum [Platform Co-Working Space] ==========+")
	fmt.Printf("|%13s%s%13s|\n", "", "Pilih yang ingin anda lakukan:", "")
	fmt.Printf("| %-54s |\n", "1. Ubah Ulasan Co-Working Space")
	fmt.Printf("| %-54s |\n", "2. Hapus Ulasan Co-Working Space")
	fmt.Printf("| %-54s |\n", "3. exit")
	fmt.Println("+========================================================+")
	fmt.Print("Pilih (1/2/3)? ")
}

func menuPekerja() {
	fmt.Println("+========== Coseum [Platform Co-Working Space] ==========+")
	fmt.Printf("|  %s  |\n", "Selamat datang di Coseum [Penyedia Co-Working Space]")
	fmt.Printf("|%13s%s%13s|\n", "", "Pilih yang ingin anda lakukan:", "")
	fmt.Printf("| %-54s |\n", "1. Tambah Ulasan Co-Working Space")
	fmt.Printf("| %-54s |\n", "2. Cari berdasarkan Nama(Sequential)")
	fmt.Printf("| %-54s |\n", "3. Cari berdasarkan Lokasi(Binary)") // data diurutin trs ditampilin trs diminta utk input keyLokasi
	fmt.Printf("| %-54s |\n", "4. Urutkan berdasarkan Harga sewa Terrendah")
	fmt.Printf("| %-54s |\n", "5. Urutkan berdasarkan Rating Tertinggi")
	fmt.Printf("| %-54s |\n", "6. Tampilkan data berdasar fasilitas")
	fmt.Printf("| %-54s |\n", "7. exit")
	fmt.Println("+========================================================+")
	fmt.Print("Pilih (1/2/3/4/5/6/7)? ")
}

func menuPemilik() {
	fmt.Println("+========== Coseum [Platform Co-Working Space] ==========+")
	fmt.Printf("|%13s%s%13s|\n", "", "Pilih yang ingin anda lakukan:", "")
	fmt.Printf("| %-54s |\n", "1. Ubah Data Co-Working Space")
	fmt.Printf("| %-54s |\n", "2. Hapus Data Co-Working Space")
	fmt.Printf("| %-54s |\n", "3. exit")
	fmt.Println("+========================================================+")
	fmt.Print("Pilih (1/2/3)? ")
}

func menuTambahTempat() {
	fmt.Println("+======================================================+")
	fmt.Printf("|%5s%s%5s|\n", "", "Silakan masukkan data co-working space anda!", "")
	fmt.Println("+======================================================+")
	fmt.Print("Banyak co-working space yang akan dimasukkan: ")
}

func menuTambahUlasan() {
	fmt.Println("+========================================================+")
	fmt.Printf("|%7s%s%7s|\n", "", "Tambah Ulasan untuk Co-Working Space berapa?", "")
	fmt.Println("+========================================================+")
	fmt.Print("Ulasan Co-Working Space Nomor: ")
}

func bacaData(cws *coWorkS, n int) {
	var i, j int
	var idCws int = 1
	for i = 0; i < n; i++ {
		cws[i].lenFasilitas = 1
		cws[i].ID = idCws
		fmt.Printf("\nCo-Working Space %d\n", i+1)
		fmt.Printf("ID Co-Working Space: %d\n", cws[i].ID)
		fmt.Print("Nama Co-Working Space [Gunakan _ sebagai pengganti spasi]: ")
		fmt.Scan(&cws[i].nama)

		fmt.Print("Lokasi Co-Working Space[Gunakan _ sebagai pengganti spasi]: ")
		fmt.Scan(&cws[i].lokasi)

		fmt.Println("Masukkan fasilitas maksimal 10 / ketik # untuk berhenti")
		j = 0
		for j < 10 {
			var fasilitas string
			fmt.Printf("Fasilitas ke %d: ", j+1)
			fmt.Scan(&fasilitas)
			if fasilitas == "#" {
				j = 10
			} else {
				cws[i].fasilitas[j] = fasilitas
				cws[i].lenFasilitas++
			}
			j++
		}
		fmt.Print("Harga Sewa Co-Working Space: ")
		fmt.Scan(&cws[i].harga_sewa)
		idCws++
	}
}

func cetakData(cws coWorkS, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("\nData Co-Working Space %d", i+1)
		fmt.Printf("\nNama tempat: %s", cws[i].nama)

		fmt.Printf("\nLokasi: %s", cws[i].lokasi)
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
		fmt.Println()
	}
}
