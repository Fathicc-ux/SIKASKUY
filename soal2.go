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
	data         [50]Mahasiswa
	jmlMahasiswa int
}

var DB dataMahasiswa

func main() {
	for {
		fmt.Println("\n========================================")
		fmt.Println("           SISTEM KAS MAHASISWA")
		fmt.Println("========================================")
		fmt.Println(" 1. Kelola Data Mahasiswa")
		fmt.Println(" 2. Pembayaran Kas")
		fmt.Println(" 3. Cari Mahasiswa")
		fmt.Println(" 4. Daftar Tunggakan Kas")
		fmt.Println(" 5. Urutkan Data")
		fmt.Println(" 6. Statistik Kas")
		fmt.Println(" 7. Keluar")
		fmt.Println("----------------------------------------")
		fmt.Print(" Pilih menu (1-7): ")

		var menu int
		fmt.Scan(&menu)
		fmt.Scanln()

		switch menu {
		case 1:
			menusiswa(&DB)
		case 2:
			bayar(&DB)
		case 3:
			binarymahasiswa(&DB)
		case 4:
			daftartunggakan(&DB)
		case 5:
			urutdata(&DB)
		case 6:
			statistikKas(&DB)
		case 7:
			fmt.Println("\n Terima kasih telah menggunakan program ini.")
			return
		default:
			fmt.Println(" ⚠️  Pilihan tidak valid! Silakan coba lagi.")
		}
	}
}

func menusiswa(A *dataMahasiswa) {
	for {
		fmt.Println("\n========================================")
		fmt.Println("        KELOLA DATA MAHASISWA")
		fmt.Println("========================================")
		fmt.Println(" 1. Tambah Data Mahasiswa")
		fmt.Println(" 2. Ubah Data Mahasiswa")
		fmt.Println(" 3. Tampilkan Data Mahasiswa")
		fmt.Println(" 4. Hapus Data Mahasiswa")
		fmt.Println(" 5. Kembali")
		fmt.Println("----------------------------------------")
		fmt.Print(" Pilih menu (1-5): ")

		var menu int
		fmt.Scan(&menu)
		fmt.Scanln()

		switch menu {
		case 1:
			tambah(A)
		case 2:
			ubah(A)
		case 3:
			tampilkan(A)
		case 4:
			hapus(A)
		case 5:
			return
		default:
			fmt.Println(" ⚠️  Pilihan tidak valid!")
		}
	}
}

// sequential search
func sequential(A *dataMahasiswa, Nim string, indeks *int) {
	*indeks = -1
	for i := 0; i < A.jmlMahasiswa; i++ {
		if A.data[i].nim == Nim {
			*indeks = i
		}
	}
}

// TAMBAH DATA
func tambah(A *dataMahasiswa) {
	fmt.Println("\n========================================")
	fmt.Println("         TAMBAH DATA MAHASISWA")
	fmt.Println("========================================")
	if A.jmlMahasiswa < 50 {
		var mhs Mahasiswa

		fmt.Print(" Masukkan NIM          : ")
		fmt.Scanln(&mhs.nim)
		fmt.Print(" Masukkan Nama Lengkap : ")
		fmt.Scanln(&mhs.nama)

		if mhs.nim == "" || mhs.nama == "" {
			fmt.Println(" ⚠️  GAGAL: NIM dan Nama tidak boleh kosong!")
			return
		}

		var cek int
		sequential(A, mhs.nim, &cek)
		if cek != -1 {
			fmt.Println(" ⚠️  NIM sudah digunakan!")
			return
		}

		for i := 0; i < A.jmlMahasiswa; i++ {
			if A.data[i].nama == mhs.nama {
				fmt.Println(" ⚠️  Nama sudah digunakan!")
				return
			}
		}

		mhs.tunggakan = 5000
		mhs.statusbayar = false
		mhs.jumlahbayar = 0
		mhs.jmlriwayat = 0
		A.data[A.jmlMahasiswa] = mhs
		A.jmlMahasiswa++
		fmt.Println(" ✅ Data berhasil ditambahkan!")
	} else {
		fmt.Println(" ⚠️  Data sudah penuh (maksimal 50 mahasiswa).")
	}
}

