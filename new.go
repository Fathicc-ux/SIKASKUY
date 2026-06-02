package main

import "fmt"

type Pembayaran struct {
	nominal int
	waktu   string
}

type Mahasiswa struct {
	nama        string
	nim         string
	statusbayar bool
	tunggakan   int
	jumlahbayar int
	riwayat     [100]Pembayaran
	jmlriwayat  int
}

type dataMahasiswa struct {
	data         [100]Mahasiswa
	jmlMahasiswa int
}

var DB dataMahasiswa

func main() {
	for {
		fmt.Println()
		fmt.Println("============================================================================")
		fmt.Println("                          SISTEM KAS MAHASISWA")
		fmt.Println("============================================================================")
		fmt.Println(" Pilih Menu:")
		fmt.Println(" 1. Kelola Data Mahasiswa")
		fmt.Println(" 2. Bayar Kas")
		fmt.Println(" 3. Cari Mahasiswa")
		fmt.Println(" 0. Keluar")
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Print(" Pilih: ")

		var menu int
		fmt.Scanln(&menu)
		fmt.Println()

		switch menu {
		case 1:
			menuMahasiswa(&DB)
		case 2:
			bayar(&DB)
		case 3:
			menucari(&DB)
		case 0:
			fmt.Println("============================================================================")
			fmt.Println("\n                 Terima kasih telah menggunakan sistem ini.")
			fmt.Println("============================================================================")
			fmt.Println()
			return
		default:
			fmt.Println(" Pilihan tidak valid. Silakan ulangi.")
			fmt.Println("============================================================================")
			fmt.Println()
		}
	}
}

func menuMahasiswa(A *dataMahasiswa) {
	for {
		fmt.Println()
		fmt.Println("============================================================================")
		fmt.Println("                         KELOLA DATA MAHASISWA")
		fmt.Println("============================================================================")
		fmt.Println(" Pilih Menu:")
		fmt.Println(" 1. Tambah Mahasiswa")
		fmt.Println(" 2. Ubah Mahasiswa")
		fmt.Println(" 3. Tampilkan Data Mahasiswa")
		fmt.Println(" 4. Hapus Mahasiswa")
		fmt.Println(" 0. Kembali")
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Print(" Pilih : ")

		var menu int
		fmt.Scanln(&menu)
		fmt.Println()

		switch menu {
		case 1:
			tambah(A)
		case 2:
			update(A)
		case 3:
			tampil(A)
		case 4:
			hapus(A)
		case 0:
			return
		default:
			fmt.Println(" Pilihan tidak valid. Silakan ulangi.")
			fmt.Println("============================================================================")
			fmt.Println()
		}
	}
}

func tambah(A *dataMahasiswa) {
	fmt.Println()
	fmt.Println("============================================================================")
	fmt.Println("                           TAMBAH MAHASISWA")
	fmt.Println("----------------------------------------------------------------------------")
	if A.jmlMahasiswa <= 100 {
		var mhs Mahasiswa

		fmt.Print(" Masukkan NIM Mahasiswa  : ")
		fmt.Scanln(&mhs.nim)
		fmt.Print(" Masukkan Nama Mahasiswa : ")
		fmt.Scanln(&mhs.nama)
		fmt.Println()

		if mhs.nim == "" || mhs.nama == "" {
			fmt.Println(" ERROR: NIM dan Nama tidak boleh kosong.")
			fmt.Println("============================================================================")
			fmt.Println()
			return
		}

		var cek int
		sequential(A, mhs.nim, &cek)
		if cek != -1 {
			fmt.Println(" ERROR: NIM sudah terdaftar.")
			fmt.Println("============================================================================")
			fmt.Println()
			return
		}

		mhs.tunggakan = 50000
		mhs.statusbayar = false
		mhs.jumlahbayar = 0
		A.data[A.jmlMahasiswa] = mhs
		A.jmlMahasiswa++
		fmt.Println(" BERHASIL: Mahasiswa berhasil ditambahkan.")
		fmt.Println("============================================================================")
		fmt.Println()
	} else {
		fmt.Println(" ERROR: Kapasitas database penuh.")
		fmt.Println("============================================================================")
		fmt.Println()
	}
}

