package main

import "fmt"

const TARGET_IURAN = 500000

type Transaksi struct {
	Tanggal string
	Nominal int
}

type Mahasiswa struct {
	NIM             string
	Nama            string
	Tunggakan       int
	StatusLunas     bool
	RiwayatBayar    [100]Transaksi
	JumlahTransaksi int
}

type Database struct {
	DataMahasiswa [100]Mahasiswa
	JumlahData    int
}

var db Database

func main() {
	for {
		fmt.Println("============================================================================")
		fmt.Println("                          SISTEM KAS MAHASISWA")
		fmt.Println("============================================================================")
		fmt.Println(" 1. Kelola Data Mahasiswa")
		fmt.Println(" 2. Bayar Kas")
		fmt.Println(" 3. Cari Mahasiswa")
		fmt.Println(" 0. Keluar")
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Print(" Pilihan Anda : ")
		var pilih int
		fmt.Scan(&pilih)
		fmt.Scanln()

		switch pilih {
		case 1:
			menuKelolaData()
		case 2:
			prosesPembayaran()
		case 3:
			menuCariMahasiswa()
		case 0:
			fmt.Println("\n                 Terima kasih telah menggunakan sistem ini.")
			fmt.Println("============================================================================")
			return
		default:
			fmt.Println(" Pilihan tidak valid. Silakan ulangi.")
			fmt.Println("============================================================================")
		}
	}
}

// ==================== KELOLA DATA ====================
func menuKelolaData() {
	for {
		fmt.Println("============================================================================")
		fmt.Println("                         KELOLA DATA MAHASISWA")
		fmt.Println("============================================================================")
		fmt.Println(" 1. Tambah Mahasiswa")
		fmt.Println(" 2. Ubah Mahasiswa")
		fmt.Println(" 3. Hapus Mahasiswa")
		fmt.Println(" 0. Kembali")
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Print(" Pilihan Anda : ")
		var pilih int
		fmt.Scan(&pilih)
		fmt.Scanln()
		switch pilih {
		case 1:
			tambahMahasiswa()
		case 2:
			ubahMahasiswa()
		case 3:
			hapusMahasiswa()
		case 0:
			return
		default:
			fmt.Println(" Pilihan tidak valid.")
			fmt.Println("============================================================================")
		}
	}
}

func cariIndeksBerdasarkanNIM(nim string) int {
	for i := 0; i < db.JumlahData; i++ {
		if db.DataMahasiswa[i].NIM == nim {
			return i
		}
	}
	return -1
}

func tambahMahasiswa() {
	fmt.Println("============================================================================")
	fmt.Println("                           TAMBAH MAHASISWA")
	fmt.Println("----------------------------------------------------------------------------")
	if db.JumlahData == 100 {
		fmt.Println(" ERROR: Kapasitas database penuh (maksimal 100).")
		fmt.Println("============================================================================")
		return
	}
	var m Mahasiswa
	fmt.Print(" NIM   : ")
	fmt.Scanln(&m.NIM)
	fmt.Print(" Nama  : ")
	fmt.Scanln(&m.Nama)
	if m.NIM == "" || m.Nama == "" {
		fmt.Println(" ERROR: NIM dan Nama tidak boleh kosong.")
		fmt.Println("============================================================================")
		return
	}
	if cariIndeksBerdasarkanNIM(m.NIM) != -1 {
		fmt.Println(" ERROR: NIM sudah terdaftar.")
		fmt.Println("============================================================================")
		return
	}
	for i := 0; i < db.JumlahData; i++ {
		if db.DataMahasiswa[i].Nama == m.Nama {
			fmt.Println(" ERROR: Nama sudah digunakan oleh mahasiswa lain.")
			fmt.Println("============================================================================")
			return
		}
	}
	m.Tunggakan = TARGET_IURAN
	m.StatusLunas = false
	m.JumlahTransaksi = 0
	db.DataMahasiswa[db.JumlahData] = m
	db.JumlahData++
	fmt.Println(" BERHASIL: Mahasiswa berhasil ditambahkan.")
	fmt.Println("============================================================================")
}

