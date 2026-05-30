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
		fmt.Println(" 4. Urutkan Mahasiswa")
		fmt.Println(" 5. Daftar Mahasiswa Bertunggakan")
		fmt.Println(" 6. Statistik Kas")
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
		case 4:
			menuUrutMahasiswa()
		case 5:
			tampilkanDaftarTunggakan()
		case 6:
			tampilkanStatistikKas()
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
		fmt.Println(" 3. Tampilkan Semua")
		fmt.Println(" 4. Hapus Mahasiswa")
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
			tampilkanSemuaMahasiswa()
		case 4:
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

func tampilkanSemuaMahasiswa() {
	fmt.Println("============================================================================")
	fmt.Println("                            DAFTAR MAHASISWA")
	fmt.Println("============================================================================")
	if db.JumlahData == 0 {
		fmt.Println(" Database mahasiswa kosong.")
		fmt.Println("============================================================================")
		return
	}
	fmt.Printf(" %-3s | %-12s | %-25s | %-11s | %-9s\n", "No", "NIM", "Nama", "Tunggakan", "Status")
	fmt.Println("-----+--------------+---------------------------+-------------+-----------")
	for i := 0; i < db.JumlahData; i++ {
		status := "Belum Lunas"
		if db.DataMahasiswa[i].StatusLunas {
			status = "Lunas"
		}
		fmt.Printf(" %-3d | %-12s | %-25s | %-11d | %-9s\n",
			i+1, db.DataMahasiswa[i].NIM, db.DataMahasiswa[i].Nama, db.DataMahasiswa[i].Tunggakan, status)
	}
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

// ==================== PENGURUTAN ====================
func menuUrutMahasiswa() {
	if db.JumlahData == 0 {
		fmt.Println("============================================================================")
		fmt.Println(" Database kosong, tidak ada data yang bisa diurut.")
		fmt.Println("============================================================================")
		return
	}
	for {
		fmt.Println("============================================================================")
		fmt.Println("                         URUT DATA MAHASISWA")
		fmt.Println("============================================================================")
		fmt.Println(" 1. Selection Sort")
		fmt.Println(" 2. Insertion Sort")
		fmt.Println(" 0. Kembali")
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Print(" Pilihan Anda : ")
		var metode int
		fmt.Scan(&metode)
		fmt.Scanln()

		if metode == 0 {
			return
		}
		if metode != 1 && metode != 2 {
			fmt.Println(" Pilihan tidak valid.")
			// tidak ada continue, langsung lanjut iterasi berikutnya
		} else {
			// metode valid
			fmt.Println("============================================================================")
			fmt.Println(" Kriteria urutan :")
			fmt.Println(" 1. Nama")
			fmt.Println(" 2. NIM")
			fmt.Println(" 3. Tunggakan")
			fmt.Println(" 0. Kembali")
			fmt.Println("----------------------------------------------------------------------------")
			fmt.Print(" Pilihan Anda : ")
			var pilihKriteria int
			fmt.Scan(&pilihKriteria)
			fmt.Scanln()

			if pilihKriteria == 0 {
				return
			}
			if pilihKriteria < 1 || pilihKriteria > 3 {
				fmt.Println(" Kriteria tidak valid.")
				// skip ke iterasi berikutnya
			} else {
				// kriteria valid
				var kriteria string
				switch pilihKriteria {
				case 1:
					kriteria = "nama"
				case 2:
					kriteria = "nim"
				case 3:
					kriteria = "tunggakan"
				}
				fmt.Println("============================================================================")
				fmt.Println(" Arah urutan :")
				fmt.Println(" 1. Ascending (kecil ke besar)")
				fmt.Println(" 2. Descending (besar ke kecil)")
				fmt.Println("----------------------------------------------------------------------------")
				fmt.Print(" Pilihan Anda : ")
				var pilihArah int
				fmt.Scan(&pilihArah)
				fmt.Scanln()

				if pilihArah != 1 && pilihArah != 2 {
					fmt.Println(" Arah tidak valid.")
					// skip
				} else {
					ascending := pilihArah == 1
					if metode == 1 {
						selectionSort(kriteria, ascending)
						fmt.Println(" BERHASIL: Data diurutkan dengan Selection Sort.")
					} else {
						insertionSort(kriteria, ascending)
						fmt.Println(" BERHASIL: Data diurutkan dengan Insertion Sort.")
					}
					fmt.Println("============================================================================")
					tampilkanSemuaMahasiswa()
				}
			}
		}
	}
}

func bandingkan(i, j int, kriteria string, ascending bool) bool {
	if kriteria == "nama" {
		if ascending {
			return db.DataMahasiswa[i].Nama < db.DataMahasiswa[j].Nama
		}
		return db.DataMahasiswa[i].Nama > db.DataMahasiswa[j].Nama
	} else if kriteria == "nim" {
		if ascending {
			return db.DataMahasiswa[i].NIM < db.DataMahasiswa[j].NIM
		}
		return db.DataMahasiswa[i].NIM > db.DataMahasiswa[j].NIM
	} else {
		if ascending {
			return db.DataMahasiswa[i].Tunggakan < db.DataMahasiswa[j].Tunggakan
		}
		return db.DataMahasiswa[i].Tunggakan > db.DataMahasiswa[j].Tunggakan
	}
}

func selectionSort(kriteria string, ascending bool) {
	for i := 0; i < db.JumlahData-1; i++ {
		idxMin := i
		for j := i + 1; j < db.JumlahData; j++ {
			if bandingkan(j, idxMin, kriteria, ascending) {
				idxMin = j
			}
		}
		if idxMin != i {
			db.DataMahasiswa[i], db.DataMahasiswa[idxMin] = db.DataMahasiswa[idxMin], db.DataMahasiswa[i]
		}
	}
}

func insertionSort(kriteria string, ascending bool) {
	for i := 1; i < db.JumlahData; i++ {
		temp := db.DataMahasiswa[i]
		j := i - 1
		for j >= 0 && bandingkan(j+1, j, kriteria, ascending) {
			db.DataMahasiswa[j+1] = db.DataMahasiswa[j]
			j--
		}
		db.DataMahasiswa[j+1] = temp
	}
}

// ==================== DAFTAR TUNGGAKAN ====================
func tampilkanDaftarTunggakan() {
	if db.JumlahData == 0 {
		fmt.Println("============================================================================")
		fmt.Println(" Belum ada data mahasiswa.")
		fmt.Println("============================================================================")
		return
	}
	var bertunggakan [100]Mahasiswa
	jumlahTunggak := 0
	for i := 0; i < db.JumlahData; i++ {
		if db.DataMahasiswa[i].Tunggakan > 0 {
			bertunggakan[jumlahTunggak] = db.DataMahasiswa[i]
			jumlahTunggak++
		}
	}
	if jumlahTunggak == 0 {
		fmt.Println("============================================================================")
		fmt.Println(" Seluruh mahasiswa telah LUNAS.")
		fmt.Println("============================================================================")
		return
	}
	fmt.Println("============================================================================")
	fmt.Println("                   DAFTAR MAHASISWA BERTUNGGAKAN")
	fmt.Println("============================================================================")
	fmt.Println(" Pilih urutan :")
	fmt.Println(" 1. Ascending (tunggakan kecil ke besar)")
	fmt.Println(" 2. Descending (tunggakan besar ke kecil)")
	fmt.Println(" 0. Kembali")
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Print(" Pilihan Anda : ")
	var pilih int
	fmt.Scan(&pilih)
	fmt.Scanln()
	if pilih == 0 {
		return
	}
	if pilih != 1 && pilih != 2 {
		fmt.Println(" Pilihan tidak valid.")
		fmt.Println("============================================================================")
		return
	}
	// bubble sort
	for i := 0; i < jumlahTunggak-1; i++ {
		for j := i + 1; j < jumlahTunggak; j++ {
			if pilih == 1 {
				if bertunggakan[i].Tunggakan > bertunggakan[j].Tunggakan {
					bertunggakan[i], bertunggakan[j] = bertunggakan[j], bertunggakan[i]
				}
			} else {
				if bertunggakan[i].Tunggakan < bertunggakan[j].Tunggakan {
					bertunggakan[i], bertunggakan[j] = bertunggakan[j], bertunggakan[i]
				}
			}
		}
	}
	fmt.Printf(" %-3s | %-12s | %-25s | %-11s\n", "No", "NIM", "Nama", "Tunggakan")
	fmt.Println("-----+--------------+---------------------------+------------")
	for i := 0; i < jumlahTunggak; i++ {
		fmt.Printf(" %-3d | %-12s | %-25s | %-11d\n", i+1, bertunggakan[i].NIM, bertunggakan[i].Nama, bertunggakan[i].Tunggakan)
	}
	fmt.Println("============================================================================")
}

// ==================== STATISTIK KAS ====================
func tampilkanStatistikKas() {
	fmt.Println("============================================================================")
	fmt.Println("                           STATISTIK KAS")
	fmt.Println("============================================================================")
	if db.JumlahData == 0 {
		fmt.Println(" Belum ada data mahasiswa.")
		fmt.Println("============================================================================")
		return
	}
	totalPembayaran := 0
	totalTunggakan := 0
	jumlahLunas := 0
	jumlahBelumLunas := 0
	for i := 0; i < db.JumlahData; i++ {
		dibayar := 0
		for j := 0; j < db.DataMahasiswa[i].JumlahTransaksi; j++ {
			dibayar += db.DataMahasiswa[i].RiwayatBayar[j].Nominal
		}
		totalPembayaran += dibayar
		if db.DataMahasiswa[i].StatusLunas {
			jumlahLunas++
		} else {
			jumlahBelumLunas++
			totalTunggakan += db.DataMahasiswa[i].Tunggakan
		}
	}
	targetTotal := TARGET_IURAN * db.JumlahData
	persentase := float64(totalPembayaran) / float64(targetTotal) * 100

	fmt.Println(" RINGKASAN KAS MAHASISWA")
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Printf(" Jumlah Mahasiswa         : %d\n", db.JumlahData)
	fmt.Printf(" Target Iuran per Orang   : Rp %d\n", TARGET_IURAN)
	fmt.Printf(" Target Total Kas         : Rp %d\n", targetTotal)
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Printf(" Mahasiswa LUNAS          : %d\n", jumlahLunas)
	fmt.Printf(" Mahasiswa BELUM LUNAS    : %d\n", jumlahBelumLunas)
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Printf(" Total Pembayaran Masuk   : Rp %d\n", totalPembayaran)
	fmt.Printf(" Total Tunggakan          : Rp %d\n", totalTunggakan)
	fmt.Printf(" Persentase Kelunasan     : %.2f%%\n", persentase)
	if jumlahBelumLunas > 0 {
		rataRata := totalTunggakan / jumlahBelumLunas
		fmt.Printf(" Rata-rata Tunggakan (BL) : Rp %d\n", rataRata)
	}
	fmt.Println("============================================================================")
}
