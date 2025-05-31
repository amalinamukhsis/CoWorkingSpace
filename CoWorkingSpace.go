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
	namaCW string
	nama   string
	isi    [15]string
	rating float64
	lenIsi int
}
type arrUlasan [NMAX]Ulasan

func main() {
	var cws coWorkS
	var nData, inputUser int
	//munculkan kata selamat datang
	jalan := true
	for jalan {
		fmt.Println("Selamat Datang di Co Working Space")

		//munculkan opsi penggguna 1 pemilik, 2 pelanggan
		fmt.Println("Masukkan Pilihan Pengguna: ")
		fmt.Println("1. Pemilik\n2. Pelanggan\n0. Keluar ")
		fmt.Print("Masukkan Pilihan: ")
		//scan input pemilik atau pelanggan
		fmt.Scan(&inputUser)
		switch inputUser {
		case 1:
			menuPemilik(&cws, &nData)

		case 2:
			menuPelanggan(&cws, nData)
		case 0:
			fmt.Println("Terima Kasih")
			jalan = false //program di berhentikan.
		default:
			fmt.Print("Pilihan tidak valid")
		}
	}
}

func menuPemilik(cws *coWorkS, nData *int) {
	var pilihanPemilik int
	kembali := true
	for kembali {
		fmt.Println("\nPilih Menu Pemilik: ")
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
			ubahDataCW(cws)
			tampilCws(cws, *nData)
		case 3:
			hapusDataCW(cws, nData)
			tampilCws(cws, *nData)

		case 0:
			kembali = false
		default:
			fmt.Print("Pilihan tidak valid")
		}
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

		fmt.Print("Masukkan nama (nama Co-Working space harus unik): ")
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
		fmt.Println("\nData Co-Working Space")
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
		fmt.Printf("\nHarga sewa: %.2f", cws[i].harga_sewa)
		fmt.Printf("\nRating: %.2f", cws[i].rating)
	}

}

func ubahDataCW(cws *coWorkS) {
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
	fmt.Println("\nMasukkan Fasilitas Co-Working Space yang baru: ")
	fmt.Println("Masukkan fasilitas maksimal 10 / ketik # untuk berhenti")
	cws[i].lenFasilitas = 1
	j = 0
	for j < 10 {
		fmt.Printf("Fasilitas ke %d: ", j+1)
		fmt.Scan(&cws[i].fasilitas[j])
		if cws[i].fasilitas[j] == "#" {
			j = 10
		} else {
			cws[i].lenFasilitas++
			j++
		}
	}

	fmt.Printf("Harga Sewa (perbulan) Co-Working Space sebelumnya: Rp.%.0f\n", cws[i].harga_sewa)
	fmt.Print("Masukkan harga sewa (perbulan) Co-Working Space yang baru: Rp.")
	fmt.Scan(&cws[i].harga_sewa)
}
func hapusDataCW(cws *coWorkS, nData *int) {
	var dataKe, i int
	fmt.Println("+======================================================+")
	fmt.Printf("|%5s%s%5s|\n", "", "Menghapus data co-working space ke berapa?", "")
	fmt.Println("+======================================================+")
	fmt.Print("Co-working space: ")
	fmt.Scan(&dataKe)
	for i = dataKe - 1; i < *nData-1; i++ {
		(*cws)[i] = (*cws)[i+1]
	}
	*nData--
}

func menuPelanggan(cws *coWorkS, nData int) {
	var pilihanPelanggan int
	jalan := true
	for jalan {
		fmt.Println("Pilih Menu Pelanggan: ")
		fmt.Println("1. Tampilkan Co-Working Space")
		fmt.Println("2. Rating dan Ulasan")
		fmt.Println("3. Mencari Co-Working Space")
		fmt.Println("4. Mengurutkan Co-Working Space")
		fmt.Println("5. Filter Co-Working Space")
		fmt.Println("0. Exit")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihanPelanggan)

		switch pilihanPelanggan {
		case 1:
			tampilCws(cws, nData)
		case 2:
			menuRatingdanUlasan(cws, nData)
		case 3:
			menuCari(cws, nData)
		case 4:
			menuSort(cws, nData)
		case 5:
			menuFilter(cws, nData)
		case 0:
			jalan = false
		default:
			fmt.Print("Pilihan tidak valid")
		}
	}
}

