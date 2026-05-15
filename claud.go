package main

import "fmt"

// ============================================================
// TIPE BENTUKAN
// ============================================================

type Payment struct {
	Amount int
	Date   string
	Note   string
}

type Student struct {
	NIM      string
	Name     string
	Payments [20]Payment
	PayCount int
	Active   bool
}

type DataStore struct {
	Students       [100]Student
	StudentCount   int
	MonthlyFee     int
	MonthsExpected int
}

var store DataStore

// ============================================================
// PROSEDUR UTILITAS
// ============================================================

// hitungDibayar mengembalikan total pembayaran seorang mahasiswa
func hitungDibayar(s Student) int {
	total := 0
	for i := 0; i < s.PayCount; i++ {
		total += s.Payments[i].Amount
	}
	return total
}

// hitungTunggakan mengembalikan sisa tunggakan (minimal 0)
func hitungTunggakan(dibayar int) int {
	wajib := store.MonthlyFee * store.MonthsExpected
	tunggakan := wajib - dibayar
	if tunggakan < 0 {
		return 0
	}
	return tunggakan
}

// hitungJumlahAktif mengembalikan jumlah mahasiswa yang aktif
func hitungJumlahAktif() int {
	jumlah := 0
	for i := 0; i < store.StudentCount; i++ {
		if store.Students[i].Active {
			jumlah++
		}
	}
	return jumlah
}

// ambilIndeksAktif mengisi arrIdx dengan indeks mahasiswa aktif dan mengembalikan jumlahnya
func ambilIndeksAktif(arrIdx *[100]int) int {
	n := 0
	for i := 0; i < store.StudentCount; i++ {
		if store.Students[i].Active {
			arrIdx[n] = i
			n++
		}
	}
	return n
}

// ============================================================
// PROSEDUR PENCARIAN
// ============================================================

// sequentialSearch mencari NIM secara sequential dalam semua mahasiswa aktif
// Mengembalikan indeks di store.Students, atau -1 jika tidak ditemukan
func sequentialSearch(nim string) int {
	for i := 0; i < store.StudentCount; i++ {
		if store.Students[i].Active && store.Students[i].NIM == nim {
			return i
		}
	}
	return -1
}

// sortNIMuntukBinarySearch mengurutkan arrIdx[0..n-1] berdasarkan NIM ascending (Selection Sort)
func sortNIMuntukBinarySearch(arrIdx *[100]int, n int) {
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if store.Students[arrIdx[j]].NIM < store.Students[arrIdx[min]].NIM {
				min = j
			}
		}
		if min != i {
			arrIdx[i], arrIdx[min] = arrIdx[min], arrIdx[i]
		}
	}
}

