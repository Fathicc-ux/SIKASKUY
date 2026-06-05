package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type mahasiswa struct {
	NIM               string        `json:"nim"`
	Nama              string        `json:"nama"`
	Tunggakan         int           `json:"tunggakan"`
	StatusPembayaraan bool          `json:"status_pembayaraan"`
	Riwayat           [12]transaksi `json:"riwayat"`
	JumlahBayar       int           `json:"jumlah_bayar"`
}

type Database struct {
	Mahasiswa      [100]mahasiswa `json:"mahasiswa"`
	MahasiswaCount int            `json:"mahasiswa_count"`
}

type transaksi struct {
	Tanggal string `json:"tanggal"`
	Nominal int    `json:"nominal"`
}

var DB Database

const dataFile = "sikas_data.json"

func saveData() error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // biar rapi saat dilihat di editor
	err = encoder.Encode(DB)
	return err
}

func loadData() error {
	file, err := os.Open(dataFile)
	if err != nil {
		// File tidak ada, berarti program baru dijalankan pertama kali
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&DB)
	return err
}

// A. Fungsi Tambah
func tambah(A *Database) {
	fmt.Println()
	fmt.Println("=== TAMBAH DATA MAHASISWA ===")
	fmt.Println("Ketik 'x' atau 'keluar' kapan saja untuk membatalkan")
	fmt.Println()

	if A.MahasiswaCount >= 100 {
		fmt.Println("== Database Mahasiswa Penuh ==")
		fmt.Println()
		return
	}

	for {
		var mhs mahasiswa

		fmt.Print("Nama: ")
		fmt.Scan(&mhs.Nama)
		if mhs.Nama == "x" || mhs.Nama == "keluar" {
			fmt.Println("== Penambahan data dibatalkan ==")
			fmt.Println()
			return
		}

		fmt.Print("NIM: ")
		fmt.Scan(&mhs.NIM)
		if mhs.NIM == "x" || mhs.NIM == "keluar" {
			fmt.Println("== Penambahan data dibatalkan ==")
			fmt.Println()
			return
		}

		var ceknim int
		squential(A, mhs.NIM, &ceknim)
		if ceknim != -1 {
			fmt.Println("== NIM sudah digunakan ==")
			fmt.Println()
		} else {
			mhs.Tunggakan = 5000
			mhs.StatusPembayaraan = false
			A.Mahasiswa[A.MahasiswaCount] = mhs
			A.MahasiswaCount++
			fmt.Println("== Mahasiswa Berhasil Ditambahkan ==")
			fmt.Println()
		}

		// Simpan ke file setelah perubahan
		if err := saveData(); err != nil {
			fmt.Println("Peringatan: Gagal menyimpan data:", err)
		}
	}
}

// B. Fungsi Ubah (update)
func ubah(A *Database) {
	fmt.Println()
	fmt.Println("=== UPDATE DATA MAHASISWA ===")
	fmt.Println("Ketik 'x' atau 'keluar' kapan saja untuk membatalkan")
	fmt.Println()

	if A.MahasiswaCount == 0 {
		fmt.Println("== Database Mahasiswa Kosong ==")
		fmt.Println()
		return
	}

	for {
		var nim string
		var indeks int

		fmt.Print("Masukkan NIM Mahasiswa yang akan diupdate: ")
		fmt.Scan(&nim)
		squential(A, nim, &indeks)

		if nim == "x" || nim == "keluar" {
			fmt.Println("== Update data dibatalkan ==")
			fmt.Println()
			return
		} else if indeks == -1 {
			fmt.Println("== Mahasiswa Tidak Ditemukan ==")
			fmt.Println()
			return
		}

		fmt.Print("Nama: ")
		fmt.Scan(&A.Mahasiswa[indeks].Nama)

		fmt.Print("NIM: ")
		fmt.Scan(&A.Mahasiswa[indeks].NIM)

		var ceknim int
		squential(A, A.Mahasiswa[indeks].NIM, &ceknim)
		if ceknim != -1 && ceknim != indeks {
			fmt.Println("== NIM sudah digunakan ==")
			fmt.Println()
			return
		}

		fmt.Println("== Mahasiswa Berhasil Diupdate ==")
		fmt.Println()

		// Simpan ke file setelah perubahan
		if err := saveData(); err != nil {
			fmt.Println("Peringatan: Gagal menyimpan data:", err)
		}
		return
	}
}