func ubahMahasiswa() {
	fmt.Println("============================================================================")
	fmt.Println("                           UPDATE MAHASISWA")
	fmt.Println("----------------------------------------------------------------------------")
	if db.JumlahData == 0 {
		fmt.Println(" Database mahasiswa kosong.")
		fmt.Println("============================================================================")
		return
	}
	var nim string
	fmt.Print(" NIM mahasiswa yang akan diubah : ")
	fmt.Scanln(&nim)
	idx := cariIndeksBerdasarkanNIM(nim)
	if idx == -1 {
		fmt.Println(" ERROR: Mahasiswa tidak ditemukan.")
		fmt.Println("============================================================================")
		return
	}
	var nimBaru, namaBaru string
	fmt.Print(" NIM baru (kosongkan jika tidak berubah) : ")
	fmt.Scanln(&nimBaru)
	fmt.Print(" Nama baru (kosongkan jika tidak berubah) : ")
	fmt.Scanln(&namaBaru)
	if nimBaru != "" {
		cek := cariIndeksBerdasarkanNIM(nimBaru)
		if cek != -1 && cek != idx {
			fmt.Println(" ERROR: NIM baru sudah digunakan oleh mahasiswa lain.")
			fmt.Println("============================================================================")
			return
		}
		db.DataMahasiswa[idx].NIM = nimBaru
	}
	if namaBaru != "" {
		db.DataMahasiswa[idx].Nama = namaBaru
	}
	fmt.Println(" BERHASIL: Data mahasiswa berhasil diubah.")
	fmt.Println("============================================================================")
}

func hapusMahasiswa() {
	fmt.Println("============================================================================")
	fmt.Println("                           HAPUS MAHASISWA")
	fmt.Println("----------------------------------------------------------------------------")
	if db.JumlahData == 0 {
		fmt.Println(" Database mahasiswa kosong.")
		fmt.Println("============================================================================")
		return
	}
	var nim string
	fmt.Print(" NIM mahasiswa yang akan dihapus : ")
	fmt.Scanln(&nim)
	idx := cariIndeksBerdasarkanNIM(nim)
	if idx == -1 {
		fmt.Println(" ERROR: Mahasiswa tidak ditemukan.")
		fmt.Println("============================================================================")
		return
	}
	for i := idx; i < db.JumlahData-1; i++ {
		db.DataMahasiswa[i] = db.DataMahasiswa[i+1]
	}
	db.JumlahData--
	fmt.Println(" BERHASIL: Mahasiswa berhasil dihapus.")
	fmt.Println("============================================================================")
}

// ==================== PEMBAYARAN ====================
func prosesPembayaran() {
	fmt.Println("============================================================================")
	fmt.Println("                            PEMBAYARAN KAS")
	fmt.Println("----------------------------------------------------------------------------")
	if db.JumlahData == 0 {
		fmt.Println(" Database mahasiswa kosong.")
		fmt.Println("============================================================================")
		return
	}
	var nim string
	fmt.Print(" NIM mahasiswa : ")
	fmt.Scanln(&nim)
	idx := cariIndeksBerdasarkanNIM(nim)
	if idx == -1 {
		fmt.Println(" ERROR: Mahasiswa tidak ditemukan.")
		fmt.Println("============================================================================")
		return
	}
	var t Transaksi
	fmt.Print(" Tanggal (dd/mm/yyyy) : ")
	fmt.Scanln(&t.Tanggal)
	fmt.Print(" Nominal pembayaran   : ")
	fmt.Scanln(&t.Nominal)
	if t.Nominal <= 0 {
		fmt.Println(" ERROR: Nominal harus lebih dari 0.")
		fmt.Println("============================================================================")
		return
	}
	if db.DataMahasiswa[idx].JumlahTransaksi >= 100 {
		fmt.Println(" ERROR: Riwayat pembayaran penuh (maksimal 100).")
		fmt.Println("============================================================================")
		return
	}
	db.DataMahasiswa[idx].RiwayatBayar[db.DataMahasiswa[idx].JumlahTransaksi] = t
	db.DataMahasiswa[idx].JumlahTransaksi++
	totalDibayar := 0
	for i := 0; i < db.DataMahasiswa[idx].JumlahTransaksi; i++ {
		totalDibayar += db.DataMahasiswa[idx].RiwayatBayar[i].Nominal
	}
	sisa := TARGET_IURAN - totalDibayar
	if sisa <= 0 {
		db.DataMahasiswa[idx].Tunggakan = 0
		db.DataMahasiswa[idx].StatusLunas = true
		fmt.Printf(" LUNAS - Total dibayar: Rp %d (kelebihan Rp %d)\n", totalDibayar, -sisa)
	} else {
		db.DataMahasiswa[idx].Tunggakan = sisa
		db.DataMahasiswa[idx].StatusLunas = false
		fmt.Printf(" BERHASIL - Sisa tunggakan: Rp %d\n", sisa)
	}
	fmt.Println("============================================================================")
}

