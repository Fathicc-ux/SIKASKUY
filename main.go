package main

import ("fmt"
		"time"
)

type mahasiswa struct {
	NIM               string
	Nama              string
	tunggakan         int
	statuspembayaraan bool
	riwayat           [100]transaksi
	jumlahbayar       int
}

type Database struct {
	Mahasiswa      [100]mahasiswa
	MahasiswaCount int
}

type transaksi struct {
	tanggal string
	nominal int
}

var DB Database

//A. Fungsi Tambah
func tambah(A *Database) {
	if A.MahasiswaCount < 100 {
	for {
	var mhs mahasiswa

		fmt.Print("Nama: ")
		fmt.Scan(&mhs.Nama)
		if mhs.Nama == "0" || mhs.Nama == "keluar" {
			return
		}
		fmt.Print("NIM: ")
		fmt.Scan(&mhs.NIM)
		if mhs.	NIM == "0" || mhs.NIM == "keluar" {
			return
		}
		mhs.tunggakan = 5000
		mhs.statuspembayaraan = false
		A.Mahasiswa[A.MahasiswaCount] = mhs
		A.MahasiswaCount++
		fmt.Println("== Mahasiswa Berhasil Ditambahkan ==")
	}
	} else {
		fmt.Println("== Database Mahasiswa Penuh ==")
	}
}

//B. Fungsi Ubah (update)
func ubah(A *Database) {
	var nim string
	var indeks int

	fmt.Print("Masukkan NIM Mahasiswa yang akan diupdate: ")
	fmt.Scan(&nim)
	squential(A, nim, &indeks)

	if indeks == -1 {
		fmt.Println("== Mahasiswa Tidak Ditemukan ==")
		return
	} else {
		fmt.Print("Nama: ")
		fmt.Scan(&A.Mahasiswa[indeks].Nama)

		fmt.Print("NIM: ")
		fmt.Scan(&A.Mahasiswa[indeks].NIM)
		
		for i := 0; i < A.MahasiswaCount; i++ {
		if i != indeks && A.Mahasiswa[i].NIM == A.Mahasiswa[indeks].NIM {

		fmt.Println("!!! NIM sudah digunakan !!!")
		fmt.Print("Masukkan NIM Baru: ")
		fmt.Scan(&A.Mahasiswa[indeks].NIM)
	}
	}
	fmt.Println("== Mahasiswa Berhasil Diupdate ==")
	}
}

//C. Fungsi Hapus
func hapus(A *Database) {
	var nim string
	var indeks int
	var i int

	fmt.Print("Masukkan NIM Mahasiswa yang akan dihapus: ")
	fmt.Scan(&nim)

	squential(A, nim, &indeks)
	if indeks == -1 {
		fmt.Println("== Mahasiswa Tidak Ditemukan ==")

	} else {
		for i = indeks; i < A.MahasiswaCount-1; i++ {
			A.Mahasiswa[i] = A.Mahasiswa[i+1]
		}
		A.MahasiswaCount--
		fmt.Println("== Mahasiswa Berhasil Dihapus ==")
	}
}
//B. Fungsi Bayar ( mencatat nominal iuran yang masuk serta tanggal pembayaran)
func bayar(A *Database) {
	var nim string
	var indeks int
	var bayar transaksi

	fmt.Print("Masukkan NIM Mahasiswa yang akan melakukan pembayaran: ")
	fmt.Scan(&nim)
	squential(A, nim, &indeks)

	if indeks == -1 {
		fmt.Println("== Mahasiswa Tidak Ditemukan ==")
		return
	}
		bayar.tanggal = time.Now().Format("02/01/2006")
		fmt.Print("Masukan Nominal Pembayaran: ")
		fmt.Scan(&bayar.nominal)

		A.Mahasiswa[indeks].riwayat[A.Mahasiswa[indeks].jumlahbayar] = bayar
		A.Mahasiswa[indeks].jumlahbayar++
		A.Mahasiswa[indeks].tunggakan = A.Mahasiswa[indeks].tunggakan - bayar.nominal
		if A.Mahasiswa[indeks].tunggakan <= 0 {
			A.Mahasiswa[indeks].tunggakan = 0
			A.Mahasiswa[indeks].statuspembayaraan = true
			fmt.Println("== Pembayaran Berhasil Dicatat ==")
			// fmt.Println("Status pembayaran: Lunas")
		} else {
			A.Mahasiswa[indeks].statuspembayaraan = false
			fmt.Println("== Pembayaran Berhasil Dicatat ==")
			fmt.Println("Status pembayaran: Belum lunas")
			fmt.Println("Sisa Tunggakan: ", A.Mahasiswa[indeks].tunggakan)
		}
	}