// C. Fungsi Hapus
func hapus(A *Database) {
	fmt.Println()
	fmt.Println("=== HAPUS DATA MAHASISWA ===")
	fmt.Println("Ketik 'x' atau 'keluar' kapan saja untuk membatalkan")
	fmt.Println()

	if A.MahasiswaCount == 0 {
		fmt.Println("== Database Mahasiswa Kosong ==")
		fmt.Println()
		return
	}

	for {
		var nim string
		var indeks int

		fmt.Print("Masukkan NIM Mahasiswa yang akan dihapus: ")
		fmt.Scan(&nim)
		if nim == "x" || nim == "keluar" {
			fmt.Println("== Hapus data dibatalkan ==")
			fmt.Println()
			return
		}

		squential(A, nim, &indeks)
		if indeks == -1 {
			fmt.Println("== Mahasiswa Tidak Ditemukan ==")
			fmt.Println()
		} else {
			for i := indeks; i < A.MahasiswaCount-1; i++ {
				A.Mahasiswa[i] = A.Mahasiswa[i+1]
			}
			A.MahasiswaCount--
			A.Mahasiswa[A.MahasiswaCount] = mahasiswa{}
			fmt.Println("== Mahasiswa Berhasil Dihapus ==")
			fmt.Println()

			if err := saveData(); err != nil {
				fmt.Println("Peringatan: Gagal menyimpan data:", err)
			}
		}
	}
}

// B. Fungsi Bayar ( mencatat nominal iuran yang masuk serta tanggal pembayaran)
func bayar(A *Database) {
	fmt.Println()
	fmt.Println("=== PEMBAYARAN KAS MAHASISWA ===")
	fmt.Println("Ketik 'x' atau 'keluar' kapan saja untuk membatalkan")
	fmt.Println()

	if A.MahasiswaCount == 0 {
		fmt.Println("== Database Mahasiswa Kosong ==")
		fmt.Println()
		return
	}

	for {
		var nim string
		var indeks int
		var bayar transaksi

		fmt.Print("Masukkan NIM Mahasiswa yang akan melakukan pembayaran: ")
		fmt.Scan(&nim)
		squential(A, nim, &indeks)

		if indeks == -1 {
			fmt.Println("== Mahasiswa Tidak Ditemukan ==")
			return
		} else if nim == "x" || nim == "keluar" {
			fmt.Println("== Pembayaran kas dibatalkan ==")
			fmt.Println()
			return
		}

		//tanggal otomatis
		bayar.Tanggal = time.Now().Format("02/01/2006")
		fmt.Print("Masukan Nominal Pembayaran: ")
		fmt.Scan(&bayar.Nominal)

		//simpan riwayat pembayaran
		if A.Mahasiswa[indeks].JumlahBayar < 12 {
			A.Mahasiswa[indeks].Riwayat[A.Mahasiswa[indeks].JumlahBayar] = bayar
			A.Mahasiswa[indeks].JumlahBayar++
		} else {
			fmt.Println("== Riwayat pembayaran sudah penuh ==")
			return
		}

		// Kurangi tunggakan
		A.Mahasiswa[indeks].Tunggakan -= bayar.Nominal
		if A.Mahasiswa[indeks].Tunggakan <= 0 {
			A.Mahasiswa[indeks].Tunggakan = 0
			A.Mahasiswa[indeks].StatusPembayaraan = true
			fmt.Println("== Pembayaran Berhasil Dicatat ==")
			fmt.Println()
		} else {
			A.Mahasiswa[indeks].StatusPembayaraan = false
			fmt.Println("== Pembayaran Berhasil Dicatat ==")
			fmt.Println("Status pembayaran: Belum lunas")
			fmt.Println("Sisa Tunggakan: ", A.Mahasiswa[indeks].Tunggakan)
			fmt.Println()
		}

		if err := saveData(); err != nil {
			fmt.Println("Peringatan: Gagal menyimpan data:", err)
		}
	}
}

// c. Fungsi (Sequential)
func squential(A *Database, nim string, indeks *int) {
	var i int
	*indeks = -1

	for i = 0; i < A.MahasiswaCount; i++ {
		if A.Mahasiswa[i].NIM == nim {
			*indeks = i
		}
	}
}