func update(A *dataMahasiswa) {
	fmt.Println()
	fmt.Println("============================================================================")
	fmt.Println("                           UPDATE MAHASISWA")
	fmt.Println("----------------------------------------------------------------------------")
	if A.jmlMahasiswa == 0 {
		fmt.Println(" Database mahasiswa kosong.")
		fmt.Println("============================================================================")
		fmt.Println()
		return
	}

	var nim string
	var indeks, cek int

	fmt.Print(" Masukkan NIM mahasiswa : ")
	fmt.Scanln(&nim)
	fmt.Println()

	if nim == "" {
		fmt.Println(" ERROR: NIM tidak boleh kosong.")
		fmt.Println("============================================================================")
		fmt.Println()
		return
	}

	sequential(A, nim, &indeks)
	if indeks == -1 {
		fmt.Println(" ERROR: Mahasiswa tidak ditemukan.")
		fmt.Println("============================================================================")
		fmt.Println()
		return
	}

	fmt.Print(" Masukkan NIM Baru  : ")
	fmt.Scanln(&A.data[indeks].nim)
	fmt.Print(" Masukkan Nama Baru : ")
	fmt.Scanln(&A.data[indeks].nama)
	fmt.Println()

	sequential(A, A.data[indeks].nim, &cek)
	if cek != -1 && cek != indeks {
		fmt.Println(" ERROR: NIM baru sudah digunakan oleh mahasiswa lain.")
		fmt.Println("============================================================================")
		fmt.Println()
	} else if A.data[indeks].nim == "" && A.data[indeks].nama == "" {
		fmt.Println(" Tidak ada perubahan data.")
		fmt.Println("============================================================================")
		fmt.Println()
	} else {
		fmt.Println(" BERHASIL: Data mahasiswa berhasil diubah.")
		fmt.Println("============================================================================")
		fmt.Println()
	}
}

func tampil(A *dataMahasiswa) {
	fmt.Println()
	fmt.Println("============================================================================")
	fmt.Println("                           DAFTAR MAHASISWA")
	fmt.Println("----------------------------------------------------------------------------")
	if A.jmlMahasiswa == 0 {
		fmt.Println(" Database mahasiswa kosong.")
		fmt.Println("============================================================================")
		fmt.Println()
		return
	}
	// Header tabel dengan lebar kolom tetap
	fmt.Printf("%-3s | %-12s | %-23s | %-12s | %-10s\n", "No", "NIM", "Nama", "Total Bayar", "Status")
	fmt.Println("----+--------------+-------------------------+--------------+---------------")
	for i := 0; i < A.jmlMahasiswa; i++ {
		status := "Belum Lunas"
		if A.data[i].statusbayar {
			status = "Lunas"
		}
		fmt.Printf("%-3d | %-12s | %-23s | Rp %-9d | %-10s\n",
			i+1, A.data[i].nim, A.data[i].nama, A.data[i].jumlahbayar, status)
		fmt.Println("----+--------------+-------------------------+--------------+---------------")
	}
	fmt.Println("============================================================================")
	fmt.Println()
}

func hapus(A *dataMahasiswa) {
	fmt.Println()
	fmt.Println("============================================================================")
	fmt.Println("                           HAPUS MAHASISWA")
	fmt.Println("----------------------------------------------------------------------------")
	if A.jmlMahasiswa == 0 {
		fmt.Println(" Database mahasiswa kosong.")
		fmt.Println("============================================================================")
		fmt.Println()
		return
	}

	var Nim string
	var indeks int

	fmt.Print(" Masukkan NIM Mahasiswa : ")
	fmt.Scanln(&Nim)
	fmt.Println()

	if Nim == "" {
		fmt.Println(" ERROR: NIM tidak boleh kosong.")
		fmt.Println("============================================================================")
		fmt.Println()
		return
	}

	sequential(A, Nim, &indeks)
	if indeks == -1 {
		fmt.Println(" ERROR: Mahasiswa tidak ditemukan.")
		fmt.Println("============================================================================")
		fmt.Println()
		return
	} else {
		for i := indeks; i < A.jmlMahasiswa-1; i++ {
			A.data[i] = A.data[i+1]
		}
		A.jmlMahasiswa--
		fmt.Println(" BERHASIL: Data mahasiswa berhasil dihapus.")
		fmt.Println("============================================================================")
		fmt.Println()
	}
}