// ==================== PENCARIAN ====================
func menuCariMahasiswa() {
	for {
		fmt.Println("============================================================================")
		fmt.Println("                            CARI MAHASISWA")
		fmt.Println("============================================================================")
		fmt.Println(" 1. Pencarian Sekuensial (Sequential)")
		fmt.Println(" 2. Pencarian Biner (Binary) - data akan diurut")
		fmt.Println(" 0. Kembali")
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Print(" Pilihan Anda : ")
		var pilih int
		fmt.Scan(&pilih)
		fmt.Scanln()
		switch pilih {
		case 1:
			cariDenganSequential()
		case 2:
			cariDenganBinary()
		case 0:
			return
		default:
			fmt.Println(" Pilihan tidak valid.")
			fmt.Println("============================================================================")
		}
	}
}

func tampilkanDetailMahasiswa(idx int) {
	fmt.Println("============================================================================")
	fmt.Println("                            DATA MAHASISWA")
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Printf(" NIM                 : %s\n", db.DataMahasiswa[idx].NIM)
	fmt.Printf(" Nama                : %s\n", db.DataMahasiswa[idx].Nama)
	fmt.Printf(" Tunggakan           : Rp %d\n", db.DataMahasiswa[idx].Tunggakan)
	status := "Belum Lunas"
	if db.DataMahasiswa[idx].StatusLunas {
		status = "Lunas"
	}
	fmt.Printf(" Status              : %s\n", status)
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Print(" Lihat riwayat pembayaran? (y/n) : ")
	var jawab string
	fmt.Scanln(&jawab)
	if jawab == "y" || jawab == "Y" {
		if db.DataMahasiswa[idx].JumlahTransaksi == 0 {
			fmt.Println(" Belum ada riwayat pembayaran.")
		} else {
			fmt.Println("------------------------------------------------------------------------")
			fmt.Println("                         RIWAYAT PEMBAYARAN")
			fmt.Println("------------------------------------------------------------------------")
			fmt.Printf(" %-3s | %-12s | %-12s\n", "No", "Tanggal", "Nominal")
			fmt.Println("-----+--------------+-------------")
			for i := 0; i < db.DataMahasiswa[idx].JumlahTransaksi; i++ {
				fmt.Printf(" %-3d | %-12s | %-12d\n", i+1, db.DataMahasiswa[idx].RiwayatBayar[i].Tanggal, db.DataMahasiswa[idx].RiwayatBayar[i].Nominal)
			}
			fmt.Println("------------------------------------------------------------------------")
		}
	}
	fmt.Println("============================================================================")
}

