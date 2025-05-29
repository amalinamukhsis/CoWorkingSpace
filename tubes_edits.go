/*
Topik 8. Aplikasi Manajemen dan Review Co-Working Space
Aplikasi ini digunakan untuk mengelola data co-working space dan
membantu pengguna dalam mencari tempat kerja yang sesuai
berdasarkan fasilitas, lokasi, dan ulasan. Data utama yang digunakan
adalah daftar co-working space, fasilitas yang disediakan, harga sewa,
serta rating pengguna. Pengguna aplikasi adalah pemilik co-working
space serta pekerja remote atau freelancer yang mencari tempat kerja.

Anggota Tim:
- Amalina Mukhsis
- Raniah Hasna Fadiyah

a. Sudah membuat subprogram untuk menginput data, cetak data, searching (seq & binary)
sudah membuat subprogram untuk edit, hapus data
b. Ulasan tidak langsung tercetak saat selesai input, menyimpannya ke struktur CoWorkSpace
c. Sorting berdasar harga dan rating belum dibuat
*/
package main

import "fmt"

const MAX_COWORKSPACE int = 50
const MAX_FASILITAS int = 10
const MAX_KOMENTAR int = 99
const MAX_ULASAN int = 999

/*
deklarasi tipe bentukan CoWorkSpace yang akan menyimpan data-data co-working space
seperti nama(string), lokasi(string), array fasilitas, harga sewa(real),
rating total(real), serta panjang elemen array fasilitas
*/
type CoWorkSpace struct {
	ID                   string
	nama, lokasi         string
	fasilitas            [MAX_FASILITAS]string
	harga_sewa, rating   float64
	lenFasilitas, length int
}
type CoWorkS [MAX_COWORKSPACE]CoWorkSpace

/*
deklarasi tipe bentukan Ulasan utk menyimpan data-data ulasan
seperti ID, coWorkingID, username, komentar, panjang elemen array komentar,
serta rating dari pengguna
*/
type Ulasan struct {
	ID, coWorkingID     int
	username            string
	komentar            [MAX_KOMENTAR]string
	rating, sumRating   float64
	lenKomentar, length int
}
type arrUlasan [MAX_ULASAN]Ulasan