// binarySearchDalamArrIdx mencari NIM di dalam arrIdx yang sudah diurutkan berdasarkan NIM
// Mengembalikan indeks di store.Students, atau -1 jika tidak ditemukan
func binarySearchDalamArrIdx(arrIdx *[100]int, n int, nim string) int {
	lo, hi := 0, n-1
	for lo <= hi {
		mid := (lo + hi) / 2
		nimMid := store.Students[arrIdx[mid]].NIM
		if nimMid == nim {
			return arrIdx[mid]
		} else if nimMid < nim {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// binarySearch mencari NIM dengan binary search di seluruh mahasiswa aktif
// Mengembalikan indeks di store.Students, atau -1 jika tidak ditemukan
func binarySearch(nim string) int {
	var arrIdx [100]int
	n := ambilIndeksAktif(&arrIdx)
	sortNIMuntukBinarySearch(&arrIdx, n)
	return binarySearchDalamArrIdx(&arrIdx, n, nim)
}

// cariMahasiswaMenu menampilkan pilihan metode dan mengembalikan indeks hasil pencarian
func cariMahasiswaMenu(prompt string) int {
	fmt.Println("Pilih metode pencarian:")
	fmt.Println(" 1. Sequential Search")
	fmt.Println(" 2. Binary Search")
	fmt.Print("Pilih: ")
	var metode int
	fmt.Scan(&metode)

	var nimCari string
	fmt.Print(prompt)
	fmt.Scan(&nimCari)

	switch metode {
	case 1:
		return sequentialSearch(nimCari)
	case 2:
		return binarySearch(nimCari)
	default:
		fmt.Println("Metode tidak valid.")
		return -1
	}
}

// ============================================================
// PROSEDUR PENGURUTAN
// ============================================================

// selectionSortNama mengurutkan arrIdx[0..n-1] berdasarkan Nama (ascending/descending)
func selectionSortNama(arrIdx *[100]int, n, arah int) {
	for i := 0; i < n-1; i++ {
		sel := i
		for j := i + 1; j < n; j++ {
			namaJ := store.Students[arrIdx[j]].Name
			namaSel := store.Students[arrIdx[sel]].Name
			if (arah == 1 && namaJ < namaSel) || (arah == 2 && namaJ > namaSel) {
				sel = j
			}
		}
		if sel != i {
			arrIdx[i], arrIdx[sel] = arrIdx[sel], arrIdx[i]
		}
	}
}

// insertionSortTunggakan mengurutkan arrIdx[0..n-1] berdasarkan Tunggakan (ascending/descending)
func insertionSortTunggakan(arrIdx *[100]int, n, arah int) {
	for i := 1; i < n; i++ {
		key := arrIdx[i]
		tunggakanKey := hitungTunggakan(hitungDibayar(store.Students[key]))

		j := i - 1
		for j >= 0 {
			tunggakanJ := hitungTunggakan(hitungDibayar(store.Students[arrIdx[j]]))
			harus := (arah == 1 && tunggakanJ > tunggakanKey) ||
				(arah == 2 && tunggakanJ < tunggakanKey)
			if harus {
				arrIdx[j+1] = arrIdx[j]
				j--
			} else {
				break
			}
		}
		arrIdx[j+1] = key
	}
}

// ============================================================
// PROSEDUR TAMPILAN
// ============================================================

// tampilkanMenu menampilkan menu utama
func tampilkanMenu() {
	fmt.Println("\n==================================================")
	fmt.Println("      SIKAS - Sistem Informasi Kas Mahasiswa")
	fmt.Println("==================================================")
	fmt.Println(" 1. Tambah Mahasiswa")
	fmt.Println(" 2. Ubah Nama Mahasiswa")
	fmt.Println(" 3. Hapus Mahasiswa")
	fmt.Println(" 4. Lihat Semua Mahasiswa")
	fmt.Println(" 5. Catat Pembayaran")
	fmt.Println(" 6. Cari Mahasiswa Belum Bayar")
	fmt.Println(" 7. Urutkan Data Mahasiswa")
	fmt.Println(" 8. Statistik Kas")
	fmt.Println(" 9. Pengaturan Iuran")
	fmt.Println(" 0. Keluar")
	fmt.Println("--------------------------------------------------")
	fmt.Print("Pilih menu: ")
}

// tampilkanTabelMahasiswa menampilkan tabel mahasiswa berdasarkan urutan arrIdx[0..n-1]
func tampilkanTabelMahasiswa(arrIdx [100]int, n int) {
	fmt.Printf("%-12s %-20s %12s %12s %s\n", "NIM", "Nama", "Dibayar", "Tunggakan", "Status")
	fmt.Println("--------------------------------------------------")
	for i := 0; i < n; i++ {
		idx := arrIdx[i]
		dibayar := hitungDibayar(store.Students[idx])
		tunggakan := hitungTunggakan(dibayar)
		status := "LUNAS"
		if tunggakan > 0 {
			status = "BELUM LUNAS"
		}
		fmt.Printf("%-12s %-20s %12d %12d %s\n",
			store.Students[idx].NIM, store.Students[idx].Name,
			dibayar, tunggakan, status)
	}
}

// ============================================================
// PROSEDUR MENU (FITUR)
// ============================================================

// tambahMahasiswa menangani proses penambahan mahasiswa baru
func tambahMahasiswa() {
	if store.StudentCount >= 100 {
		fmt.Println("Data mahasiswa sudah penuh (maks 100).")
		return
	}

	var nim string
	fmt.Print("NIM           : ")
	fmt.Scan(&nim)

	if sequentialSearch(nim) != -1 {
		fmt.Println("NIM sudah terdaftar.")
		return
	}

	var nama string
	fmt.Print("Nama Mahasiswa: ")
	fmt.Scan(&nama)

	s := &store.Students[store.StudentCount]
	s.NIM = nim
	s.Name = nama
	s.PayCount = 0
	s.Active = true
	store.StudentCount++
	fmt.Println("Mahasiswa berhasil ditambahkan.")
}

// ubahNamaMahasiswa menangani proses pengubahan nama mahasiswa
func ubahNamaMahasiswa() {
	if hitungJumlahAktif() == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	idx := cariMahasiswaMenu("NIM yang akan diubah: ")
	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	fmt.Printf("Nama lama: %s\n", store.Students[idx].Name)
	var namaBaru string
	fmt.Print("Nama baru : ")
	fmt.Scan(&namaBaru)
	store.Students[idx].Name = namaBaru
	fmt.Println("Nama berhasil diubah.")
}

// hapusMahasiswa menangani proses penghapusan (soft delete) mahasiswa
func hapusMahasiswa() {
	if hitungJumlahAktif() == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	idx := cariMahasiswaMenu("NIM yang akan dihapus: ")
	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	fmt.Printf("Hapus %s (%s)? (y/n): ", store.Students[idx].NIM, store.Students[idx].Name)
	var konfirm string
	fmt.Scan(&konfirm)
	if konfirm == "y" || konfirm == "Y" {
		store.Students[idx].Active = false
		fmt.Println("Mahasiswa berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

// lihatSemuaMahasiswa menampilkan daftar seluruh mahasiswa aktif
func lihatSemuaMahasiswa() {
	var arrIdx [100]int
	n := ambilIndeksAktif(&arrIdx)

	if n == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	fmt.Println("\n--------------------------------------------------")
	tampilkanTabelMahasiswa(arrIdx, n)
	fmt.Println("--------------------------------------------------")
}

// catatPembayaran menangani proses pencatatan pembayaran mahasiswa
func catatPembayaran() {
	if hitungJumlahAktif() == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	var nimCari string
	fmt.Print("NIM mahasiswa: ")
	fmt.Scan(&nimCari)

	idx := sequentialSearch(nimCari)
	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}
	if store.Students[idx].PayCount >= 20 {
		fmt.Println("Riwayat pembayaran penuh (maks 20).")
		return
	}

	fmt.Printf("Mahasiswa: %s - %s\n", store.Students[idx].NIM, store.Students[idx].Name)

	var amount int
	fmt.Print("Nominal (Rp): ")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Nominal harus lebih dari 0.")
		return
	}

	var tanggal, catatan string
	fmt.Print("Tanggal (DD/MM/YYYY): ")
	fmt.Scan(&tanggal)
	fmt.Print("Catatan: ")
	fmt.Scan(&catatan)

	pi := store.Students[idx].PayCount
	store.Students[idx].Payments[pi] = Payment{Amount: amount, Date: tanggal, Note: catatan}
	store.Students[idx].PayCount++
	fmt.Printf("Pembayaran Rp%d berhasil dicatat.\n", amount)
}

// cariMahasiswaBelumBayar menampilkan daftar dan mencari mahasiswa yang belum lunas
func cariMahasiswaBelumBayar() {
	// Kumpulkan mahasiswa belum lunas
	var unpaidIdx [100]int
	jumlahUnpaid := 0
	wajib := store.MonthlyFee * store.MonthsExpected

	for i := 0; i < store.StudentCount; i++ {
		if store.Students[i].Active && hitungDibayar(store.Students[i]) < wajib {
			unpaidIdx[jumlahUnpaid] = i
			jumlahUnpaid++
		}
	}

	if jumlahUnpaid == 0 {
		fmt.Println("Semua mahasiswa sudah lunas.")
		return
	}

	// Tampilkan daftar belum lunas
	fmt.Println("\n--- Daftar Mahasiswa Belum Lunas ---")
	fmt.Printf("%-12s %-20s %12s\n", "NIM", "Nama", "Tunggakan")
	fmt.Println("--------------------------------------------")
	for i := 0; i < jumlahUnpaid; i++ {
		idx := unpaidIdx[i]
		tunggakan := hitungTunggakan(hitungDibayar(store.Students[idx]))
		fmt.Printf("%-12s %-20s %12d\n",
			store.Students[idx].NIM, store.Students[idx].Name, tunggakan)
	}
	fmt.Println("--------------------------------------------")

	// Pilih metode pencarian dalam daftar belum lunas
	fmt.Println("Pilih metode pencarian:")
	fmt.Println(" 1. Sequential Search")
	fmt.Println(" 2. Binary Search")
	fmt.Print("Pilih: ")
	var metode int
	fmt.Scan(&metode)

	var nimCari string
	fmt.Print("NIM yang dicari: ")
	fmt.Scan(&nimCari)

	idx := -1
	switch metode {
	case 1:
		for i := 0; i < jumlahUnpaid && idx == -1; i++ {
			if store.Students[unpaidIdx[i]].NIM == nimCari {
				idx = unpaidIdx[i]
			}
		}
	case 2:
		sortNIMuntukBinarySearch(&unpaidIdx, jumlahUnpaid)
		idx = binarySearchDalamArrIdx(&unpaidIdx, jumlahUnpaid, nimCari)
	default:
		fmt.Println("Metode tidak valid.")
		return
	}

	if idx == -1 {
		fmt.Println("Tidak ditemukan dalam daftar belum lunas.")
		return
	}

	dibayar := hitungDibayar(store.Students[idx])
	tunggakan := hitungTunggakan(dibayar)
	fmt.Println("\n--- Hasil Pencarian ---")
	fmt.Printf("NIM      : %s\n", store.Students[idx].NIM)
	fmt.Printf("Nama     : %s\n", store.Students[idx].Name)
	fmt.Printf("Dibayar  : Rp%d\n", dibayar)
	fmt.Printf("Tunggakan: Rp%d\n", tunggakan)
}

// urutkanDataMahasiswa menampilkan menu pengurutan dan menampilkan hasil urutan
func urutkanDataMahasiswa() {
	var arrIdx [100]int
	n := ambilIndeksAktif(&arrIdx)

	if n == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	fmt.Println("Urut berdasarkan:")
	fmt.Println(" 1. Nama           (Selection Sort)")
	fmt.Println(" 2. Total Tunggakan (Insertion Sort)")
	fmt.Print("Pilih: ")
	var kategori int
	fmt.Scan(&kategori)

	fmt.Println("Arah:")
	fmt.Println(" 1. Ascending  (A-Z / Kecil ke Besar)")
	fmt.Println(" 2. Descending (Z-A / Besar ke Kecil)")
	fmt.Print("Pilih: ")
	var arah int
	fmt.Scan(&arah)

	labelArah := map[int]string{1: "Ascending", 2: "Descending"}

	switch kategori {
	case 1:
		selectionSortNama(&arrIdx, n, arah)
		fmt.Printf("\n--- Urut Nama %s (Selection Sort) ---\n", labelArah[arah])
	case 2:
		insertionSortTunggakan(&arrIdx, n, arah)
		fmt.Printf("\n--- Urut Tunggakan %s (Insertion Sort) ---\n", labelArah[arah])
	default:
		fmt.Println("Kategori tidak valid.")
		return
	}

	tampilkanTabelMahasiswa(arrIdx, n)
}

// statistikKas menampilkan ringkasan statistik keuangan kas
func statistikKas() {
	totalKas := 0
	jumlahLunas := 0
	jumlahAktif := 0
	wajib := store.MonthlyFee * store.MonthsExpected

	for i := 0; i < store.StudentCount; i++ {
		if store.Students[i].Active {
			jumlahAktif++
			dibayar := hitungDibayar(store.Students[i])
			totalKas += dibayar
			if dibayar >= wajib {
				jumlahLunas++
			}
		}
	}

	fmt.Println("\n========== STATISTIK KAS ==========")
	fmt.Printf("Total saldo kas terkumpul : Rp%d\n", totalKas)
	fmt.Printf("Mahasiswa lunas           : %d dari %d\n", jumlahLunas, jumlahAktif)
	fmt.Printf("Mahasiswa belum lunas     : %d\n", jumlahAktif-jumlahLunas)
	fmt.Printf("Iuran per bulan           : Rp%d\n", store.MonthlyFee)
	fmt.Printf("Jumlah bulan wajib        : %d bulan\n", store.MonthsExpected)
	fmt.Printf("Total wajib per mahasiswa : Rp%d\n", wajib)
	fmt.Println("====================================")
}

// pengaturanIuran menangani perubahan iuran bulanan dan jumlah bulan wajib
func pengaturanIuran() {
	fmt.Printf("Iuran per bulan saat ini: Rp%d\n", store.MonthlyFee)
	fmt.Print("Iuran baru (0 = tidak ubah): Rp")
	var fee int
	fmt.Scan(&fee)
	if fee > 0 {
		store.MonthlyFee = fee
		fmt.Printf("Iuran diubah menjadi Rp%d\n", store.MonthlyFee)
	}

	fmt.Printf("Jumlah bulan wajib saat ini: %d\n", store.MonthsExpected)
	fmt.Print("Jumlah bulan baru (0 = tidak ubah): ")
	var bulan int
	fmt.Scan(&bulan)
	if bulan > 0 {
		store.MonthsExpected = bulan
		fmt.Printf("Jumlah bulan diubah menjadi %d\n", store.MonthsExpected)
	}
}

// ============================================================
// MAIN
// ============================================================

func main() {
	store.MonthlyFee = 50000
	store.MonthsExpected = 6

	jalan := true
	for jalan {
		tampilkanMenu()

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahMahasiswa()
		case 2:
			ubahNamaMahasiswa()
		case 3:
			hapusMahasiswa()
		case 4:
			lihatSemuaMahasiswa()
		case 5:
			catatPembayaran()
		case 6:
			cariMahasiswaBelumBayar()
		case 7:
			urutkanDataMahasiswa()
		case 8:
			statistikKas()
		case 9:
			pengaturanIuran()
		case 0:
			fmt.Println("Terima kasih. Program selesai.")
			jalan = false
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}