//c. Fungsi (Sequential)
func squential(A *Database, nim string, indeks *int) {
	var i int
	*indeks = -1

	for i = 0; i < A.MahasiswaCount; i++ {
		if A.Mahasiswa[i].NIM == nim {
			*indeks = i
		}
	}
}
//C. Fungsi cari squential sub fungsi squential untuk mencari mahasiswa yang belum bayar berdasarkan NIM dan Nama yang diinputkan
func cari_squential(A *Database) {
	var nim string
	var indeks int

	fmt.Println("====")
	fmt.Print("Masukkan NIM Mahasiswa yang akan dicari: ")
	fmt.Scan(&nim)
	squential(A, nim, &indeks)
	if indeks == -1 {
		fmt.Println("== Mahasiswa Tidak Ditemukan ==")
	} else {
		if A.Mahasiswa[indeks].statuspembayaraan == false {
			fmt.Println("== Daftar Mahasiswa Belum Lunas ==")
			fmt.Println("NIM  :", A.Mahasiswa[indeks].NIM)
			fmt.Println("Nama :", A.Mahasiswa[indeks].Nama)
		} else {
			fmt.Println("Mahasiswa Sudah Membayar secara Lunas")
			fmt.Println("NIM  :", A.Mahasiswa[indeks].NIM)
			fmt.Println("Nama :", A.Mahasiswa[indeks].Nama)
		}
	}
}