// UBAH DATA
func ubah(A *dataMahasiswa) {
	fmt.Println("\n========================================")
	fmt.Println("         UPDATE DATA MAHASISWA")
	fmt.Println("========================================")

	var Nim string
	var indeks int

	fmt.Print(" Masukkan NIM mahasiswa : ")
	fmt.Scanln(&Nim)

	if Nim == "" {
		fmt.Println(" ⚠️  NIM tidak boleh kosong!")
		return
	}

	sequential(A, Nim, &indeks)
	if indeks == -1 {
		fmt.Println(" ⚠️  Data mahasiswa tidak ditemukan!")
		return
	}

	var Nimbaru, Namabaru string
	fmt.Print(" Masukkan NIM baru (kosongkan jika tidak berubah) : ")
	fmt.Scanln(&Nimbaru)
	fmt.Print(" Masukkan Nama baru (kosongkan jika tidak berubah): ")
	fmt.Scanln(&Namabaru)

	if Nimbaru != "" {
		var cek int
		sequential(A, Nimbaru, &cek)
		if cek != -1 && cek != indeks {
			fmt.Println(" ⚠️  NIM baru sudah digunakan oleh mahasiswa lain!")
			return
		}
		A.data[indeks].nim = Nimbaru
	}

	if Namabaru != "" {
		A.data[indeks].nama = Namabaru
	}

	if Nimbaru == "" && Namabaru == "" {
		fmt.Println(" ℹ️  Tidak ada perubahan data.")
	} else {
		fmt.Println(" ✅ Data berhasil diubah!")
	}
}

// TAMPILKAN DATA (rapi)
func tampilkan(A *dataMahasiswa) {
	fmt.Println("\n========================================")
	fmt.Println("         DAFTAR MAHASISWA")
	fmt.Println("========================================")
	if A.jmlMahasiswa == 0 {
		fmt.Println(" ⚠️  Belum ada data mahasiswa.")
	} else {
		// Header
		fmt.Printf(" %-3s | %-10s | %-25s | %-12s | %-10s\n", "No", "NIM", "Nama", "Total Bayar", "Status")
		fmt.Println("-----+------------+---------------------------+--------------+------------")
		for i := 0; i < A.jmlMahasiswa; i++ {
			status := "Belum Lunas"
			if A.data[i].statusbayar {
				status = "Lunas"
			}
			fmt.Printf(" %-3d | %-10s | %-25s | Rp %-9d | %-10s\n",
				i+1, A.data[i].nim, A.data[i].nama, A.data[i].jumlahbayar, status)
		}
	}
	fmt.Println()
}

// HAPUS DATA
func hapus(A *dataMahasiswa) {
	fmt.Println("\n========================================")
	fmt.Println("         HAPUS DATA MAHASISWA")
	fmt.Println("========================================")
	var Nim string
	var indeks int

	fmt.Print(" Masukkan NIM mahasiswa : ")
	fmt.Scanln(&Nim)

	if Nim == "" {
		fmt.Println(" ⚠️  NIM tidak boleh kosong!")
		return
	}

	sequential(A, Nim, &indeks)
	if indeks == -1 {
		fmt.Println(" ⚠️  Data mahasiswa tidak ditemukan!")
		return
	}

	for i := indeks; i < A.jmlMahasiswa-1; i++ {
		A.data[i] = A.data[i+1]
	}
	A.jmlMahasiswa--
	fmt.Println(" ✅ Data mahasiswa berhasil dihapus!")
}

// PEMBAYARAN KAS (rapi)
func bayar(A *dataMahasiswa) {
	fmt.Println("\n========================================")
	fmt.Println("           PEMBAYARAN KAS")
	fmt.Println("========================================")

	var nim string
	var indeks int

	fmt.Print(" Masukkan NIM : ")
	fmt.Scan(&nim)
	fmt.Scanln()

	sequential(A, nim, &indeks)
	if indeks == -1 {
		fmt.Println(" ⚠️  Data mahasiswa tidak ditemukan!")
		return
	}

	fmt.Printf(" Nama         : %s\n", A.data[indeks].nama)
	fmt.Printf(" Tunggakan    : Rp %d\n", A.data[indeks].tunggakan)

	if A.data[indeks].tunggakan == 0 {
		fmt.Println(" ℹ️  Mahasiswa sudah lunas, tidak perlu membayar.")
		return
	}

	var nominal int
	fmt.Print(" Masukkan nominal pembayaran : Rp ")
	fmt.Scan(&nominal)
	fmt.Scanln()

	if nominal <= 0 {
		fmt.Println(" ⚠️  Nominal harus positif!")
		return
	}
	if nominal > A.data[indeks].tunggakan {
		fmt.Printf(" ⚠️  Nominal melebihi tunggakan (Rp %d)!\n", A.data[indeks].tunggakan)
		return
	}

	A.data[indeks].tunggakan -= nominal
	A.data[indeks].jumlahbayar += nominal

	if A.data[indeks].jmlriwayat < 100 {
		A.data[indeks].riwayat[A.data[indeks].jmlriwayat] = Pembayaran{
			nominal: nominal,
			waktu:   "tercatat",
		}
		A.data[indeks].jmlriwayat++
	}

	if A.data[indeks].tunggakan == 0 {
		A.data[indeks].statusbayar = true
		fmt.Println(" ✅ PEMBAYARAN LUNAS! Status menjadi LUNAS.")
	} else {
		fmt.Printf(" ✅ Pembayaran berhasil. Sisa tunggakan : Rp %d\n", A.data[indeks].tunggakan)
	}
}

