package main

// @jebb_24
// SIKAS - Sistem Informasi Kas Mahasiswa
// Aplikasi untuk mengelola pencatatan iuran kas kelas secara digital.
// Pengguna: bendahara kelas.

import "fmt"

// ============================================================
// TIPE BENTUKAN
// ============================================================

// Payment menyimpan data satu transaksi pembayaran iuran (@jebb_24)
type Payment struct {
	Amount int    // nominal pembayaran dalam rupiah (@jebb_24)
	Date   string // tanggal pembayaran format DD/MM/YYYY (@jebb_24)
	Note   string // catatan tambahan untuk transaksi (@jebb_24)
}

// Student menyimpan identitas dan riwayat pembayaran seorang mahasiswa (@jebb_24)
type Student struct {
	NIM      string      // nomor induk mahasiswa yang bersifat unik (@jebb_24)
	Name     string      // nama lengkap mahasiswa (@jebb_24)
	Payments [20]Payment // array riwayat pembayaran maksimal 20 transaksi (@jebb_24)
	PayCount int         // jumlah transaksi yang sudah tercatat (@jebb_24)
	Active   bool        // status aktif; false berarti sudah dihapus secara soft delete (@jebb_24)
}

// DataStore adalah penyimpanan utama seluruh data aplikasi (@jebb_24)
type DataStore struct {
	Students       [100]Student // array data mahasiswa maksimal 100 orang (@jebb_24)
	StudentCount   int          // jumlah slot yang sudah terpakai di array Students (@jebb_24)
	MonthlyFee     int          // besar iuran per bulan dalam rupiah (@jebb_24)
	MonthsExpected int          // jumlah bulan yang wajib dibayar (@jebb_24)
}

// store adalah variabel global tunggal yang menjadi sumber data seluruh program (@jebb_24)
var store DataStore

// ============================================================
// PROSEDUR UTILITAS
// ============================================================

// hitungDibayar menjumlahkan seluruh nominal pembayaran milik satu mahasiswa (@jebb_24)
func hitungDibayar(s Student) int {
	total := 0
	// menjumlahkan setiap transaksi dari indeks 0 sampai PayCount-1 (@jebb_24)
	for i := 0; i < s.PayCount; i++ {
		total += s.Payments[i].Amount
	}
	return total
}

// hitungTunggakan menghitung sisa kewajiban yang belum dibayar; hasilnya minimal 0 (@jebb_24)
func hitungTunggakan(dibayar int) int {
	// total kewajiban adalah iuran per bulan dikali jumlah bulan wajib (@jebb_24)
	wajib := store.MonthlyFee * store.MonthsExpected
	tunggakan := wajib - dibayar
	// tunggakan tidak boleh negatif meski mahasiswa membayar lebih dari kewajiban (@jebb_24)
	if tunggakan < 0 {
		return 0
	}
	return tunggakan
}

// hitungJumlahAktif menghitung berapa mahasiswa yang berstatus aktif (@jebb_24)
func hitungJumlahAktif() int {
	jumlah := 0
	// memeriksa setiap slot yang sudah terpakai di array (@jebb_24)
	for i := 0; i < store.StudentCount; i++ {
		if store.Students[i].Active {
			jumlah++
		}
	}
	return jumlah
}