// C. Fungsi cari squential sub fungsi squential untuk mencari mahasiswa yang belum bayar berdasarkan NIM dan Nama yang diinputkan
func cari_squential(A *Database) {
	fmt.Println()
	fmt.Println("=== CARI SEQUENTIAL ===")
	fmt.Println("Ketik 'x' atau 'keluar' kapan saja untuk membatalkan")
	fmt.Println()

	if A.MahasiswaCount == 0 {
		fmt.Println("== Database Mahasiswa Kosong ==")
		fmt.Println()
		return
	}

	for {
		var nim string
		var indeks int

		fmt.Println("====")
		fmt.Print("Masukkan NIM Mahasiswa yang akan dicari: ")
		fmt.Scan(&nim)

		if nim == "x" || nim == "keluar" {
			fmt.Println("== Cari data dibatalkan ==")
			fmt.Println()
			return
		}

		squential(A, nim, &indeks)
		if indeks == -1 {
			fmt.Println("== Mahasiswa Tidak Ditemukan ==")
		} else {
			if A.Mahasiswa[indeks].StatusPembayaraan == false {
				fmt.Println("== Daftar Mahasiswa Belum Lunas ==")
				fmt.Println("NIM  :", A.Mahasiswa[indeks].NIM)
				fmt.Println("Nama :", A.Mahasiswa[indeks].Nama)
			} else {
				fmt.Println("Mahasiswa Sudah Membayar secara Lunas")
				fmt.Println("NIM  :", A.Mahasiswa[indeks].NIM)
				fmt.Println("Nama :", A.Mahasiswa[indeks].Nama)
			}
			fmt.Println()
		}
	}
}

// c. Fungsi (Binary)
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

// C. Fungsi Cari Sequential sub fungsi squential untuk mencari mahasiswa yang belum bayar berdasarkan NIM dan Nama yang diinputkan
func cari_binary(A *Database) {
	fmt.Println()
	fmt.Println("=== CARI SEQUENTIAL ===")
	fmt.Println("Ketik 'x' atau 'keluar' kapan saja untuk membatalkan")
	fmt.Println()

	if A.MahasiswaCount == 0 {
		fmt.Println("== Database Mahasiswa Kosong ==")
		fmt.Println()
		return
	}

	for {
		var nim string
		var indeks int

		fmt.Print("Masukkan NIM Mahasiswa yang akan dicari: ")
		fmt.Scan(&nim)
		insertionsortNIM(A)
		binary(A, nim, &indeks)

		if nim == "x" || nim == "keluar" {
			fmt.Println("== Pembayaran kas dibatalkan ==")
			fmt.Println()
			return
		} else if indeks == -1 {
			fmt.Println("== Mahasiswa Tidak Ditemukan ==")
		} else {
			fmt.Println("NIM  :", A.Mahasiswa[indeks].NIM)
			fmt.Println("Nama :", A.Mahasiswa[indeks].Nama)
			fmt.Println()
		}
	}
}

// D. Fungsi Insertionsort berdasarkan NIM untuk binary search
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

// D. Fungsi Selectionsort Asscending berdasakan nama dan tunggakan
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
				if A.Mahasiswa[j].Tunggakan < A.Mahasiswa[indeks].Tunggakan {
					indeks = j
				}
			}
		}
		sementara = A.Mahasiswa[i]
		A.Mahasiswa[i] = A.Mahasiswa[indeks]
		A.Mahasiswa[indeks] = sementara
	}
}

// D. Fungsi Selectionsort Descending berdasarkan nama dan tunggakan
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
				if A.Mahasiswa[j].Tunggakan > A.Mahasiswa[indeks].Tunggakan {
					indeks = j
				}
			}
		}
		sementara = A.Mahasiswa[i]
		A.Mahasiswa[i] = A.Mahasiswa[indeks]
		A.Mahasiswa[indeks] = sementara
	}
}

// D. Fungsi insetionsort Ascending berdasarkan nama dan tunggakan
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
			for j >= 0 && A.Mahasiswa[j].Tunggakan > sementara.Tunggakan {
				A.Mahasiswa[j+1] = A.Mahasiswa[j]
				j--
			}
		}
		A.Mahasiswa[j+1] = sementara
	}
}

// D. Fungsi insertionsort Descending berdasarkan nama dan tunggakan
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
			for j >= 0 && A.Mahasiswa[j].Tunggakan < sementara.Tunggakan {
				A.Mahasiswa[j+1] = A.Mahasiswa[j]
				j--
			}
		}
		A.Mahasiswa[j+1] = sementara
	}
}

