package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Payment merepresentasikan riwayat pembayaran
type Payment struct {
	Amount int    `json:"amount"`
	Date   string `json:"date"`
	Note   string `json:"note"`
}

// Student menggunakan NIM sebagai ID unik
type Student struct {
	NIM      string    `json:"nim"`
	Name     string    `json:"name"`
	Payments []Payment `json:"payments"`
}

// DataStore untuk menyimpan semua data aplikasi
type DataStore struct {
	Students       []Student `json:"students"`
	MonthlyFee     int       `json:"monthly_fee"`
	MonthsExpected int       `json:"months_expected"`
}

var (
	store   DataStore
	scanner = bufio.NewScanner(os.Stdin)
)

func init() {
	store.MonthlyFee = 50000
	store.MonthsExpected = 6
}

func main() {
	loadData()
	for {
		showMenu()
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		switch choice {
		case "1":
			manageStudents()
		case "2":
			recordPayment()
		case "3":
			searchUnpaidStudents()
		case "4":
			sortMenu()
		case "5":
			showStatistics()
		case "6":
			settingsMenu()
		case "7":
			saveData()
			fmt.Println("\n✓ Data disimpan. Terima kasih!")
			return
		case "8":
			exportMenu()
		default:
			fmt.Println("✗ Pilihan tidak valid.")
		}
	}
}

func showMenu() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("     SIKAS - Informasi Kas Mahasiswa")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println(" 1. Kelola Data Mahasiswa")
	fmt.Println(" 2. Catat Pembayaran Iuran")
	fmt.Println(" 3. Cari Mahasiswa Belum Bayar")
	fmt.Println(" 4. Urutkan Data Mahasiswa")
	fmt.Println(" 5. Statistik Kas")
	fmt.Println(" 6. Pengaturan Iuran")
	fmt.Println(" 7. Simpan & Keluar")
	fmt.Println(" 8. Ekspor Data ke File (CSV/TXT)")
	fmt.Println(strings.Repeat("-", 50))
}

// ==================== MANAJEMEN MAHASISWA ====================

func manageStudents() {
	for {
		fmt.Println("\n" + strings.Repeat("-", 40))
		fmt.Println("     MANAJEMEN MAHASISWA")
		fmt.Println(strings.Repeat("-", 40))
		fmt.Println(" 1. Tambah Mahasiswa")
		fmt.Println(" 2. Ubah Mahasiswa")
		fmt.Println(" 3. Hapus Mahasiswa")
		fmt.Println(" 4. Lihat Semua Mahasiswa")
		fmt.Println(" 5. Kembali")
		fmt.Print("Pilih: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		switch choice {
		case "1":
			addStudent()
		case "2":
			editStudent()
		case "3":
			deleteStudent()
		case "4":
			viewAllStudents()
		case "5":
			return
		default:
			fmt.Println("✗ Pilihan tidak valid.")
		}
	}
}

func addStudent() {
	fmt.Print("NIM: ")
	scanner.Scan()
	nim := strings.TrimSpace(scanner.Text())
	if nim == "" {
		fmt.Println("✗ NIM tidak boleh kosong!")
		return
	}
	if findStudentIndexByNIM(nim) != -1 {
		fmt.Println("✗ Mahasiswa dengan NIM tersebut sudah ada.")
		return
	}
	fmt.Print("Nama Mahasiswa: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	if name == "" {
		fmt.Println("✗ Nama tidak boleh kosong!")
		return
	}
	student := Student{
		NIM:      nim,
		Name:     name,
		Payments: []Payment{},
	}
	store.Students = append(store.Students, student)
	fmt.Printf("✓ Mahasiswa dengan NIM %s (%s) berhasil ditambahkan.\n", nim, name)
	saveData()
}

func editStudent() {
	if len(store.Students) == 0 {
		fmt.Println("✗ Belum ada data mahasiswa.")
		return
	}
	fmt.Print("Masukkan NIM mahasiswa yang akan diubah: ")
	scanner.Scan()
	nim := strings.TrimSpace(scanner.Text())
	idx := findStudentIndexByNIM(nim)
	if idx == -1 {
		fmt.Println("✗ Mahasiswa tidak ditemukan.")
		return
	}
	fmt.Printf("Nama lama: %s\n", store.Students[idx].Name)
	fmt.Print("Nama baru: ")
	scanner.Scan()
	newName := strings.TrimSpace(scanner.Text())
	if newName != "" {
		store.Students[idx].Name = newName
		fmt.Println("✓ Data berhasil diubah.")
		saveData()
	} else {
		fmt.Println("✗ Nama tidak boleh kosong, perubahan dibatalkan.")
	}
}

