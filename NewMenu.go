package main

import "fmt"

const NMAX int = 50

type coWorkSpace struct {
	nama                               string
	fasilitas                          [15]string
	lokasi, ulasan                     [NMAX]string
	harga_sewa                         float64
	rating                             float64
	lenLokasi, lenFasilitas, lenUlasan int
}
type coWorkS [NMAX]coWorkSpace

func main() {
	var cws coWorkS
	var nData, n int
	var pilih, pilTambah, aksesKe int
	aksesKe = 1
	for valid := false; !valid; {
		mainMenu()
		fmt.Print("Pilih (1/2/3)? ")
		fmt.Scan(&pilih)
		if pilih == 3 {
			valid = true
		} else if pilih == 1 && aksesKe == 1 {
			fmt.Printf("+===========================+\n")
			fmt.Printf("|%4s%s%4s|\n", " ", "Data belum tersedia", " ")
			fmt.Printf("+===========================+\n")
			menuPemilik()
			fmt.Scan(&nData)
			bacaData(&cws, nData)
		} else if pilih == 1 && aksesKe > 1 {
			cetakData(cws, nData)
		} else {
			menuTambah()
			fmt.Scan(&pilTambah)
			switch pilTambah {
			case 1:
				menuPemilik()
				fmt.Scan(&nData)
				bacaData(&cws, nData)
			case 2:
				cetakData(cws, nData)
				menuUlasan()
				fmt.Scan(&n)
				bacaUlasan(&cws, n-1)
			case 3:
				//ubahData()
				//insertion sort (?)
			case 4:
				//ubahUlasan()
				//insertion sort (/)
			}
		}
		aksesKe++
	}
}

func mainMenu() {
	fmt.Println("+========== Coseum [Platform Co-Working Space] ==========+")
	fmt.Printf("|  %s  |\n", "Selamat datang di Coseum [Penyedia Co-Working Space]")
	fmt.Printf("|%13s%s%13s|\n", "", "Pilih yang ingin anda lakukan:", "")
	fmt.Printf("| %-54s |\n", "1. Tampilkan data co-working space")
	fmt.Printf("| %-54s |\n", "2. Tambahkan data co-working space")
	fmt.Printf("| %-54s |\n", "3. exit")
	fmt.Println("+========================================================+")
}

/*
func menuPekerja() {
	fmt.Println("Ingin mencari co-working space yang bagaimana?")
	fmt.Println("1. Rating tertinggi")
	fmt.Println("2. Fasilitas +Wifi")
	fmt.Println("3. Fasilitas +AC")
	fmt.Println("4. Back")
}*/

func menuPemilik() {
	fmt.Println("+======================================================+")
	fmt.Printf("|%5s%s%5s|\n", "", "Silakan masukkan data co-working space anda!", "")
	fmt.Println("+======================================================+")
	fmt.Print("Banyak co-working space yang akan dimasukkan: ")
}

func menuTambah() {
	fmt.Println("+========================================================+")
	fmt.Printf("|%16s%s%16s|\n", "", "Silakan Pilih Aktivitas!", "")
	fmt.Printf("| %-54s |\n", "1. Tambah Tempat Co-Working Space")
	fmt.Printf("| %-54s |\n", "2. Tambah Ulasan Co-Working Space")
	fmt.Printf("| %-54s |\n", "3. Ubah Data Co-Working Space")
	fmt.Printf("| %-54s |\n", "4. Ubah Ulasan")
	fmt.Println("+========================================================+")
	fmt.Print("Pilih (1/2/3/4): ")
}

func menuUlasan() {
	fmt.Println("+========================================================+")
	fmt.Printf("|%7s%s%7s|\n", "", "Tambah Ulasan untuk Co-Working Space berapa?", "")
	fmt.Println("+========================================================+")
	fmt.Print("Ulasan Co-Working Space Nomor: ")
}

func bacaUlasan(cws *coWorkS, n int) {
	var j int = 0
	var sum float64 = 0
	fmt.Println("+======================================+")
	fmt.Printf("| %-36s |\n", "gunakan _ sebagai pengganti spasi")
	fmt.Printf("| %-36s |\n", "akhiri dengan #")
	fmt.Println("+======================================+")
	fmt.Printf("\nUlasan %d: ", j+1)
	fmt.Scan(&cws[n].ulasan[j])
	for cws[n].ulasan[j] != "#" {
		fmt.Print("Rating: ")
		fmt.Scan(&cws[n].rating)
		sum = sum + cws[n].rating
		j++
		cws[n].lenUlasan++
		fmt.Printf("\nUlasan %d: ", j+1)
		fmt.Scan(&cws[n].ulasan[j])
	}
	cws[n].rating = sum / float64(j)
}

func bacaData(cws *coWorkS, n int) {
	var i, j int
	for i = 0; i < n; i++ {
		cws[i].lenLokasi = 1
		cws[i].lenFasilitas = 1
		cws[i].lenUlasan = 1
		j = 0
		fmt.Printf("\nCo-Working Space %d\n", i+1)
		fmt.Print("Nama Co-Working Space [Gunakan _ sebagai pengganti spasi]: ")
		fmt.Scan(&cws[i].nama)

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
	}
}

func cetakData(cws coWorkS, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("\nData Co-Working Space %d", i+1)
		fmt.Printf("\nNama tempat: %s", cws[i].nama)

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
		fmt.Println()
	}
}