// Fungsi tampil untuk menampilkan daftar mahasiswa beserta status pembayaran, tunggakan, dan riwayat pembayaran jika ada
func tampil(A *Database) {
	var i, k int

	if A.MahasiswaCount == 0 {
		fmt.Println("== Database Mahasiswa Kosong ==")
		return
	}

	fmt.Println("=== Daftar Status Pembayaran Mahasiswa ===")
	fmt.Println()

	for i = 0; i < A.MahasiswaCount; i++ {
		fmt.Println("Nama :", A.Mahasiswa[i].Nama)
		fmt.Println("NIM  :", A.Mahasiswa[i].NIM)
		fmt.Println("Tunggakan :", A.Mahasiswa[i].Tunggakan)

		if A.Mahasiswa[i].StatusPembayaraan {
			fmt.Println("Status Pembayaran : Lunas")
		} else {
			fmt.Println("Status Pembayaran : Belum Lunas")
		}

		fmt.Println("Riwayat Pembayaran:")

		if A.Mahasiswa[i].JumlahBayar == 0 {
			fmt.Println("Belum ada pembayaran")
		} else {
			for k = 0; k < A.Mahasiswa[i].JumlahBayar; k++ {
				tr := A.Mahasiswa[i].Riwayat[k]
				fmt.Printf("  %d. Tanggal: %s | Nominal: %d\n",
					k+1, tr.Tanggal, tr.Nominal)
			}
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

		if A.Mahasiswa[i].StatusPembayaraan == true {
			jumlahLunas++
		}

		for j = 0; j < A.Mahasiswa[i].JumlahBayar; j++ {
			totalKas += A.Mahasiswa[i].Riwayat[j].Nominal
		}
	}
	fmt.Println("STATISTIK KAS")
	fmt.Println("==============================")
	fmt.Println("Total Saldo Kas       :", totalKas)
	fmt.Println("Mahasiswa Lunas       :", jumlahLunas)
	fmt.Println("Total Mahasiswa       :", A.MahasiswaCount)
	fmt.Println("==============================")
}

func pilihselection(A *Database) {
	if A.MahasiswaCount == 0 {
		fmt.Println("== Database Mahasiswa Kosong ==")
		fmt.Println()
		return
	}

	for {
		fmt.Println("=== Pilih pengurutan ===")
		//var kategori int
		fmt.Println("1. Nama ascending")
		fmt.Println("2. Nama discending")
		fmt.Println("3. Tunggakan ascending")
		fmt.Println("4. Tunggakan discending")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		var kategori int
		fmt.Scan(&kategori)

		if kategori == 1 {
			selectionsortAscending(&DB, "Nama")
		} else if kategori == 2 {
			selectionsortDescending(&DB, "Nama")
		} else if kategori == 3 {
			selectionsortAscending(&DB, "Tunggakan")
		} else if kategori == 4 {
			selectionsortDescending(&DB, "Tunggakan")
		} else if kategori == 0 {
			return
		}
		tampil(&DB)
	}
}

func pilihinsertion(A *Database) {
	if A.MahasiswaCount == 0 {
		fmt.Println("== Database Mahasiswa Kosong ==")
		fmt.Println()
		return
	}

	for {
		fmt.Println("=== Pilih pengurutan ===")
		//var kategori int
		fmt.Println("1. Nama ascending")
		fmt.Println("2. Nama discending")
		fmt.Println("3. Tunggakan ascending")
		fmt.Println("4. Tunggakan discending")
		fmt.Println("0. kembali")
		fmt.Print("Pilih: ")

		var kategori int
		fmt.Scan(&kategori)

		if kategori == 1 {
			insertionsortAscending(&DB, "Nama")
		} else if kategori == 2 {
			insertionsortDescending(&DB, "Nama")
		} else if kategori == 3 {
			insertionsortAscending(&DB, "Tunggakan")
		} else if kategori == 4 {
			insertionsortDescending(&DB, "Tunggakan")
		} else if kategori == 0 {
			return
		}
		tampil(&DB)
	}
}

func main() {
	var pilih int

	for {
		fmt.Println()
		fmt.Println("===== SIKAS =====")
		fmt.Println("1.  Tambah")
		fmt.Println("2.  Tampil")
		fmt.Println("3.  Ubah")
		fmt.Println("4.  Hapus")
		fmt.Println("5.  Bayar")
		fmt.Println("6.  Cari Sequential")
		fmt.Println("7.  Cari Binary")
		fmt.Println("8.  Selectionsort")
		fmt.Println("9.  Insertionsort")
		fmt.Println("10. Statistik")
		fmt.Println("0.  Keluar")

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

			pilihselection(&DB)

		} else if pilih == 9 {

			pilihinsertion(&DB)

		} else if pilih == 10 {
			statistik(&DB)
		} else if pilih == 0 {

			fmt.Println("Selesai")
			return

		} else {
			fmt.Println("Menu salah")
		}
	}
}