func menuCari(cws *coWorkS, nData int) {
	var untukCari int
	program := true
	for program == true {
		fmt.Println("Cari Co-Working Space berdasarkan:")
		fmt.Println("1. Lokasi\n2. Nama\n0. exit")
		fmt.Print("Pilihan yang mau dicari: ")
		fmt.Scan(&untukCari)
		switch untukCari {
		case 1:
			cariLokasicws(*cws, nData)
		case 2:
			cariNamaCws(*cws, nData)
		case 0:
			program = false
		default:
			fmt.Println("Tidak valid")
		}
	}
}

func menuSort(cws *coWorkS, nData int) {
	var urut int
	program := true
	for program == true {
		fmt.Println("Urutkan Co-Working Space berdasarkan:")
		fmt.Println("1. Harga Sewa\n2. Rating \n0. exit")
		fmt.Print("Pilihan yang mau di urutkan: ")
		fmt.Scan(&urut)
		switch urut {
		case 1:
			selectionSortHargaSewa(cws, nData)
		case 2:
			insertionSortRating(cws, nData)
		case 0:
			program = false
		default:
			fmt.Println("Tidak valid")
		}
	}
}

func menuFilter(cws *coWorkS, nData int) {
	var filter string
	fmt.Print("Masukkan Fasilitas yang akan difilter: ")
	fmt.Scan(&filter)
	found := false
	for i := 0; i < nData; i++ {
		for j := 0; j < 10; j++ {
			if cws[i].fasilitas[j] == filter {
				fmt.Println("Nama Co-Working Space: ", cws[i].nama)
				fmt.Println("Lokasi Co-Working Space: ", cws[i].lokasi)
				fmt.Println("Harga Sewa Co-Working Space: ", cws[i].harga_sewa)
				fmt.Println("Rating: ", cws[i].rating)
				found = true
			}
		}
	}
	if found == false {
		fmt.Print("Tidak ditemukan")
	}
}

func menuRatingdanUlasan(cws *coWorkS, nData int) {
	var menuRating int
	program := true
	for program == true {
		fmt.Println("\nMenu rating dan ulasan:")
		fmt.Println("1. Tambah ulasan dan rating\n2. Edit ulasan dan rating\n3. hapus ulasan dan rating\n4. Tampilkan ulasan satu Co-Working Space\n0. Exit")
		fmt.Print("Masukkan Pilihan menu rating: ")
		fmt.Scan(&menuRating)
		switch menuRating {
		case 1:
			tambahUlasan(cws, nData)
		case 2:
			editUlasan(&daftarUlasan, jumlahUlasan)
		case 3:
			hapusUlasan(&daftarUlasan, &jumlahUlasan)
		case 4:
			tampilkanUlasan(&daftarUlasan, jumlahUlasan)
		case 0:
			program = false
		default:
			fmt.Println("Tidak valid")
		}
	}
}