func main() {
	// deklarasi variabel yang dibutuhkan
	var cws CoWorkS
	var ulas arrUlasan
	var nData int
	var pilih, pilPemilik, pilPekerja, pilOpsiUlasan int
	var aksesKe int
	var indexCws int
	var keyNama, keyLokasi, keyFasilitas string
	var valid, back, back2, dataTerurut bool

	dataTerurut = false
	nData = 0
	valid = false
	aksesKe = 1

	for !valid {
		back = false
		mainMenu(&pilih)
		// pengecekan pilihan pengguna
		if pilih == 3 { // pilih exit
			valid = true
		} else if pilih == 1 { // masuk sebagai pemilik
			for !back {
				menuPemilik(&pilPemilik)
				if pilPemilik == 1 { // nambahin data
					menuTambahTempat()
					bacaData(&cws, &nData)
					cetakData(cws, ulas, nData)
				} else if pilPemilik == 2 { // ubah data
					ubahData(&cws)
					cetakData(cws, ulas, nData)
				} else if pilPemilik == 3 { // hapus data
					hapusData(&cws, &nData)
					cetakData(cws, ulas, nData)
				} else {
					back = true
				}
			}
		} else { // masuk sebagai pekerja / pengguna co-working space
			if aksesKe == 1 {
				DataKosong()
				fmt.Println("Masukkan data Co-working space dahulu!! [pilih 1. Pemilik]")
			} else {
				for !back {
					menuPekerja(&pilPekerja)
					if pilPekerja == 7 { // memilih exit
						back = true
					} else if pilPekerja == 1 { // tambah ulasan
						cetakData(cws, ulas, nData)
						back2 = false
						for !back2 {
							menuUlasan(&pilOpsiUlasan)
							if pilOpsiUlasan == 4 { // memilih exit
								back2 = true
							} else if pilOpsiUlasan == 1 { // tambah ulasan
								menuTambahUlasan()
								bacaUlasan(&ulas)
								cetakData(cws, ulas, nData)
							} else if pilOpsiUlasan == 2 { // edit/ubah ulasan
								editUlasan(&ulas, nData)
								cetakData(cws, ulas, nData)
							} else { // hapus ulasan
								hapusUlasan(&ulas, &nData)
								cetakData(cws, ulas, nData)
							}
						}
					} else if pilPekerja == 2 { // cari co-working space berdasar nama (binary search)
						fmt.Print("Masukkan nama co-working space yang anda cari: ")
						fmt.Scan(&keyNama)
						indexCws = CariNama(cws, nData, keyNama, dataTerurut) // cari index nya
						if indexCws == -1 {                                   // ga ketemu yg key nya
							// co-working space tidak ada
							DataKosong()
						} else if indexCws == -2 {
							// data belum terurut
							fmt.Println("Data belum terurut berdasarkan nama!")
						} else { // keynya ketemu
							// co-working space ada
							fmt.Println("Data ditemukan!")
							printDataSatuan(cws, indexCws, ulas)
						}
					} else if pilPekerja == 3 { // cari co-working space berdasar lokasi (seq search)
						// cari co-working space berdasarkan lokasi (sequential search)
						fmt.Print("Masukkan lokasi co-working space yang anda cari: ")
						fmt.Scan(&keyLokasi)
						indexCws = CariLokasi(cws, nData, keyLokasi)
						if indexCws == -1 {
							// co-working space tidak ada
							DataKosong()
						} else {
							// co-working space ada
							fmt.Println("Data ditemukan!")
							printDataSatuan(cws, indexCws, ulas)
						}
					} else if pilPekerja == 4 {
						// sortSewaRendah() selection
					} else if pilPekerja == 5 {
						// sortRatingTinggi() insertion
					} else if pilPekerja == 6 { // cari berdasar fasilitas
						menuCariFasilitas(&keyFasilitas)
						cariFasilitas(cws, nData, keyFasilitas, ulas)
					} else { // mengurutkan data berdasar nama cws nya
						UrutNama(&cws, nData)
						dataTerurut = true
					}
				}
			}
		}
		aksesKe++
	}
}