func bayar(A *dataMahasiswa) {
	fmt.Println()
	fmt.Println("============================================================================")
	fmt.Println("                            PEMBAYARAN KAS")
	fmt.Println("----------------------------------------------------------------------------")
	if A.jmlMahasiswa == 0 {
		fmt.Println(" Database mahasiswa kosong.")
		fmt.Println("============================================================================")
		fmt.Println()
		return
	}

	var nim string
	var indeks int

	fmt.Print(" NIM mahasiswa : ")
	fmt.Scanln(&nim)
	fmt.Println()

	sequential(A, nim, &indeks)
	if indeks == -1 {
		fmt.Println(" ERROR: Mahasiswa tidak ditemukan.")
		fmt.Println("============================================================================")
		fmt.Println()
		return
	} else {
		var bayar Pembayaran
		fmt.Print(" Tanggal (dd/mm/yyyy) : ")
		fmt.Scanln(&bayar.waktu)
		fmt.Print(" Nominal pembayaran   : ")
		fmt.Scanln(&bayar.nominal)
		fmt.Println()

		//jika nominal melebihi tunggakan
		if bayar.nominal > A.data[indeks].tunggakan {
			kembalian := bayar.nominal - A.data[indeks].tunggakan
			fmt.Printf(" PERINGATAN: Nominal melebihi tunggakan (Rp %d). Pembayaran disesuaikan menjadi Rp %d. Kembalian Rp %d\n",
				A.data[indeks].tunggakan, A.data[indeks].tunggakan, kembalian)
			bayar.nominal = A.data[indeks].tunggakan
			fmt.Println()
		}

		//proses pembayaran
		A.data[indeks].tunggakan -= bayar.nominal
		A.data[indeks].jumlahbayar += bayar.nominal

		//proces simpan riwayat
		if A.data[indeks].jmlriwayat < 100 {
			A.data[indeks].riwayat[A.data[indeks].jmlriwayat] = Pembayaran{
				nominal: bayar.nominal,
				waktu:   bayar.waktu,
			}
			A.data[indeks].jmlriwayat++
		}

		//update status
		if A.data[indeks].tunggakan == 0 {
			A.data[indeks].statusbayar = true
			fmt.Println(" PEMBAYARAN LUNAS! Status menjadi LUNAS.")
			fmt.Println()
		} else {
			fmt.Printf(" Pembayaran berhasil. Sisa tunggakan : Rp %d\n", A.data[indeks].tunggakan)
			fmt.Println()
		}
	}
}

func menucari(A *dataMahasiswa) {
	for {
		fmt.Println()
		fmt.Println("============================================================================")
		fmt.Println("                            CARI MAHASISWA")
		fmt.Println("============================================================================")
		fmt.Println(" 1. Pencarian Sekuensial (Sequential)")
		fmt.Println(" 2. Pencarian Biner (Binary) - data akan diurut")
		fmt.Println(" 0. Kembali")
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Print(" Pilih : ")

		var pilih int
		fmt.Scanln(&pilih)
		fmt.Println()

		switch pilih {
		case 1:
			cari_sequential(A)
		case 2:
			cari_binary(A)
		case 0:
			return
		default:
			fmt.Println(" Pilihan tidak valid.")
			fmt.Println("============================================================================")
			fmt.Println()
		}
	}
}

//================SEQUENTIAL=====================
func sequential(A *dataMahasiswa, Nim string, indeks *int) {
	*indeks = -1
	for i := 0; i < A.jmlMahasiswa; i++ {
		if A.data[i].nim == Nim {
			*indeks = i
		}
	}
}

func sequentialNama(A *dataMahasiswa, Nama string, indeks *int) {
	*indeks = -1
	for i := 0; i < A.jmlMahasiswa; i++ {
		if A.data[i].nama == Nama {
			*indeks = i
		}
	}
}

func cari_sequential(A *dataMahasiswa) {
	if A.jmlMahasiswa == 0 {
		fmt.Println("============================================================================")
		fmt.Println("Database mahasiswa kosong.")
		fmt.Println("============================================================================")
		fmt.Println()
		return
	}

	var pilih int
	valid := false
	for !valid {
		fmt.Println()
		fmt.Println("============================================================================")
		fmt.Println("                       PENCARIAN SEQUENSIAL")
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Println(" Pilih Menu :")
		fmt.Println(" 1. NIM")
		fmt.Println(" 2. Nama")
		fmt.Println(" 0. Kembali")
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Print(" Pilih : ")
		fmt.Scanln(&pilih)
		fmt.Println()

		switch pilih {
		case 0:
			return
		case 1, 2:
			valid = true
		default:
			fmt.Println(" Pilihan tidak valid. Silakan ulangi.")
			fmt.Println("============================================================================")
			fmt.Println()
		}
	}

	var indeks int
	switch pilih {
	case 1:
		var nim string
		fmt.Print(" Masukkan NIM : ")
		fmt.Scanln(&nim)
		fmt.Println()
		if nim == "" {
			fmt.Println(" NIM tidak boleh kosong.")
			fmt.Println("============================================================================")
			fmt.Println()
		}
		sequential(A, nim, &indeks)

	case 2:
		var nama string
		fmt.Print(" Masukkan Nama : ")
		fmt.Scanln(&nama)
		fmt.Println()
		if nama == "" {
			fmt.Println(" Nama tidak boleh kosong.")
			fmt.Println("============================================================================")
			fmt.Println()
		}
		sequentialNama(A, nama, &indeks)
	}

	if indeks == -1 {
		fmt.Println(" Hasil: Mahasiswa tidak ditemukan.")
		fmt.Println("============================================================================")
		fmt.Println()
	} else {
		tampilkanDetailMahasiswa(A, indeks)
	}
}

