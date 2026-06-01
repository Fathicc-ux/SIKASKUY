//contoh saja
//baru contoh

package main

import "fmt"

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

func tambah(A *Database) {
	if A.MahasiswaCount < 100 {

		var mhs mahasiswa

		fmt.Print("Nama: ")
		fmt.Scan(&mhs.Nama)

		fmt.Print("NIM: ")
		fmt.Scan(&mhs.NIM)

		mhs.tunggakan = 0
		mhs.statuspembayaraan = false
		A.Mahasiswa[A.MahasiswaCount] = mhs
		A.MahasiswaCount++
		fmt.Println("== Mahasiswa Berhasil Ditambahkan ==")
	} else {
		fmt.Println("== Database Mahasiswa Penuh ==")
	}
}

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

		fmt.Println("== Mahasiswa Berhasil Diupdate ==")
	}
}

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

// func cariMahasiswa(A *Database, nim string) int {
// 	var i int
// 	var indeks int
// 	indeks = -1

// 	for i = 0; i < A.MahasiswaCount; i++ {
// 		if A.Mahasiswa[i].NIM == nim {
// 			indeks = i

// 			i = i + 1
// 		}

// 	}
// 	return indeks
// }

func squential(A *Database, nim string, indeks *int) {
	var i int
	*indeks = -1

	for i = 0; i < A.MahasiswaCount; i++ {
		if A.Mahasiswa[i].NIM == nim {
			*indeks = i
		}
	}
}
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

func tampil(A *Database) {
	var i int
	if A.MahasiswaCount == 0 {
		fmt.Println("== Database Mahasiswa Kosong ==")
	} else {
		fmt.Println("== Daftar Status Pembayaran Mahasiswa ==")

		for i = 0; i < A.MahasiswaCount; i++ {

			fmt.Println("Nama :", A.Mahasiswa[i].Nama)
			fmt.Println("NIM  :", A.Mahasiswa[i].NIM)

			if A.Mahasiswa[i].statuspembayaraan == true {
				fmt.Println("Status Pembayaran: Lunas")
			} else {
				fmt.Println("Status Pembayaran: Belum Lunas")
			}
		}
	}
}

func bayar(A *Database) {
	var nim string
	var indeks int
	var bayar transaksi

	fmt.Print("Masukkan NIM Mahasiswa yang akan melakukan pembayaran: ")
	fmt.Scan(&nim)
	squential(A, nim, &indeks)

	if indeks == -1 {
		fmt.Println("== Mahasiswa Tidak Ditemukan ==")
	} else {
		fmt.Print("Masukan Tanggal Pembayaran (DD/MM/YYYY): ")
		fmt.Scan(&bayar.tanggal)

		fmt.Print("Masukan Nominal Pembayaran: ")
		fmt.Scan(&bayar.nominal)

		A.Mahasiswa[indeks].riwayat[A.Mahasiswa[indeks].jumlahbayar] = bayar
		A.Mahasiswa[indeks].jumlahbayar++
		A.Mahasiswa[indeks].tunggakan = A.Mahasiswa[indeks].tunggakan - bayar.nominal
		A.Mahasiswa[indeks].statuspembayaraan = true
		fmt.Println("== Pembayaran Berhasil Dicatat ==")
	}
}

func main() {
	var pilih int
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
		fmt.Println("8. selectionsortAscending")
		fmt.Println("9. 3selectionsortDescending")
		fmt.Println("10. InsertionsortAscending")
		fmt.Println("11. InsertionsortDescending")
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

			selectionsortAscending(&DB, "Nama")
			tampil(&DB)
		} else if pilih == 9 {

			selectionsortDescending(&DB, "Nama")
			tampil(&DB)
		} else if pilih == 10 {

			insertionsortAscending(&DB, "Nama")
			tampil(&DB)
		} else if pilih == 11 {

			insertionsortDescending(&DB, "Nama")
			tampil(&DB)
		} else if pilih == 0 {

			fmt.Println("Selesai")

		} else {

			fmt.Println("Menu salah")
		}
	}
}

//00
//ubah