func editUlasan(arr *arrUlasan, n int) {
	var user string
	fmt.Print("Masukkan username yang ingin diedit ulasannya: ")
	fmt.Scan(&user)

	var edit bool
	var i, j int
	edit = false
	for i = 0; i < n; i++ {
		if arr[i].username == user {
			edit = true

			// Tampilkan ulasan sebelumnya
			fmt.Println("\nUlasan Sebelumnya: ")
			fmt.Println("Username:", arr[i].username)
			fmt.Print("Komentar: ")
			for j = 0; j < 15; j++ {
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
			for j = 0; j < 15; j++ {
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
			for j = 0; j < 15; j++ {
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

func bacaUlasan(ulas *arrUlasan) {
	var IdCws int
	var i, j, n int
	var sum float64 = 0
	var idUlas int = 1
	fmt.Scan(&IdCws)
	fmt.Print("Banyak ulasan yang akan diinput [maks 20, akhiri dengan #]: ")
	fmt.Scan(&n)

	ulas[i].length = 1
	i = IdCws - 1
	ulas[i].ID = idUlas
	ulas[i].coWorkingID = IdCws
	for x := 0; x < n; x++ {
		fmt.Printf("\nUlasan Co-Working Space dengan ID:%d\n", IdCws)
		fmt.Printf("\nUlasan ke-%d\n", i+1)
		ulas[i].ID = IdCws
		fmt.Printf("ID Ulasan: %d\n", ulas[i].ID)
		fmt.Print("Username: ")
		fmt.Scan(&ulas[i].username)

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
		fmt.Printf("| %-36s |\n", "Jika telah selesai memilih kategori")
		fmt.Printf("| %-36s |\n", "diakhiri dengan #")
		fmt.Printf("| %-36s |\n", "EX: Nyaman Satpam_ramah #")
		fmt.Println("+======================================+")
		fmt.Print("Komentar: ")
		j = 0
		ulas[i].lenKomentar = 1
		for j < 15 {
			var komentar string
			fmt.Scan(&komentar)
			if komentar == "#" {
				j = 15
			} else {
				ulas[i].komentar[j] = komentar
				ulas[i].lenKomentar++
				j++
			}
		}
		fmt.Scan(&ulas[i].rating)
		sum = sum + ulas[i].rating
		ulas[i].length++
	}
	ulas[i].sumRating = sum
}

// func cetakUlasan() {

// }

func rerataRating(cws *CoWorkS, ulas arrUlasan, i int) float64 {
	cws[i].rating = ulas[i].sumRating / float64(ulas[i].length)
	return cws[i].rating
}

func hapusData(cws *CoWorkS, n *int) {
	var dataKe, i int
	fmt.Println("+======================================================+")
	fmt.Printf("|%5s%s%5s|\n", "", "Menghapus data co-working space ke berapa?", "")
	fmt.Println("+======================================================+")
	fmt.Print("Co-working space: ")
	fmt.Scan(&dataKe)
	for i = dataKe - 1; i < *n; i++ {
		cws[i] = cws[i+1]
	}
	*n--
	cws[i].length--
}

func ubahData(cws *CoWorkS) {
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

func cariFasilitas(cws CoWorkS, nData int, keyWord string, ulas arrUlasan) {
	var found bool = false
	var i, j, idx int
	for i = 0; i < nData; i++ {
		for j = 0; j < 10; j++ {
			if keyWord == cws[i].fasilitas[j] {
				fmt.Println("Data ditemukan!")
				idx = i
				printDataSatuan(cws, idx, ulas)
				found = true
			}
		}
	}
	if !found {
		fmt.Println("Fasilitas tidak ditemukan.")
	}
}

func printDataSatuan(cws CoWorkS, idx int, ulas arrUlasan) {
	var j int
	var rating float64
	fmt.Printf("\nData Co-Working Space %d", idx+1)
	fmt.Printf("\nID Co-Working Space: %s", cws[idx].ID)
	fmt.Printf("\nNama tempat: %s", cws[idx].nama)

	fmt.Printf("\nLokasi: %s", cws[idx].lokasi)
	fmt.Print("\nFasilitas: ")
	for j = 0; j < cws[idx].lenFasilitas-1; j++ {
		fmt.Print(cws[idx].fasilitas[j])
		if j < cws[idx].lenFasilitas-1 {
			fmt.Print(", ")
		}
	}
	fmt.Printf("\nHarga sewa /bulan: Rp.%.0f", cws[idx].harga_sewa)

	fmt.Print("\nUlasan: ")
	rating = rerataRating(&cws, ulas, idx)
	fmt.Printf("\nRating: %.1f/5\n", rating)
	fmt.Println()
}

func CariNama(cws CoWorkS, n int, key string, dataTerurut bool) int { // binary search A-Z mencari nama
	var left, mid, right, idx int
	if dataTerurut {
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
	} else {
		return -2
	}
}

func UrutNama(cws *CoWorkS, n int) { // sorting nama ascending (a-z)
	var temp CoWorkSpace
	var i, j, min int
	for i = 0; i < n; i++ {
		min = i
		for j = i + 1; j < n-1; j++ {
			if cws[j].nama < cws[min].nama {
				min = j
			}
		}
		temp = cws[min]
		cws[min] = cws[i]
		cws[i] = temp
	}
	fmt.Println("Data telah terurut berdasarkan nama!")
}

func CariLokasi(cws CoWorkS, n int, key string) int {
	var idx, left, right, mid int
	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if cws[mid].lokasi == key {
			idx = mid
		} else if cws[mid].lokasi < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
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

func mainMenu(pilih *int) {
	fmt.Println("+========== Coseum [Platform Co-Working Space] ==========+")
	fmt.Printf("|  %s  |\n", "Selamat datang di Coseum [Penyedia Co-Working Space]")
	fmt.Printf("|%21s%s%21s|\n", "", "Masuk Sebagai:", "")
	fmt.Printf("| %-54s |\n", "1. Pemilik Co-Working Space")
	fmt.Printf("| %-54s |\n", "2. Pengguna Co-Working Space (Pekerja)")
	fmt.Printf("| %-54s |\n", "3. exit")
	fmt.Println("+========================================================+")
	fmt.Print("Pilih (1/2/3): ")
	fmt.Scan(pilih)
}

func menuCariFasilitas(kunci *string) {
	fmt.Println("+======================================================+")
	fmt.Printf("|%5s%s%5s|\n", "", "Fasilitas apa yang anda cari?", "")
	fmt.Println("+======================================================+")
	menuPilihanFasilitas()
	fmt.Print("Fasilitas: ")
	fmt.Scan(kunci)
}

func menuUlasan(pilih *int) {
	fmt.Println("+========== Coseum [Platform Co-Working Space] ==========+")
	fmt.Printf("|%13s%s%13s|\n", "", "Pilih yang ingin anda lakukan:", "")
	fmt.Printf("| %-54s |\n", "1. Tambah Ulasan Co-Working Space")
	fmt.Printf("| %-54s |\n", "2. Ubah Ulasan Co-Working Space")
	fmt.Printf("| %-54s |\n", "3. Hapus Ulasan Co-Working Space")
	fmt.Printf("| %-54s |\n", "4. exit")
	fmt.Println("+========================================================+")
	fmt.Print("Pilih (1/2/3/4): ")
	fmt.Scan(pilih)
}

func menuPekerja(pilih *int) {
	fmt.Println("+========== Coseum [Platform Co-Working Space] ==========+")
	fmt.Printf("|  %s  |\n", "Selamat datang di Coseum [Penyedia Co-Working Space]")
	fmt.Printf("|%13s%s%13s|\n", "", "Pilih yang ingin anda lakukan:", "")
	fmt.Printf("| %-54s |\n", "1. Tambah Ulasan Co-Working Space")
	fmt.Printf("| %-54s |\n", "2. Cari berdasarkan Nama(Binary)")
	fmt.Printf("| %-54s |\n", "3. Cari berdasarkan Lokasi(Sequential)")
	fmt.Printf("| %-54s |\n", "4. Urutkan berdasarkan Harga sewa Terrendah")
	fmt.Printf("| %-54s |\n", "5. Urutkan berdasarkan Rating Tertinggi")
	fmt.Printf("| %-54s |\n", "6. Tampilkan data berdasar fasilitas")
	fmt.Printf("| %-54s |\n", "7. Urutkan berdasarkan Nama Co-Working Space")
	fmt.Printf("| %-54s |\n", "8. exit")
	fmt.Println("+========================================================+")
	fmt.Print("Pilih (1/2/3/4/5/6/7/8): ")
	fmt.Scan(pilih)
}

func menuPemilik(pilih *int) {
	fmt.Println("+========== Coseum [Platform Co-Working Space] ==========+")
	fmt.Printf("|%13s%s%13s|\n", "", "Pilih yang ingin anda lakukan:", "")
	fmt.Printf("| %-54s |\n", "1. Tambah Data Co-Working Space")
	fmt.Printf("| %-54s |\n", "2. Ubah Data Co-Working Space")
	fmt.Printf("| %-54s |\n", "3. Hapus Data Co-Working Space")
	fmt.Printf("| %-54s |\n", "4. exit")
	fmt.Println("+========================================================+")
	fmt.Print("Pilih (1/2/3/4): ")
	fmt.Scan(pilih)
}

func menuTambahTempat() {
	fmt.Println("+======================================================+")
	fmt.Printf("|%5s%s%5s|\n", "", "Silakan masukkan data co-working space anda!", "")
	fmt.Println("+======================================================+")
}

func menuTambahUlasan() {
	fmt.Println("+========================================================+")
	fmt.Printf("|%7s%s%7s|\n", "", "Tambah Ulasan untuk Co-Working Space berapa?", "")
	fmt.Println("+========================================================+")
	fmt.Print("Ulasan Co-Working Space Nomor: ")
}

func bacaData(cws *CoWorkS, n *int) {
	var i, nFasil int
	var id, nama, lokasi string
	var fasilitas [MAX_FASILITAS]string
	var harga float64
	var end bool = false

	fmt.Printf("\nCo-Working Space %d\n", *n+1)
	fmt.Print("ID Co-Working Space: ")
	fmt.Scan(&id)

	fmt.Print("Nama Co-Working Space\n[Gunakan _ sebagai pengganti spasi]: ")
	fmt.Scan(&nama)

	fmt.Print("Lokasi Co-Working Space\n[Gunakan _ sebagai pengganti spasi]: ")
	fmt.Scan(&lokasi)

	fmt.Println("Masukkan fasilitas maksimal 10 / ketik # untuk berhenti")
	menuPilihanFasilitas()
	i = 0
	nFasil = 0
	for !end {
		fmt.Scan(&fasilitas[i])
		nFasil++
		end = fasilitas[i] == "#"
		i++
	}

	fmt.Print("Harga Sewa Co-Working Space /bulan: Rp.")
	fmt.Scan(&harga)

	if *n < MAX_COWORKSPACE {
		cws[*n].ID = id
		cws[*n].nama = nama
		cws[*n].lokasi = lokasi
		for i = 0; i < nFasil-1; i++ {
			cws[*n].fasilitas[i] = fasilitas[i]
		}
		cws[*n].lenFasilitas = nFasil
		cws[*n].harga_sewa = harga
		*n++
		fmt.Println("Co-Working Space berhasil ditambahkan!")
	} else {
		fmt.Println("Co-Working Space gagal ditambahkan!")
	}
}

func menuPilihanFasilitas() {
	fmt.Println("+======================================+")
	fmt.Printf("| %-36s |\n", "Komentar: isi dengan kategori dibawah ini:")
	fmt.Printf("| %-36s |\n", "1. Wifi")
	fmt.Printf("| %-36s |\n", "2. Meeting_room")
	fmt.Printf("| %-36s |\n", "3. Pantry")
	fmt.Printf("| %-36s |\n", "4. AC")
	fmt.Printf("| %-36s |\n", "5. Toilet")
	fmt.Printf("| %-36s |\n", "6. Mushola")
	fmt.Printf("| %-36s |\n", "7. Whiteboard")
	fmt.Printf("| %-36s |\n", "8. Meja_Kursi")
	fmt.Printf("| %-36s |\n", "9. Proyektor")
	fmt.Printf("| %-36s |\n", "10. Printer")
	fmt.Printf("| %-36s |\n", "Jika telah selesai memilih kategori")
	fmt.Printf("| %-36s |\n", "akhiri dengan #")
	fmt.Printf("| %-36s |\n", "EX: AC Meja_Kursi #")
	fmt.Println("+======================================+")
}

func cetakData(cws CoWorkS, ulas arrUlasan, n int) {
	var i int
	var rating float64
	if n != 0 {
		fmt.Println("\nData Co-Working Space")
		for i = 0; i < n; i++ {
			fmt.Printf("ID Co-Working Space: %s", cws[i].ID)
			fmt.Printf("\nNama tempat: %s", cws[i].nama)

			fmt.Printf("\nLokasi: %s", cws[i].lokasi)
			fmt.Print("\nFasilitas: ")
			for j := 0; j < cws[i].lenFasilitas-1; j++ {
				fmt.Print(cws[i].fasilitas[j])
				if j < cws[i].lenFasilitas-1 {
					fmt.Print(", ")
				}
			}
			fmt.Printf("\nHarga sewa /bulan: Rp.%.0f", cws[i].harga_sewa)

			fmt.Print("\nUlasan: ")
			// cetakUlasan()
			rating = rerataRating(&cws, ulas, i)
			fmt.Printf("\nRating: %.1f/5\n", rating)
			fmt.Println()
		}
	} else {
		fmt.Println("Data Kosong!")
	}
}