//c. Fungsi (Binary)
func binary(A *Database, nim string, indeks *int) {
	var kanan, kiri, tengah int

	*indeks = -1
	kiri = 0
	kanan = A.MahasiswaCount - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if A.Mahasiswa[tengah].NIM == nim {
			*indeks = tengah
			kiri = kanan + 1
		} else if A.Mahasiswa[tengah].NIM < nim {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
}

//C. Fungsi Cari Sequential sub fungsi squential untuk mencari mahasiswa yang belum bayar berdasarkan NIM dan Nama yang diinputkan
func cari_binary(A *Database) {
	var nim string
	var indeks int

	fmt.Print("Masukkan NIM Mahasiswa yang akan dicari: ")
	fmt.Scan(&nim)
	insertionsortNIM(A)
	binary(A, nim, &indeks)
	if indeks == -1 {
		fmt.Println("== Mahasiswa Tidak Ditemukan ==")
	} else {
		fmt.Println("NIM  :", A.Mahasiswa[indeks].NIM)
		fmt.Println("Nama :", A.Mahasiswa[indeks].Nama)
	}
}

//D. Fungsi Insertionsort berdasarkan NIM untuk binary search
func insertionsortNIM(A *Database) {
	var i, j int
	var sementara mahasiswa
	for i = 1; i < A.MahasiswaCount; i++ {
		sementara = A.Mahasiswa[i]
		j = i - 1
		for j >= 0 && A.Mahasiswa[j].NIM > sementara.NIM {
			A.Mahasiswa[j+1] = A.Mahasiswa[j]
			j--
		}
		A.Mahasiswa[j+1] = sementara
	}
}

//D. Fungsi Selectionsort Asscending berdasakan nama dan tunggakan
func selectionsortAscending(A *Database, kategori string) {
	var i, j, indeks int
	var sementara mahasiswa

	for i = 0; i < A.MahasiswaCount-1; i++ {
		indeks = i
		for j = i + 1; j < A.MahasiswaCount; j++ {
			if kategori == "Nama" {
				if A.Mahasiswa[j].Nama < A.Mahasiswa[indeks].Nama {
					indeks = j
				}
			} else if kategori == "Tunggakan" {
				if A.Mahasiswa[j].tunggakan < A.Mahasiswa[indeks].tunggakan {
					indeks = j
				}
			}
		}
		sementara = A.Mahasiswa[i]
		A.Mahasiswa[i] = A.Mahasiswa[indeks]
		A.Mahasiswa[indeks] = sementara
	}
}

//D. Fungsi Selectionsort Descending berdasarkan nama dan tunggakan
func selectionsortDescending(A *Database, kategori string) {
	var i, j, indeks int
	var sementara mahasiswa

	for i = 0; i < A.MahasiswaCount-1; i++ {
		indeks = i
		for j = i + 1; j < A.MahasiswaCount; j++ {
			if kategori == "Nama" {
				if A.Mahasiswa[j].Nama > A.Mahasiswa[indeks].Nama {
					indeks = j
				}
			} else if kategori == "Tunggakan" {
				if A.Mahasiswa[j].tunggakan > A.Mahasiswa[indeks].tunggakan {
					indeks = j
				}
			}
		}
		sementara = A.Mahasiswa[i]
		A.Mahasiswa[i] = A.Mahasiswa[indeks]
		A.Mahasiswa[indeks] = sementara
	}
}

//D. Fungsi insetionsort Ascending berdasarkan nama dan tunggakan
func insertionsortAscending(A *Database, kategori string) {
	var i, j int
	var sementara mahasiswa

	for i = 1; i < A.MahasiswaCount; i++ {
		sementara = A.Mahasiswa[i]
		j = i - 1
		if kategori == "Nama" {
			for j >= 0 && A.Mahasiswa[j].Nama > sementara.Nama {
				A.Mahasiswa[j+1] = A.Mahasiswa[j]
				j--
			}
		} else if kategori == "Tunggakan" {
			for j >= 0 && A.Mahasiswa[j].tunggakan > sementara.tunggakan {
				A.Mahasiswa[j+1] = A.Mahasiswa[j]
				j--
			}
		}
		A.Mahasiswa[j+1] = sementara
	}
}

//D. Fungsi insertionsort Descending berdasarkan nama dan tunggakan
func insertionsortDescending(A *Database, kategori string) {
	var i, j int
	var sementara mahasiswa
	for i = 1; i < A.MahasiswaCount; i++ {
		sementara = A.Mahasiswa[i]
		j = i - 1

		if kategori == "Nama" {
			for j >= 0 && A.Mahasiswa[j].Nama < sementara.Nama {
				A.Mahasiswa[j+1] = A.Mahasiswa[j]
				j--
			}
		} else if kategori == "Tunggakan" {
			for j >= 0 && A.Mahasiswa[j].tunggakan < sementara.tunggakan {
				A.Mahasiswa[j+1] = A.Mahasiswa[j]
				j--
			}
		}
		A.Mahasiswa[j+1] = sementara
	}
}

//Fungsi tampil untuk menampilkan daftar mahasiswa beserta status pembayaran, tunggakan, dan riwayat pembayaran jika ada
func tampil(A *Database) {
	var i, k int

	if A.MahasiswaCount == 0 {
		fmt.Println("== Database Mahasiswa Kosong ==")
		return
	}

	fmt.Println("== Daftar Status Pembayaran Mahasiswa ==")
	fmt.Println()

	for i = 0; i < A.MahasiswaCount; i++ {
		fmt.Println("Nama :", A.Mahasiswa[i].Nama)
		fmt.Println("NIM  :", A.Mahasiswa[i].NIM)
		fmt.Println("Tunggakan :", A.Mahasiswa[i].tunggakan)

		if A.Mahasiswa[i].statuspembayaraan {
			fmt.Println("Status Pembayaran : Lunas")
		} else {
			fmt.Println("Status Pembayaran : Belum Lunas")
		}

		fmt.Println("Riwayat Pembayaran:")

		if A.Mahasiswa[i].jumlahbayar == 0 {
			fmt.Println("Belum ada pembayaran")
		}

		for k = 0; k < A.Mahasiswa[i].jumlahbayar; k++ {
			tr := A.Mahasiswa[i].riwayat[k]
			fmt.Printf("  %d. Tanggal: %s | Nominal: %d\n",
				k+1, tr.tanggal, tr.nominal)
		}

		fmt.Println("--------------------------------")
	}
}

func statistik(A *Database) {
	var i, j int
	var totalKas int
	var jumlahLunas int

	totalKas = 0
	jumlahLunas = 0

	for i = 0; i < A.MahasiswaCount; i++ {

		if A.Mahasiswa[i].statuspembayaraan == true {
			jumlahLunas++
		}

		for j = 0; j < A.Mahasiswa[i].jumlahbayar; j++ {
			totalKas += A.Mahasiswa[i].riwayat[j].nominal
		}
	}
	fmt.Println("STATISTIK KAS")
	fmt.Println("==============================")
	fmt.Println("Total Saldo Kas       :", totalKas)
	fmt.Println("Mahasiswa Lunas       :", jumlahLunas)
	fmt.Println("Total Mahasiswa       :", A.MahasiswaCount)
	fmt.Println("==============================")
}
func main() {
	var pilih int
	var kategori int

	pilih = -1

	for pilih != 0 {

		fmt.Println()
		fmt.Println("===== SIKAS =====")
		fmt.Println("1. Tambah")
		fmt.Println("2. Tampil")
		fmt.Println("3. Ubah")
		fmt.Println("4. Hapus")
		fmt.Println("5. Bayar")
		fmt.Println("6. Cari Sequential")
		fmt.Println("7. Cari Binary")
		fmt.Println("8. Selectionsort")
		fmt.Println("9. Insertionsort")
		fmt.Println("10. Statistik")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {

			tambah(&DB)

		} else if pilih == 2 {

			tampil(&DB)

		} else if pilih == 3 {

			ubah(&DB)

		} else if pilih == 4 {

			hapus(&DB)
		} else if pilih == 5 {

			bayar(&DB)

		} else if pilih == 6 {

			cari_squential(&DB)

		} else if pilih == 7 {

			cari_binary(&DB)

		} else if pilih == 8 {
			fmt.Println("=== Pilih pengurutan ===")
			//var kategori int
			fmt.Println("1. Nama ascending")
			fmt.Println("2. Nama discending")
			fmt.Println("3. Tunggakan ascending")
			fmt.Println("4. Tunggakan discending")
			fmt.Print("Pilih: ")
			fmt.Scan(&kategori)

			if kategori == 1 {
				selectionsortAscending(&DB, "Nama")
			} else if kategori == 2 {
				selectionsortDescending(&DB, "Nama")
			} else if kategori == 3 {
				selectionsortAscending(&DB, "Tunggakan")
			} else if kategori == 4 {
				selectionsortDescending(&DB, "Tunggakan")
			}
			tampil(&DB)
		} else if pilih == 9 {
			fmt.Println("=== Pilih pengurutan ===")
			//var kategori int
			fmt.Println("1. Nama ascending")
			fmt.Println("2. Nama discending")
			fmt.Println("3. Tunggakan ascending")
			fmt.Println("4. Tunggakan discending")
			fmt.Print("Pilih: ")
			fmt.Scan(&kategori)

			if kategori == 1 {
				insertionsortAscending(&DB, "Nama")
			} else if kategori == 2 {
				insertionsortDescending(&DB, "Nama")
			} else if kategori == 3 {
				insertionsortAscending(&DB, "Tunggakan")
			} else if kategori == 4 {
				insertionsortDescending(&DB, "Tunggakan")
			}
			tampil(&DB)
		} else if pilih == 10 {
			statistik(&DB)
		} else if pilih == 0 {

			fmt.Println("Selesai")

		} else {
			fmt.Println("Menu salah")
		}
	}
}

//00
//ubah