func deleteStudent() {
	if len(store.Students) == 0 {
		fmt.Println("✗ Belum ada data mahasiswa.")
		return
	}
	fmt.Print("Masukkan NIM mahasiswa yang akan dihapus: ")
	scanner.Scan()
	nim := strings.TrimSpace(scanner.Text())
	idx := findStudentIndexByNIM(nim)
	if idx == -1 {
		fmt.Println("✗ Mahasiswa tidak ditemukan.")
		return
	}
	fmt.Printf("Anda yakin ingin menghapus %s (NIM: %s)? (y/n): ", store.Students[idx].Name, nim)
	scanner.Scan()
	confirm := strings.ToLower(strings.TrimSpace(scanner.Text()))
	if confirm == "y" {
		store.Students = append(store.Students[:idx], store.Students[idx+1:]...)
		fmt.Println("✓ Mahasiswa berhasil dihapus.")
		saveData()
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

func viewAllStudents() {
	if len(store.Students) == 0 {
		fmt.Println("✗ Belum ada data mahasiswa.")
		return
	}
	printStudentTable(store.Students, "DATA SELURUH MAHASISWA")
}

// Fungsi untuk mencetak tabel mahasiswa dengan rapi
func printStudentTable(students []Student, title string) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 80))

	// Header tabel
	fmt.Printf("| %-12s | %-20s | %13s | %13s | %-8s |\n", "NIM", "Nama", "Total Dibayar", "Tunggakan", "Status")
	fmt.Println(strings.Repeat("-", 80))

	expectedTotal := store.MonthlyFee * store.MonthsExpected
	for _, s := range students {
		totalPaid := totalPayments(s)
		arrears := expectedTotal - totalPaid
		if arrears < 0 {
			arrears = 0
		}
		status := "LUNAS"
		if arrears > 0 {
			status = "BELUM LUNAS"
		}
		fmt.Printf("| %-12s | %-20s | Rp%10d | Rp%10d | %-8s |\n",
			s.NIM, truncateString(s.Name, 20), totalPaid, arrears, status)
	}
	fmt.Println(strings.Repeat("=", 80))
}

// Memotong string jika terlalu panjang
func truncateString(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen-3] + "..."
	}
	return s
}

// ==================== PEMBAYARAN ====================

func recordPayment() {
	if len(store.Students) == 0 {
		fmt.Println("✗ Belum ada data mahasiswa. Tambahkan mahasiswa terlebih dahulu.")
		return
	}
	fmt.Print("Masukkan NIM mahasiswa: ")
	scanner.Scan()
	nim := strings.TrimSpace(scanner.Text())
	idx := findStudentIndexByNIM(nim)
	if idx == -1 {
		fmt.Println("✗ Mahasiswa tidak ditemukan.")
		return
	}
	fmt.Printf("Mahasiswa: %s (NIM: %s)\n", store.Students[idx].Name, nim)
	fmt.Print("Nominal iuran (Rp): ")
	scanner.Scan()
	amtStr := strings.TrimSpace(scanner.Text())
	amount, err := strconv.Atoi(amtStr)
	if err != nil || amount <= 0 {
		fmt.Println("✗ Nominal harus berupa angka positif.")
		return
	}

	currentDate := time.Now().Format("2006-01-02")
	fmt.Printf("Tanggal pembayaran (default %s): ", currentDate)
	scanner.Scan()
	dateInput := strings.TrimSpace(scanner.Text())
	if dateInput == "" {
		dateInput = currentDate
	}

	fmt.Print("Catatan (opsional): ")
	scanner.Scan()
	note := strings.TrimSpace(scanner.Text())

	payment := Payment{
		Amount: amount,
		Date:   dateInput,
		Note:   note,
	}
	store.Students[idx].Payments = append(store.Students[idx].Payments, payment)
	fmt.Printf("✓ Pembayaran Rp%d berhasil dicatat untuk %s.\n", amount, store.Students[idx].Name)
	saveData()
}

// ==================== PENCARIAN MAHASISWA BELUM BAYAR ====================