// ambilIndeksAktif mengisi arrIdx dengan indeks semua mahasiswa aktif
// dan mengembalikan jumlahnya (@jebb_24)
func ambilIndeksAktif(arrIdx *[100]int) int {
	n := 0
	// hanya memasukkan indeks jika mahasiswa berstatus aktif (@jebb_24)
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

// sequentialSearch mencari NIM secara linear dari awal hingga akhir array (@jebb_24)
// mengembalikan indeks di store.Students, atau -1 jika tidak ditemukan (@jebb_24)
func sequentialSearch(nim string) int {
	// memeriksa satu per satu hingga NIM cocok atau semua data habis diperiksa (@jebb_24)
	for i := 0; i < store.StudentCount; i++ {
		if store.Students[i].Active && store.Students[i].NIM == nim {
			return i
		}
	}
	// mengembalikan -1 sebagai tanda data tidak ditemukan (@jebb_24)
	return -1
}

// sortNIMuntukBinarySearch mengurutkan arrIdx[0..n-1] berdasarkan NIM secara ascending
// menggunakan Selection Sort agar siap digunakan oleh Binary Search (@jebb_24)
func sortNIMuntukBinarySearch(arrIdx *[100]int, n int) {
	// selection sort: mencari elemen terkecil dan menukarnya ke posisi i (@jebb_24)
	for i := 0; i < n-1; i++ {
		min := i
		// mencari indeks NIM terkecil di sisa array yang belum terurut (@jebb_24)
		for j := i + 1; j < n; j++ {
			if store.Students[arrIdx[j]].NIM < store.Students[arrIdx[min]].NIM {
				min = j
			}
		}
		// menukar posisi jika ditemukan elemen yang lebih kecil dari posisi min (@jebb_24)
		if min != i {
			arrIdx[i], arrIdx[min] = arrIdx[min], arrIdx[i]
		}
	}
}

// binarySearchDalamArrIdx mencari NIM di dalam arrIdx yang sudah diurutkan (@jebb_24)
// mengembalikan indeks di store.Students, atau -1 jika tidak ditemukan (@jebb_24)
func binarySearchDalamArrIdx(arrIdx *[100]int, n int, nim string) int {
	lo, hi := 0, n-1
	// membagi rentang pencarian menjadi dua setiap iterasi hingga ditemukan atau habis (@jebb_24)
	for lo <= hi {
		mid := (lo + hi) / 2
		nimMid := store.Students[arrIdx[mid]].NIM
		if nimMid == nim {
			// NIM ditemukan tepat di posisi tengah (@jebb_24)
			return arrIdx[mid]
		} else if nimMid < nim {
			// target ada di separuh kanan, geser batas bawah ke kanan (@jebb_24)
			lo = mid + 1
		} else {
			// target ada di separuh kiri, geser batas atas ke kiri (@jebb_24)
			hi = mid - 1
		}
	}
	// mengembalikan -1 jika seluruh rentang sudah diperiksa dan tidak ditemukan (@jebb_24)
	return -1
}

// binarySearch membungkus proses sort lalu binary search untuk seluruh mahasiswa aktif (@jebb_24)
// mengembalikan indeks di store.Students, atau -1 jika tidak ditemukan (@jebb_24)
func binarySearch(nim string) int {
	var arrIdx [100]int
	// mengumpulkan indeks mahasiswa aktif sebelum diurutkan (@jebb_24)
	n := ambilIndeksAktif(&arrIdx)
	// mengurutkan indeks berdasarkan NIM agar binary search dapat bekerja dengan benar (@jebb_24)
	sortNIMuntukBinarySearch(&arrIdx, n)
	return binarySearchDalamArrIdx(&arrIdx, n, nim)
}

// cariMahasiswaMenu menampilkan pilihan metode pencarian, membaca NIM dari pengguna,
// lalu mengembalikan indeks hasil pencarian di store.Students (@jebb_24)
func cariMahasiswaMenu(prompt string) int {
	// meminta pengguna memilih metode pencarian yang diinginkan (@jebb_24)
	fmt.Println("Pilih metode pencarian:")
	fmt.Println(" 1. Sequential Search")
	fmt.Println(" 2. Binary Search")
	fmt.Print("Pilih: ")
	var metode int
	fmt.Scan(&metode)

	var nimCari string
	fmt.Print(prompt)
	fmt.Scan(&nimCari)

	// mengarahkan ke fungsi pencarian sesuai pilihan pengguna (@jebb_24)
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

// selectionSortNama mengurutkan arrIdx[0..n-1] berdasarkan Nama
// menggunakan Selection Sort dengan arah ascending (1) atau descending (2) (@jebb_24)
func selectionSortNama(arrIdx *[100]int, n, arah int) {
	// pada setiap iterasi luar, mencari elemen terpilih (terkecil/terbesar) (@jebb_24)
	for i := 0; i < n-1; i++ {
		sel := i
		for j := i + 1; j < n; j++ {
			namaJ := store.Students[arrIdx[j]].Name
			namaSel := store.Students[arrIdx[sel]].Name
			// membandingkan nama sesuai arah yang dipilih pengguna (@jebb_24)
			if (arah == 1 && namaJ < namaSel) || (arah == 2 && namaJ > namaSel) {
				sel = j
			}
		}
		// menukar posisi jika elemen terpilih bukan di posisi saat ini (@jebb_24)
		if sel != i {
			arrIdx[i], arrIdx[sel] = arrIdx[sel], arrIdx[i]
		}
	}
}

// insertionSortTunggakan mengurutkan arrIdx[0..n-1] berdasarkan total tunggakan
// menggunakan Insertion Sort dengan arah ascending (1) atau descending (2) (@jebb_24)
func insertionSortTunggakan(arrIdx *[100]int, n, arah int) {
	// memproses setiap elemen mulai indeks ke-1 dan menyisipkannya ke posisi yang tepat (@jebb_24)
	for i := 1; i < n; i++ {
		key := arrIdx[i]
		// menghitung tunggakan elemen yang akan disisipkan (@jebb_24)
		tunggakanKey := hitungTunggakan(hitungDibayar(store.Students[key]))

		j := i - 1
		// menggeser elemen ke kanan selama kondisi urutan belum terpenuhi (@jebb_24)
		for j >= 0 {
			tunggakanJ := hitungTunggakan(hitungDibayar(store.Students[arrIdx[j]]))
			// menentukan apakah elemen j perlu digeser berdasarkan arah pengurutan (@jebb_24)
			harus := (arah == 1 && tunggakanJ > tunggakanKey) ||
				(arah == 2 && tunggakanJ < tunggakanKey)
			if harus {
				arrIdx[j+1] = arrIdx[j]
				j--
			} else {
				break
			}
		}
		// menempatkan key di posisi yang sesuai hasil penyisipan (@jebb_24)
		arrIdx[j+1] = key
	}
}

// ============================================================
// PROSEDUR TAMPILAN
// ============================================================

// tampilkanMenu mencetak menu utama beserta seluruh pilihan yang tersedia (@jebb_24)
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

// tampilkanTabelMahasiswa mencetak baris data setiap mahasiswa dalam arrIdx[0..n-1] (@jebb_24)
func tampilkanTabelMahasiswa(arrIdx [100]int, n int) {
	// mencetak header kolom tabel (@jebb_24)
	fmt.Printf("%-12s %-20s %12s %12s %s\n", "NIM", "Nama", "Dibayar", "Tunggakan", "Status")
	fmt.Println("--------------------------------------------------")
	// mencetak satu baris per mahasiswa sesuai urutan yang ada di arrIdx (@jebb_24)
	for i := 0; i < n; i++ {
		idx := arrIdx[i]
		dibayar := hitungDibayar(store.Students[idx])
		tunggakan := hitungTunggakan(dibayar)
		// menentukan label status berdasarkan ada tidaknya tunggakan (@jebb_24)
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

// tambahMahasiswa membaca data dari pengguna dan mendaftarkan mahasiswa baru (@jebb_24)
func tambahMahasiswa() {
	// memastikan array mahasiswa belum penuh sebelum menambah data (@jebb_24)
	if store.StudentCount >= 100 {
		fmt.Println("Data mahasiswa sudah penuh (maks 100).")
		return
	}

	var nim string
	fmt.Print("NIM           : ")
	fmt.Scan(&nim)

	// mencegah duplikasi NIM menggunakan sequential search (@jebb_24)
	if sequentialSearch(nim) != -1 {
		fmt.Println("NIM sudah terdaftar.")
		return
	}

	var nama string
	fmt.Print("Nama Mahasiswa: ")
	fmt.Scan(&nama)

	// mengisi slot berikutnya di array dan menandainya sebagai aktif (@jebb_24)
	s := &store.Students[store.StudentCount]
	s.NIM = nim
	s.Name = nama
	s.PayCount = 0
	s.Active = true
	store.StudentCount++
	fmt.Println("Mahasiswa berhasil ditambahkan.")
}

// ubahNamaMahasiswa mencari mahasiswa berdasarkan NIM lalu memperbarui namanya (@jebb_24)
func ubahNamaMahasiswa() {
	// menghentikan proses jika belum ada mahasiswa yang terdaftar (@jebb_24)
	if hitungJumlahAktif() == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	// menggunakan prosedur pencarian terpusat untuk mendapatkan indeks (@jebb_24)
	idx := cariMahasiswaMenu("NIM yang akan diubah: ")
	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	// menampilkan nama lama terlebih dahulu sebelum meminta nama baru (@jebb_24)
	fmt.Printf("Nama lama: %s\n", store.Students[idx].Name)
	var namaBaru string
	fmt.Print("Nama baru : ")
	fmt.Scan(&namaBaru)
	store.Students[idx].Name = namaBaru
	fmt.Println("Nama berhasil diubah.")
}

// hapusMahasiswa menandai mahasiswa sebagai tidak aktif menggunakan soft delete (@jebb_24)
func hapusMahasiswa() {
	// menghentikan proses jika belum ada mahasiswa yang terdaftar (@jebb_24)
	if hitungJumlahAktif() == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	// menggunakan prosedur pencarian terpusat untuk mendapatkan indeks (@jebb_24)
	idx := cariMahasiswaMenu("NIM yang akan dihapus: ")
	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	// meminta konfirmasi pengguna sebelum melakukan penghapusan (@jebb_24)
	fmt.Printf("Hapus %s (%s)? (y/n): ", store.Students[idx].NIM, store.Students[idx].Name)
	var konfirm string
	fmt.Scan(&konfirm)
	if konfirm == "y" || konfirm == "Y" {
		// menandai Active = false sebagai pengganti penghapusan fisik dari array (@jebb_24)
		store.Students[idx].Active = false
		fmt.Println("Mahasiswa berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

// lihatSemuaMahasiswa menampilkan tabel seluruh mahasiswa yang masih aktif (@jebb_24)
func lihatSemuaMahasiswa() {
	var arrIdx [100]int
	// mengumpulkan indeks mahasiswa aktif ke dalam arrIdx (@jebb_24)
	n := ambilIndeksAktif(&arrIdx)

	if n == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	fmt.Println("\n--------------------------------------------------")
	// mencetak tabel menggunakan prosedur tampilan bersama (@jebb_24)
	tampilkanTabelMahasiswa(arrIdx, n)
	fmt.Println("--------------------------------------------------")
}

// catatPembayaran menambahkan satu transaksi pembayaran ke riwayat mahasiswa (@jebb_24)
func catatPembayaran() {
	// menghentikan proses jika belum ada mahasiswa yang terdaftar (@jebb_24)
	if hitungJumlahAktif() == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	var nimCari string
	fmt.Print("NIM mahasiswa: ")
	fmt.Scan(&nimCari)

	// menggunakan sequential search untuk menemukan mahasiswa yang dituju (@jebb_24)
	idx := sequentialSearch(nimCari)
	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}
	// memastikan slot riwayat pembayaran belum penuh (@jebb_24)
	if store.Students[idx].PayCount >= 20 {
		fmt.Println("Riwayat pembayaran penuh (maks 20).")
		return
	}

	fmt.Printf("Mahasiswa: %s - %s\n", store.Students[idx].NIM, store.Students[idx].Name)

	var amount int
	fmt.Print("Nominal (Rp): ")
	fmt.Scan(&amount)
	// memvalidasi bahwa nominal pembayaran bernilai positif (@jebb_24)
	if amount <= 0 {
		fmt.Println("Nominal harus lebih dari 0.")
		return
	}

	var tanggal, catatan string
	fmt.Print("Tanggal (DD/MM/YYYY): ")
	fmt.Scan(&tanggal)
	fmt.Print("Catatan: ")
	fmt.Scan(&catatan)

	// menyimpan data transaksi di slot berikutnya lalu menambah PayCount (@jebb_24)
	pi := store.Students[idx].PayCount
	store.Students[idx].Payments[pi] = Payment{Amount: amount, Date: tanggal, Note: catatan}
	store.Students[idx].PayCount++
	fmt.Printf("Pembayaran Rp%d berhasil dicatat.\n", amount)
}

// cariMahasiswaBelumBayar menampilkan daftar mahasiswa yang belum lunas,
// lalu memungkinkan pencarian individual dengan Sequential atau Binary Search (@jebb_24)
func cariMahasiswaBelumBayar() {
	// mengumpulkan indeks mahasiswa yang total bayarnya kurang dari kewajiban (@jebb_24)
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

	// mencetak daftar mahasiswa belum lunas beserta sisa tunggakannya (@jebb_24)
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

	// meminta pengguna memilih metode pencarian dalam daftar belum lunas (@jebb_24)
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
		// menelusuri daftar belum lunas satu per satu secara sequential (@jebb_24)
		for i := 0; i < jumlahUnpaid && idx == -1; i++ {
			if store.Students[unpaidIdx[i]].NIM == nimCari {
				idx = unpaidIdx[i]
			}
		}
	case 2:
		// mengurutkan daftar belum lunas berdasarkan NIM sebelum binary search (@jebb_24)
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

	// menampilkan detail mahasiswa yang berhasil ditemukan (@jebb_24)
	dibayar := hitungDibayar(store.Students[idx])
	tunggakan := hitungTunggakan(dibayar)
	fmt.Println("\n--- Hasil Pencarian ---")
	fmt.Printf("NIM      : %s\n", store.Students[idx].NIM)
	fmt.Printf("Nama     : %s\n", store.Students[idx].Name)
	fmt.Printf("Dibayar  : Rp%d\n", dibayar)
	fmt.Printf("Tunggakan: Rp%d\n", tunggakan)
}

// urutkanDataMahasiswa menampilkan menu pilihan urutan lalu mencetak tabel yang sudah diurut (@jebb_24)
func urutkanDataMahasiswa() {
	var arrIdx [100]int
	// mengumpulkan indeks mahasiswa aktif sebelum diurutkan (@jebb_24)
	n := ambilIndeksAktif(&arrIdx)

	if n == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	// meminta pengguna memilih kriteria pengurutan (@jebb_24)
	fmt.Println("Urut berdasarkan:")
	fmt.Println(" 1. Nama           (Selection Sort)")
	fmt.Println(" 2. Total Tunggakan (Insertion Sort)")
	fmt.Print("Pilih: ")
	var kategori int
	fmt.Scan(&kategori)

	// meminta pengguna memilih arah pengurutan (@jebb_24)
	fmt.Println("Arah:")
	fmt.Println(" 1. Ascending  (A-Z / Kecil ke Besar)")
	fmt.Println(" 2. Descending (Z-A / Besar ke Kecil)")
	fmt.Print("Pilih: ")
	var arah int
	fmt.Scan(&arah)

	// memetakan angka arah ke label teks untuk ditampilkan di judul hasil (@jebb_24)
	labelArah := map[int]string{1: "Ascending", 2: "Descending"}

	switch kategori {
	case 1:
		// mengurutkan berdasarkan nama menggunakan Selection Sort (@jebb_24)
		selectionSortNama(&arrIdx, n, arah)
		fmt.Printf("\n--- Urut Nama %s (Selection Sort) ---\n", labelArah[arah])
	case 2:
		// mengurutkan berdasarkan tunggakan menggunakan Insertion Sort (@jebb_24)
		insertionSortTunggakan(&arrIdx, n, arah)
		fmt.Printf("\n--- Urut Tunggakan %s (Insertion Sort) ---\n", labelArah[arah])
	default:
		fmt.Println("Kategori tidak valid.")
		return
	}

	// mencetak tabel hasil pengurutan menggunakan prosedur bersama (@jebb_24)
	tampilkanTabelMahasiswa(arrIdx, n)
}

// statistikKas menghitung dan menampilkan ringkasan keuangan kas kelas (@jebb_24)
func statistikKas() {
	totalKas := 0
	jumlahLunas := 0
	jumlahAktif := 0
	wajib := store.MonthlyFee * store.MonthsExpected

	// mengakumulasi data dari setiap mahasiswa yang masih aktif (@jebb_24)
	for i := 0; i < store.StudentCount; i++ {
		if store.Students[i].Active {
			jumlahAktif++
			dibayar := hitungDibayar(store.Students[i])
			totalKas += dibayar
			// mahasiswa dianggap lunas jika total bayarnya memenuhi seluruh kewajiban (@jebb_24)
			if dibayar >= wajib {
				jumlahLunas++
			}
		}
	}

	// mencetak ringkasan statistik kas dalam format yang rapi (@jebb_24)
	fmt.Println("\n========== STATISTIK KAS ==========")
	fmt.Printf("Total saldo kas terkumpul : Rp%d\n", totalKas)
	fmt.Printf("Mahasiswa lunas           : %d dari %d\n", jumlahLunas, jumlahAktif)
	fmt.Printf("Mahasiswa belum lunas     : %d\n", jumlahAktif-jumlahLunas)
	fmt.Printf("Iuran per bulan           : Rp%d\n", store.MonthlyFee)
	fmt.Printf("Jumlah bulan wajib        : %d bulan\n", store.MonthsExpected)
	fmt.Printf("Total wajib per mahasiswa : Rp%d\n", wajib)
	fmt.Println("====================================")
}

// pengaturanIuran memungkinkan bendahara mengubah besaran iuran dan jumlah bulan wajib (@jebb_24)
func pengaturanIuran() {
	// menampilkan nilai iuran saat ini sebelum meminta perubahan (@jebb_24)
	fmt.Printf("Iuran per bulan saat ini: Rp%d\n", store.MonthlyFee)
	fmt.Print("Iuran baru (0 = tidak ubah): Rp")
	var fee int
	fmt.Scan(&fee)
	// hanya memperbarui iuran jika nilai baru lebih dari 0 (@jebb_24)
	if fee > 0 {
		store.MonthlyFee = fee
		fmt.Printf("Iuran diubah menjadi Rp%d\n", store.MonthlyFee)
	}

	// menampilkan jumlah bulan saat ini sebelum meminta perubahan (@jebb_24)
	fmt.Printf("Jumlah bulan wajib saat ini: %d\n", store.MonthsExpected)
	fmt.Print("Jumlah bulan baru (0 = tidak ubah): ")
	var bulan int
	fmt.Scan(&bulan)
	// hanya memperbarui jumlah bulan jika nilai baru lebih dari 0 (@jebb_24)
	if bulan > 0 {
		store.MonthsExpected = bulan
		fmt.Printf("Jumlah bulan diubah menjadi %d\n", store.MonthsExpected)
	}
}

// ============================================================
// MAIN
// ============================================================

// main adalah titik masuk program; menginisialisasi data awal dan menjalankan loop menu (@jebb_24)
func main() {
	// menetapkan iuran default dan jumlah bulan wajib saat program pertama dijalankan (@jebb_24)
	store.MonthlyFee = 50000
	store.MonthsExpected = 6

	jalan := true
	// menjalankan loop utama program selama pengguna belum memilih keluar (@jebb_24)
	for jalan {
		tampilkanMenu()

		var pilih int
		fmt.Scan(&pilih)

		// mengarahkan eksekusi ke prosedur yang sesuai berdasarkan pilihan pengguna (@jebb_24)
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