func tampilCws(cws *coWorkS, nData int) {
	for i := 0; i < nData; i++ {
		fmt.Println("Nama Co-Working Space: ", cws[i].nama)
		fmt.Println("Lokasi Co-Working Space: ", cws[i].lokasi)
		fmt.Print("Fasilitas Co-working Space: ")
		for j := 0; j < cws[i].lenFasilitas; j++ {
			fmt.Print(cws[i].fasilitas[j])
			if j < cws[i].lenFasilitas-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("\nHarga Sewa Co-Working Space: ", cws[i].harga_sewa)
		fmt.Printf("Rating: %.2f ", cws[i].rating)
	}
}

/*cari lokasi berdasarkan sequential search*/
func cariLokasicws(cws coWorkS, nData int) {
	var lokasi string
	var found = false
	fmt.Print("Masukkan lokasi yang anda inginkan : ")
	fmt.Scan(&lokasi)
	fmt.Print("Co-Working space di ", lokasi)
	for i := 0; i < nData; i++ {
		if cws[i].lokasi == lokasi {
			fmt.Println("\nNama Co-Working Space: ", cws[i].nama)
			fmt.Println("Lokasi Co-Working Space: ", cws[i].lokasi)
			fmt.Println("Harga Sewa Co-Working Space: ", cws[i].harga_sewa)
			fmt.Println("Rating: ", cws[i].rating)
			found = true
			/*tidak bisa dibuat else karena akan selalu print out, perlu varibael penanda kalau sidah ditemukan tidak perlu masuk
			ke else tapi langsung berhenti*/
			// } else {
			// 	fmt.Println("Lokasi Co-Working space yang anda cari tidak di temukan")
		}
	}
	if found == false {
		fmt.Println("Lokasi Co-Working space yang anda cari tidak di temukan")

	}
}

/*prosedur untuk mengurutkan nama menggunakan selection sort*/
func selectionSortNama(cws coWorkS, nData int) coWorkS {
	for i := 0; i < nData-1; i++ {
		min := i
		for j := i + 1; j < nData; j++ {
			if cws[j].nama < cws[min].nama {
				min = j
			}
		}
		temp := cws[i].nama
		cws[i].nama = cws[min].nama
		cws[min].nama = temp
	}
	return cws
}

func binarySearchNama(cws coWorkS, nData int, nama string) bool {
	bawah := 0
	atas := nData

	for bawah <= atas {
		mid := (atas + bawah) / 2
		if cws[mid].nama == nama {
			return true
		} else if cws[mid].nama < nama {
			bawah = mid + 1
		} else {
			atas = mid - 1
		}
	}
	return false
}

func cariNamaCws(cws coWorkS, nData int) {
	var nama string
	fmt.Print("Masukkan nama Co-Working yang anda inginkan : ")
	fmt.Scan(&nama)
	fmt.Print("Nama Coworking ", nama)
	selectionSortNama(cws, nData)
	found := binarySearchNama(cws, nData, nama)
	if found == true {
		for i := 0; i < nData; i++ {
			if cws[i].nama == nama {
				fmt.Println("\nNama Co-Working Space: ", cws[i].nama)
				fmt.Println("Lokasi Co-Working Space: ", cws[i].lokasi)
				fmt.Println("Harga Sewa Co-Working Space: ", cws[i].harga_sewa)
				fmt.Println("Rating: ", cws[i].rating)
			}
		}
	} else {
		fmt.Println("Nama Co-Working space yang anda cari tidak di temukan")

	}

}

func selectionSortHargaSewa(cws *coWorkS, nData int) {
	for i := 0; i < nData-1; i++ {
		min := i
		for j := i + 1; j < nData; j++ {
			if cws[j].harga_sewa < cws[min].harga_sewa {
				min = j
			}
		}
		temp := cws[i].harga_sewa
		cws[i].harga_sewa = cws[min].harga_sewa
		cws[min].harga_sewa = temp
	}
	tampilCws(cws, nData)
}

func insertionSortRating(cws *coWorkS, nData int) {
	for i := 1; i < nData; i++ {
		max := i
		j := i - 1
		for j >= 0 && cws[j].rating < cws[max].rating {
			cws[j+1] = cws[j]
			j = j - 1
		}
		cws[j+1] = cws[i]
	}
	tampilCws(cws, nData)
}

func tambahUlasan(cws *coWorkS, nData int) {
	var u Ulasan
	var found bool

	fmt.Print("Masukkan nama co-working space yang ingin diulas: ")
	fmt.Scan(&u.namaCW)
	for i := 0; i < nData; i++ {
		if u.namaCW == cws[i].nama {
			found = true
		}
	}
	if found == false {
		fmt.Println("Co-working space tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan nama Anda: ")
	fmt.Scan(&u.nama)
	fmt.Println("Masukkan isi ulasan Anda perbaris, maksimal 15, ketik (.) untuk berhenti")
	jumlah := 0
	u.lenIsi = 0
	for jumlah < 15 {
		var ulas string
		fmt.Printf("masukkan ulasan anda: ")
		fmt.Scan(&ulas)
		if ulas == "." {
			jumlah = 15
		} else {
			u.isi[jumlah] = ulas
			u.lenIsi++
		}
		jumlah++
	}
	for i := 0; i < u.lenIsi; i++ {
		fmt.Printf("Ulasan ke-%d: %s\n", i+1, u.isi[i])
	}
	fmt.Print("Masukkan rating (0.0 - 5.0): ")
	fmt.Scan(&u.rating)
	daftarUlasan[jumlahUlasan] = u  // masukkan ulasan ke struct
	jumlahUlasan = jumlahUlasan + 1 //lalu tambhjumlah ulasan
	//setrlah rating masuk, hitung rata-rata rating untuk ditampilkan di sort by rating
	hitungRataRating(cws, nData)
	fmt.Println("Ulasan berhasil ditambahkan.")
}

var daftarUlasan [100]Ulasan
var jumlahUlasan int

func hitungRataRating(cws *coWorkS, nData int) {
	for i := 0; i < nData; i++ {
		var total float64
		var jml int
		for j := 0; j < jumlahUlasan; j++ {
			if daftarUlasan[j].namaCW == cws[i].nama {
				total += daftarUlasan[j].rating
				jml = jml + 1
			}
		}
		if jml > 0 {
			cws[i].rating = total / float64(jml)
		} else {
			cws[i].rating = 0
		}
	}
}

func editUlasan(u *[100]Ulasan, jumlahUlasan int) {
	var namaCW, namaUser string
	var found bool

	fmt.Print("Masukkan nama co-working space: ")
	fmt.Scan(&namaCW)
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scan(&namaUser)
	for i := 0; i < jumlahUlasan; i++ {
		if daftarUlasan[i].namaCW == namaCW && daftarUlasan[i].nama == namaUser {
			fmt.Print("Masukkan ulasan baru: ")
			fmt.Scan(&daftarUlasan[i].isi)
			fmt.Println("Ulasan berhasil diperbarui.")
			fmt.Println("Ulasan terbaru")
			fmt.Println(u[i].nama)
			fmt.Println(u[i].namaCW)
			fmt.Println(u[i].isi)
			fmt.Printf("%.2f", u[i].rating)
			found = true
		}
	}
	if found == false {
		fmt.Println("nama tidak ditemukan.")
	}
}
func hapusUlasan(u *[100]Ulasan, jumlahUlasan *int) {
	var namaCW, namaUser string
	var found bool

	fmt.Print("Masukkan nama co-working space: ")
	fmt.Scan(&namaCW)
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scan(&namaUser)
	for i := 0; i < *jumlahUlasan; i++ {
		if daftarUlasan[i].namaCW == namaCW && daftarUlasan[i].nama == namaUser && !found {
			// Geser elemen ke kiri
			for j := i; j < *jumlahUlasan-1; j++ {
				daftarUlasan[j] = daftarUlasan[j+1]
			}
			*jumlahUlasan = *jumlahUlasan - 1
			found = true
		}
	}

	if found == true {
		fmt.Println("Ulasan berhasil dihapus.")
	} else {
		fmt.Println("Ulasan tidak ditemukan.")
	}
}

func tampilkanUlasan(u *[100]Ulasan, jumlahUlasan int) {
	var namaSpace string
	var found bool
	fmt.Print("Masukkan nama co-working space: ")
	fmt.Scan(&namaSpace)
	fmt.Printf("\nDaftar ulasan: %s\n", namaSpace)
	for i := 0; i < jumlahUlasan; i++ {
		if u[i].namaCW == namaSpace {
			fmt.Println("Nama pengulas: ", u[i].nama)
			fmt.Println("Komentar: ", u[i].isi)
			fmt.Printf("Rating: %.2f", u[i].rating)
			found = true
		}
	}
	if found == false {
		fmt.Println("Belum ada ulasan")
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