func cariDenganSequential() {
	if db.JumlahData == 0 {
		fmt.Println("============================================================================")
		fmt.Println(" Database mahasiswa kosong.")
		fmt.Println("============================================================================")
		return
	}
	fmt.Println("============================================================================")
	fmt.Println(" Pencarian berdasarkan :")
	fmt.Println(" 1. NIM")
	fmt.Println(" 2. Nama")
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Print(" Pilihan Anda : ")
	var pilih int
	fmt.Scan(&pilih)
	fmt.Scanln()
	switch pilih {
	case 1:
		var nim string
		fmt.Print(" Masukkan NIM : ")
		fmt.Scanln(&nim)
		idx := cariIndeksBerdasarkanNIM(nim)
		if idx == -1 {
			fmt.Println(" Hasil: Mahasiswa tidak ditemukan.")
			fmt.Println("============================================================================")
		} else {
			tampilkanDetailMahasiswa(idx)
		}
	case 2:
		var nama string
		fmt.Print(" Masukkan Nama : ")
		fmt.Scanln(&nama)
		ditemukan := -1
		for i := 0; i < db.JumlahData; i++ {
			if db.DataMahasiswa[i].Nama == nama {
				ditemukan = i
				break
			}
		}
		if ditemukan == -1 {
			fmt.Println(" Hasil: Mahasiswa tidak ditemukan.")
			fmt.Println("============================================================================")
		} else {
			tampilkanDetailMahasiswa(ditemukan)
		}
	default:
		fmt.Println(" Pilihan tidak valid.")
		fmt.Println("============================================================================")
	}
}

// ==================== BINARY SEARCH ====================
func urutkanDataBerdasarkanNIM() {
	for i := 1; i < db.JumlahData; i++ {
		temp := db.DataMahasiswa[i]
		j := i - 1
		for j >= 0 && db.DataMahasiswa[j].NIM > temp.NIM {
			db.DataMahasiswa[j+1] = db.DataMahasiswa[j]
			j--
		}
		db.DataMahasiswa[j+1] = temp
	}
}

func urutkanDataBerdasarkanNama() {
	for i := 1; i < db.JumlahData; i++ {
		temp := db.DataMahasiswa[i]
		j := i - 1
		for j >= 0 && db.DataMahasiswa[j].Nama > temp.Nama {
			db.DataMahasiswa[j+1] = db.DataMahasiswa[j]
			j--
		}
		db.DataMahasiswa[j+1] = temp
	}
}

func binarySearchBerdasarkanNIM(nim string) int {
	kiri, kanan := 0, db.JumlahData-1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if db.DataMahasiswa[tengah].NIM == nim {
			return tengah
		} else if db.DataMahasiswa[tengah].NIM < nim {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func binarySearchBerdasarkanNama(nama string) int {
	kiri, kanan := 0, db.JumlahData-1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if db.DataMahasiswa[tengah].Nama == nama {
			return tengah
		} else if db.DataMahasiswa[tengah].Nama < nama {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func cariDenganBinary() {
	if db.JumlahData == 0 {
		fmt.Println("============================================================================")
		fmt.Println(" Database mahasiswa kosong.")
		fmt.Println("============================================================================")
		return
	}
	fmt.Println("============================================================================")
	fmt.Println(" Pencarian Biner (data akan diurutkan otomatis)")
	fmt.Println(" 1. Berdasarkan NIM")
	fmt.Println(" 2. Berdasarkan Nama")
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Print(" Pilihan Anda : ")
	var pilih int
	fmt.Scan(&pilih)
	fmt.Scanln()
	switch pilih {
	case 1:
		urutkanDataBerdasarkanNIM()
		var nim string
		fmt.Print(" Masukkan NIM : ")
		fmt.Scanln(&nim)
		idx := binarySearchBerdasarkanNIM(nim)
		if idx == -1 {
			fmt.Println(" Hasil: Mahasiswa tidak ditemukan.")
			fmt.Println("============================================================================")
		} else {
			tampilkanDetailMahasiswa(idx)
		}
	case 2:
		urutkanDataBerdasarkanNama()
		var nama string
		fmt.Print(" Masukkan Nama : ")
		fmt.Scanln(&nama)
		idx := binarySearchBerdasarkanNama(nama)
		if idx == -1 {
			fmt.Println(" Hasil: Mahasiswa tidak ditemukan.")
			fmt.Println("============================================================================")
		} else {
			tampilkanDetailMahasiswa(idx)
		}
	default:
		fmt.Println(" Pilihan tidak valid.")
		fmt.Println("============================================================================")
	}
}