func searchUnpaidStudents() {
	unpaid := getUnpaidStudents()
	if len(unpaid) == 0 {
		fmt.Println("✓ Tidak ada mahasiswa dengan tunggakan.")
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("  DAFTAR MAHASISWA YANG BELUM LUNAS")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("| %-12s | %-20s | %13s |\n", "NIM", "Nama", "Tunggakan")
	fmt.Println(strings.Repeat("-", 70))
	for _, s := range unpaid {
		fmt.Printf("| %-12s | %-20s | Rp%10d |\n", s.NIM, truncateString(s.Name, 20), getArrears(s))
	}
	fmt.Println(strings.Repeat("=", 70))

	fmt.Println("\nPilih metode pencarian berdasarkan NIM:")
	fmt.Println(" 1. Sequential Search")
	fmt.Println(" 2. Binary Search (data akan diurutkan berdasarkan NIM)")
	fmt.Print("Pilih: ")
	scanner.Scan()
	method := strings.TrimSpace(scanner.Text())

	fmt.Print("Masukkan NIM mahasiswa yang dicari: ")
	scanner.Scan()
	targetNIM := strings.TrimSpace(scanner.Text())
	if targetNIM == "" {
		fmt.Println("✗ NIM tidak boleh kosong.")
		return
	}

	var found *Student
	if method == "1" {
		found = sequentialSearchByNIM(unpaid, targetNIM)
	} else if method == "2" {
		sortedUnpaid := make([]Student, len(unpaid))
		copy(sortedUnpaid, unpaid)
		sort.Slice(sortedUnpaid, func(i, j int) bool {
			return sortedUnpaid[i].NIM < sortedUnpaid[j].NIM
		})
		fmt.Println("✓ Data telah diurutkan berdasarkan NIM untuk binary search.")
		found = binarySearchByNIM(sortedUnpaid, targetNIM)
	} else {
		fmt.Println("✗ Metode tidak valid.")
		return
	}

	if found != nil {
		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Println("  HASIL PENCARIAN")
		fmt.Println(strings.Repeat("=", 50))
		fmt.Printf("NIM          : %s\n", found.NIM)
		fmt.Printf("Nama         : %s\n", found.Name)
		fmt.Printf("Total Dibayar: Rp%d\n", totalPayments(*found))
		fmt.Printf("Tunggakan    : Rp%d\n", getArrears(*found))
		fmt.Println(strings.Repeat("=", 50))
	} else {
		fmt.Printf("✗ Mahasiswa dengan NIM '%s' tidak ditemukan dalam daftar belum lunas.\n", targetNIM)
	}
}

func getUnpaidStudents() []Student {
	var unpaid []Student
	for _, s := range store.Students {
		if getArrears(s) > 0 {
			unpaid = append(unpaid, s)
		}
	}
	return unpaid
}

func sequentialSearchByNIM(students []Student, nim string) *Student {
	for i := range students {
		if students[i].NIM == nim {
			return &students[i]
		}
	}
	return nil
}

func binarySearchByNIM(students []Student, nim string) *Student {
	low, high := 0, len(students)-1
	for low <= high {
		mid := (low + high) / 2
		if students[mid].NIM == nim {
			return &students[mid]
		} else if students[mid].NIM < nim {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

// ==================== PENGURUTAN ====================

func sortMenu() {
	fmt.Println("\n" + strings.Repeat("-", 40))
	fmt.Println("     PENGURUTAN DATA MAHASISWA")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println(" 1. Urutkan berdasarkan Nama (Selection Sort)")
	fmt.Println(" 2. Urutkan berdasarkan Total Tunggakan (Insertion Sort)")
	fmt.Print("Pilih: ")
	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	if len(store.Students) == 0 {
		fmt.Println("✗ Belum ada data mahasiswa.")
		return
	}

	temp := make([]Student, len(store.Students))
	copy(temp, store.Students)

	switch choice {
	case "1":
		selectionSortByName(temp)
		printStudentTable(temp, "HASIL URUTAN BERDASARKAN NAMA (Selection Sort)")
	case "2":
		insertionSortByArrears(temp)
		printStudentTable(temp, "HASIL URUTAN BERDASARKAN TUNGGAKAN (Insertion Sort)")
	default:
		fmt.Println("✗ Pilihan tidak valid.")
	}
}

func selectionSortByName(students []Student) {
	n := len(students)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(students[j].Name) < strings.ToLower(students[minIdx].Name) {
				minIdx = j
			}
		}
		if minIdx != i {
			students[i], students[minIdx] = students[minIdx], students[i]
		}
	}
}

func insertionSortByArrears(students []Student) {
	n := len(students)
	for i := 1; i < n; i++ {
		key := students[i]
		keyArrears := getArrears(key)
		j := i - 1
		for j >= 0 && getArrears(students[j]) > keyArrears {
			students[j+1] = students[j]
			j--
		}
		students[j+1] = key
	}
}

// ==================== STATISTIK ====================

func showStatistics() {
	totalCash := 0
	fullyPaidCount := 0
	expectedTotal := store.MonthlyFee * store.MonthsExpected

	for _, s := range store.Students {
		totalPaid := totalPayments(s)
		totalCash += totalPaid
		if totalPaid >= expectedTotal {
			fullyPaidCount++
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("        STATISTIK KAS MAHASISWA")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf(" Total saldo kas terkumpul     : Rp%12d\n", totalCash)
	fmt.Printf(" Jumlah mahasiswa lunas        : %d dari %d\n", fullyPaidCount, len(store.Students))
	fmt.Printf(" Iuran per bulan               : Rp%12d\n", store.MonthlyFee)
	fmt.Printf(" Jumlah bulan wajib            : %12d\n", store.MonthsExpected)
	fmt.Printf(" Total iuran wajib per mhs     : Rp%12d\n", expectedTotal)
	fmt.Println(strings.Repeat("=", 50))
}

// ==================== PENGATURAN ====================

func settingsMenu() {
	fmt.Println("\n" + strings.Repeat("-", 40))
	fmt.Println("     PENGATURAN IURAN")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf(" Iuran per bulan saat ini: Rp%d\n", store.MonthlyFee)
	fmt.Print(" Masukkan iuran bulanan baru (atau tekan Enter untuk tetap): ")
	scanner.Scan()
	feeStr := strings.TrimSpace(scanner.Text())
	if feeStr != "" {
		newFee, err := strconv.Atoi(feeStr)
		if err == nil && newFee > 0 {
			store.MonthlyFee = newFee
			fmt.Printf(" ✓ Iuran per bulan diubah menjadi Rp%d\n", store.MonthlyFee)
		} else {
			fmt.Println(" ✗ Nominal tidak valid, iuran tetap.")
		}
	}

	fmt.Printf(" Jumlah bulan wajib saat ini: %d\n", store.MonthsExpected)
	fmt.Print(" Masukkan jumlah bulan wajib baru (atau tekan Enter untuk tetap): ")
	scanner.Scan()
	monthStr := strings.TrimSpace(scanner.Text())
	if monthStr != "" {
		newMonths, err := strconv.Atoi(monthStr)
		if err == nil && newMonths > 0 {
			store.MonthsExpected = newMonths
			fmt.Printf(" ✓ Jumlah bulan wajib diubah menjadi %d\n", store.MonthsExpected)
		} else {
			fmt.Println(" ✗ Jumlah bulan tidak valid, tetap.")
		}
	}
	saveData()
}

// ==================== EKSPOR DATA KE FILE ====================

func exportMenu() {
	fmt.Println("\n" + strings.Repeat("-", 40))
	fmt.Println("     EKSPOR DATA")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println(" 1. Ekspor ke CSV (bisa dibuka dengan Excel)")
	fmt.Println(" 2. Ekspor ke TXT (format teks terstruktur)")
	fmt.Print("Pilih: ")
	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	switch choice {
	case "1":
		exportToCSV()
	case "2":
		exportToTXT()
	default:
		fmt.Println("✗ Pilihan tidak valid.")
	}
}

func exportToCSV() {
	// File 1: Data Mahasiswa
	mahasiswaFile, err := os.Create("data_mahasiswa.csv")
	if err != nil {
		fmt.Println("✗ Gagal membuat file CSV:", err)
		return
	}
	defer mahasiswaFile.Close()

	writer := csv.NewWriter(mahasiswaFile)
	defer writer.Flush()

	header := []string{"NIM", "Nama", "Total Dibayar", "Tunggakan", "Status"}
	writer.Write(header)

	expectedTotal := store.MonthlyFee * store.MonthsExpected
	for _, s := range store.Students {
		totalPaid := totalPayments(s)
		arrears := expectedTotal - totalPaid
		if arrears < 0 {
			arrears = 0
		}
		status := "LUNAS"
		if arrears > 0 {
			status = "BELUM LUNAS"
		}
		row := []string{s.NIM, s.Name, strconv.Itoa(totalPaid), strconv.Itoa(arrears), status}
		writer.Write(row)
	}
	fmt.Println("✓ Data mahasiswa berhasil diekspor ke data_mahasiswa.csv")

	// File 2: Riwayat Pembayaran
	paymentsFile, err := os.Create("riwayat_pembayaran.csv")
	if err != nil {
		fmt.Println("✗ Gagal membuat file CSV riwayat:", err)
		return
	}
	defer paymentsFile.Close()
	writer2 := csv.NewWriter(paymentsFile)
	defer writer2.Flush()

	header2 := []string{"NIM", "Nama", "Tanggal", "Nominal", "Catatan"}
	writer2.Write(header2)

	for _, s := range store.Students {
		for _, p := range s.Payments {
			row := []string{s.NIM, s.Name, p.Date, strconv.Itoa(p.Amount), p.Note}
			writer2.Write(row)
		}
	}
	fmt.Println("✓ Riwayat pembayaran berhasil diekspor ke riwayat_pembayaran.csv")
	fmt.Println("✓ File CSV dapat dibuka dengan Microsoft Excel.")
}

func exportToTXT() {
	file, err := os.Create("laporan_kas.txt")
	if err != nil {
		fmt.Println("✗ Gagal membuat file TXT:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	writer.WriteString("=== LAPORAN KAS MAHASISWA (SIKAS) ===\n")
	writer.WriteString(fmt.Sprintf("Dibuat pada: %s\n", time.Now().Format("2006-01-02 15:04:05")))
	writer.WriteString(fmt.Sprintf("Iuran per bulan: Rp%d\n", store.MonthlyFee))
	writer.WriteString(fmt.Sprintf("Jumlah bulan wajib: %d\n", store.MonthsExpected))
	writer.WriteString(fmt.Sprintf("Total iuran wajib per mahasiswa: Rp%d\n\n", store.MonthlyFee*store.MonthsExpected))

	writer.WriteString("--- DATA MAHASISWA ---\n")
	writer.WriteString(fmt.Sprintf("%-12s %-20s %15s %15s %-8s\n", "NIM", "Nama", "Total Dibayar", "Tunggakan", "Status"))
	writer.WriteString(strings.Repeat("-", 80) + "\n")
	for _, s := range store.Students {
		totalPaid := totalPayments(s)
		expected := store.MonthlyFee * store.MonthsExpected
		arrears := expected - totalPaid
		if arrears < 0 {
			arrears = 0
		}
		status := "LUNAS"
		if arrears > 0 {
			status = "BELUM LUNAS"
		}
		writer.WriteString(fmt.Sprintf("%-12s %-20s Rp%12d Rp%12d %-8s\n", s.NIM, s.Name, totalPaid, arrears, status))
	}

	writer.WriteString("\n--- RIWAYAT PEMBAYARAN ---\n")
	writer.WriteString(fmt.Sprintf("%-12s %-20s %-12s %12s %s\n", "NIM", "Nama", "Tanggal", "Nominal", "Catatan"))
	writer.WriteString(strings.Repeat("-", 80) + "\n")
	for _, s := range store.Students {
		for _, p := range s.Payments {
			writer.WriteString(fmt.Sprintf("%-12s %-20s %-12s Rp%9d %s\n", s.NIM, s.Name, p.Date, p.Amount, p.Note))
		}
	}

	totalCash := 0
	fullyPaidCount := 0
	for _, s := range store.Students {
		totalPaid := totalPayments(s)
		totalCash += totalPaid
		if totalPaid >= store.MonthlyFee*store.MonthsExpected {
			fullyPaidCount++
		}
	}
	writer.WriteString("\n--- STATISTIK ---\n")
	writer.WriteString(fmt.Sprintf("Total saldo kas terkumpul: Rp%d\n", totalCash))
	writer.WriteString(fmt.Sprintf("Jumlah mahasiswa lunas: %d dari %d\n", fullyPaidCount, len(store.Students)))

	fmt.Println("✓ Laporan berhasil diekspor ke laporan_kas.txt")
}

// ==================== FUNGSI BANTU ====================

func totalPayments(s Student) int {
	sum := 0
	for _, p := range s.Payments {
		sum += p.Amount
	}
	return sum
}

func getArrears(s Student) int {
	expected := store.MonthlyFee * store.MonthsExpected
	paid := totalPayments(s)
	arrears := expected - paid
	if arrears < 0 {
		arrears = 0
	}
	return arrears
}

func findStudentIndexByNIM(nim string) int {
	for i, s := range store.Students {
		if s.NIM == nim {
			return i
		}
	}
	return -1
}

// ==================== PENYIMPANAN DATA INTERNAL (JSON) ====================

const dataFile = "sikas_data.json"

func saveData() {
	file, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		fmt.Println("✗ Gagal menyimpan data:", err)
		return
	}
	err = os.WriteFile(dataFile, file, 0644)
	if err != nil {
		fmt.Println("✗ Gagal menulis file:", err)
	}
}

func loadData() {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("✗ Gagal membaca file data:", err)
		return
	}
	err = json.Unmarshal(file, &store)
	if err != nil {
		fmt.Println("✗ Gagal memproses data:", err)
	}
}