//=================BINARY======================
func binary(A *dataMahasiswa, nim string) int {
	kiri := 0
	kanan := A.jmlMahasiswa - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if A.data[tengah].nim == nim {
			return tengah
		} else if A.data[tengah].nim < nim {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func binaryNama(A *dataMahasiswa, nama string) int {
	kiri := 0
	kanan := A.jmlMahasiswa - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if A.data[tengah].nama == nama {
			return tengah
		} else if A.data[tengah].nama < nama {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func cari_binary(A *dataMahasiswa) {
	if A.jmlMahasiswa == 0 {
		fmt.Println("============================================================================")
		fmt.Println("Database mahasiswa kosong.")
		fmt.Println("============================================================================")
		fmt.Println()
		return
	}

	var pilih int
	valid := false
	for !valid {
		fmt.Println()
		fmt.Println("============================================================================")
		fmt.Println("                         PENCARIAN BINERY")
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Println(" Pilih Menu :")
		fmt.Println(" 1. NIM")
		fmt.Println(" 2. Nama")
		fmt.Println(" 0. Kembali")
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Print(" Pilih : ")
		fmt.Scanln(&pilih)
		fmt.Println()

		switch pilih {
		case 0:
			return
		case 1, 2:
			valid = true
		default:
			fmt.Println(" Pilihan tidak valid. Silakan ulangi.")
			fmt.Println("============================================================================")
			fmt.Println()
		}
	}

	var indeks int
	switch pilih {
	case 1:
		var nim string
		fmt.Print(" Masukkan NIM : ")
		fmt.Scanln(&nim)
		fmt.Println()
		if nim == "" {
			fmt.Println(" NIM tidak boleh kosong.")
			fmt.Println("============================================================================")
			fmt.Println()
		}
		sequential(A, nim, &indeks)
		binary(A, nim)

	case 2:
		var nama string
		fmt.Print(" Masukkan Nama : ")
		fmt.Scanln(&nama)
		fmt.Println()
		if nama == "" {
			fmt.Println(" Nama tidak boleh kosong.")
			fmt.Println("============================================================================")
			fmt.Println()
		}
		sequentialNama(A, nama, &indeks)
		binaryNama(A, nama)
	}

	if indeks == -1 {
		fmt.Println(" Hasil: Mahasiswa tidak ditemukan.")
		fmt.Println("============================================================================")
		fmt.Println()
	} else {
		tampilkanDetailMahasiswa(A, indeks)
	}
}

func tampilkanDetailMahasiswa(A *dataMahasiswa, indeks int) {
	fmt.Println("===============================================================================")
	fmt.Println("                            DATA MAHASISWA")
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Printf("%-20s : %s\n", "NIM", A.data[indeks].nim)
	fmt.Printf("%-20s : %s\n", "Nama", A.data[indeks].nama)
	fmt.Printf("%-20s : Rp %d\n", "Tunggakan", A.data[indeks].tunggakan)
	status := "Belum Lunas"
	if A.data[indeks].statusbayar {
		status = "Lunas"
	}
	fmt.Printf("%-20s : %s\n", "Status", status)
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println()

	fmt.Print(" Lihat riwayat pembayaran? (y/n) : ")
	var jawab string
	fmt.Scanln(&jawab)
	fmt.Println()

	if jawab == "y" || jawab == "Y" {
		if A.data[indeks].jmlriwayat == 0 {
			fmt.Println(" Belum ada riwayat pembayaran.")
			fmt.Println()
		} else {
			fmt.Println("-------------------------------------------------------------------------------")
			fmt.Println("                         RIWAYAT PEMBAYARAN")
			fmt.Println("-------------------------------------------------------------------------------")
			fmt.Printf("%-4s %-12s %-12s\n", "No", "Tanggal", "Nominal")
			fmt.Println("----+------------+------------")
			for i := 0; i < A.data[indeks].jmlriwayat; i++ {
				fmt.Printf("%-4d %-12s Rp %-9d\n", i+1, A.data[indeks].riwayat[i].waktu, A.data[indeks].riwayat[i].nominal)
			}
			fmt.Println("-------------------------------------------------------------------------------")
		}
	}
	fmt.Println("===============================================================================")
	fmt.Println()
}