// BINARY SEARCH MAHASISWA (rapi)
func binary(A *dataMahasiswa, nim string) int {
	left := 0
	right := A.jmlMahasiswa - 1
	for left <= right {
		mid := (left + right) / 2
		if A.data[mid].nim == nim {
			return mid
		} else if A.data[mid].nim < nim {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func binarynim(A *dataMahasiswa) {
	if A.jmlMahasiswa > 1 {
		selection(A, "NIM", true)
	}
}

func binarymahasiswa(A *dataMahasiswa) {
	fmt.Println("\n========================================")
	fmt.Println("      CARI MAHASISWA (BINARY SEARCH)")
	fmt.Println("========================================")
	if A.jmlMahasiswa == 0 {
		fmt.Println(" ⚠️  Belum ada data mahasiswa.")
		return
	}

	binarynim(A)

	var Nim string
	fmt.Print(" Masukkan NIM Mahasiswa: ")
	fmt.Scanln(&Nim)

	indeks := binary(A, Nim)
	if indeks == -1 {
		fmt.Println(" ⚠️  Mahasiswa dengan NIM tersebut tidak ditemukan.")
	} else {
		fmt.Println(" ✅ Data ditemukan:")
		status := "Belum Lunas"
		if A.data[indeks].statusbayar {
			status = "Lunas"
		}
		fmt.Printf(" NIM         : %s\n", A.data[indeks].nim)
		fmt.Printf(" Nama        : %s\n", A.data[indeks].nama)
		fmt.Printf(" Total Bayar : Rp %d\n", A.data[indeks].jumlahbayar)
		fmt.Printf(" Tunggakan   : Rp %d\n", A.data[indeks].tunggakan)
		fmt.Printf(" Status      : %s\n", status)
	}
	fmt.Println()
}

// DAFTAR TUNGGAKAN (rapi)
func daftartunggakan(A *dataMahasiswa) {
	fmt.Println("\n========================================")
	fmt.Println("       DAFTAR MAHASISWA BERTUNGGAKAN")
	fmt.Println("========================================")
	if A.jmlMahasiswa == 0 {
		fmt.Println(" ⚠️  Belum ada data mahasiswa.")
		return
	}

	ada := false
	fmt.Printf(" %-3s | %-10s | %-25s | %-12s\n", "No", "NIM", "Nama", "Tunggakan")
	fmt.Println("----+------------+---------------------------+-------------")
	for i := 0; i < A.jmlMahasiswa; i++ {
		if A.data[i].tunggakan > 0 {
			ada = true
			fmt.Printf(" %-3d | %-10s | %-25s | Rp %-9d\n",
				i+1, A.data[i].nim, A.data[i].nama, A.data[i].tunggakan)
		}
	}
	if !ada {
		fmt.Println(" ✅ Semua mahasiswa sudah LUNAS!")
	}
	fmt.Println()
}

// FUNGSI PEMBANDING (terkecil) – tidak diubah
func terkecil(a, b Mahasiswa, kategori string, asc bool) bool {
	switch kategori {
	case "NIM":
		if asc {
			return a.nim < b.nim
		} else {
			return a.nim > b.nim
		}
	case "Nama":
		if asc {
			return a.nama < b.nama
		} else {
			return a.nama > b.nama
		}
	case "Tunggakan":
		if asc {
			return a.tunggakan < b.tunggakan
		} else {
			return a.tunggakan > b.tunggakan
		}
	case "Jumlah Bayar":
		if asc {
			return a.jumlahbayar < b.jumlahbayar
		} else {
			return a.jumlahbayar > b.jumlahbayar
		}
	}
	return false
}

// SELECTION SORT (tidak diubah)
func selection(A *dataMahasiswa, kategori string, asc bool) {
	for i := 0; i < A.jmlMahasiswa-1; i++ {
		target := i
		for j := i + 1; j < A.jmlMahasiswa; j++ {
			if terkecil(A.data[j], A.data[target], kategori, asc) {
				target = j
			}
		}
		if target != i {
			A.data[i], A.data[target] = A.data[target], A.data[i]
		}
	}
}

// INSERTION SORT (tidak diubah)
func insertion(A *dataMahasiswa, kategori string, asc bool) {
	for i := 1; i < A.jmlMahasiswa; i++ {
		kunci := A.data[i]
		j := i - 1
		for j >= 0 && terkecil(kunci, A.data[j], kategori, asc) {
			A.data[j+1] = A.data[j]
			j--
		}
		A.data[j+1] = kunci
	}
}

// URUTKAN DATA (tampilan rapi)
func urutdata(A *dataMahasiswa) {
	if A.jmlMahasiswa == 0 {
		fmt.Println("\n⚠️  Data kosong, tidak bisa diurutkan.")
		return
	}

	fmt.Println("\n========================================")
	fmt.Println("        PENGURUTAN DATA MAHASISWA")
	fmt.Println("========================================")
	fmt.Println(" Pilih metode sorting:")
	fmt.Println(" 1. Selection Sort")
	fmt.Println(" 2. Insertion Sort")
	fmt.Print(" Pilih (1/2): ")
	var menu int
	fmt.Scan(&menu)
	fmt.Scanln()

	fmt.Println("\n Pilih kategori:")
	fmt.Println(" 1. NIM")
	fmt.Println(" 2. Nama")
	fmt.Println(" 3. Tunggakan")
	fmt.Println(" 4. Jumlah Bayar")
	fmt.Print(" Pilih (1-4): ")
	var kat int
	fmt.Scan(&kat)
	fmt.Scanln()

	var kategori string
	switch kat {
	case 1:
		kategori = "NIM"
	case 2:
		kategori = "Nama"
	case 3:
		kategori = "Tunggakan"
	case 4:
		kategori = "Jumlah Bayar"
	default:
		fmt.Println(" ⚠️  Kategori tidak valid!")
		return
	}

	fmt.Println("\n Pilih urutan:")
	fmt.Println(" 1. Ascending (kecil ke besar)")
	fmt.Println(" 2. Descending (besar ke kecil)")
	fmt.Print(" Pilih (1/2): ")
	var urut int
	fmt.Scan(&urut)
	fmt.Scanln()
	asc := (urut == 1)

	if menu == 1 {
		selection(A, kategori, asc)
		fmt.Println(" ✅ Data berhasil diurutkan dengan Selection Sort.")
	} else if menu == 2 {
		insertion(A, kategori, asc)
		fmt.Println(" ✅ Data berhasil diurutkan dengan Insertion Sort.")
	} else {
		fmt.Println(" ⚠️  Metode sorting tidak valid!")
		return
	}
	tampilkan(A)
}

// STATISTIK KAS (rapi)
func statistikKas(A *dataMahasiswa) {
	fmt.Println("\n========================================")
	fmt.Println("          STATISTIK KAS MAHASISWA")
	fmt.Println("========================================")

	if A.jmlMahasiswa == 0 {
		fmt.Println(" ⚠️  Belum ada data mahasiswa.")
		return
	}

	totalKas := 0
	totalTunggakan := 0
	jmlLunas := 0
	maxBayar := -1
	idxMaxBayar := -1

	for i := 0; i < A.jmlMahasiswa; i++ {
		totalKas += A.data[i].jumlahbayar
		totalTunggakan += A.data[i].tunggakan
		if A.data[i].statusbayar {
			jmlLunas++
		}
		if A.data[i].jumlahbayar > maxBayar {
			maxBayar = A.data[i].jumlahbayar
			idxMaxBayar = i
		}
	}

	rataRata := float64(totalKas) / float64(A.jmlMahasiswa)
	jmlBelumLunas := A.jmlMahasiswa - jmlLunas

	fmt.Printf(" Total Kas Terkumpul   : Rp %d\n", totalKas)
	fmt.Printf(" Rata-rata Pembayaran  : Rp %.2f\n", rataRata)
	fmt.Printf(" Mahasiswa Lunas       : %d orang\n", jmlLunas)
	fmt.Printf(" Mahasiswa Belum Lunas : %d orang\n", jmlBelumLunas)
	fmt.Printf(" Total Tunggakan       : Rp %d\n", totalTunggakan)

	if idxMaxBayar != -1 && maxBayar > 0 {
		fmt.Printf(" Pembayaran Tertinggi  : %s (NIM: %s) dengan total Rp %d\n",
			A.data[idxMaxBayar].nama, A.data[idxMaxBayar].nim, maxBayar)
	} else {
		fmt.Println(" Belum ada pembayaran yang tercatat.")
	}
	fmt.Println()
}
