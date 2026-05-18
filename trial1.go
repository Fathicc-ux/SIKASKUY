package main
import "fmt"

type mahasiswa struct {
	NIM	 string
	Nama string
	tunggakan int
	statuspembayaraan bool
	riwayat [100]transaksi
	jumlahbayar int
	
}

type Database struct {
	mahasiswa [100]mahasiswa
	mahasiswaCount int

}

type transaksi struct {
	tanggal string
	nominal int
}

var DB Database

func tambah(A *Database) {
	if A.mahasiswaCount < 100 {
	
		var mhs mahasiswa

		fmt.Print("Nama: ")
		fmt.Scan(&mhs.Nama)

		fmt.Print("NIM: ")
		fmt.Scan(&mhs.NIM)

		mhs.tunggakan = 0
		mhs.statuspembayaraan = false
		A.mahasiswa[A.mahasiswaCount] = mhs
		A.mahasiswaCount++
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
	}else {
		fmt.Print("Nama: ")
		fmt.Scan(&A.mahasiswa[indeks].Nama)

		fmt.Print("NIM: ")
		fmt.Scan(&A.mahasiswa[indeks].NIM)
		
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
		
	}else {
		for i = indeks; i < A.mahasiswaCount-1; i++ {
			A.mahasiswa[i] = A.mahasiswa[i+1]
		}
		A.mahasiswaCount--
		fmt.Println("== Mahasiswa Berhasil Dihapus ==")
	}
}

// func cariMahasiswa(A *Database, nim string) int {
// 	var i int
// 	var indeks int
// 	indeks = -1
	
// 	for i = 0; i < A.mahasiswaCount; i++ {
// 		if A.mahasiswa[i].NIM == nim {
// 			indeks = i
		
// 			i = i + 1
// 		}

// 	}
// 	return indeks
// }

func squential(A *Database, nim string, indeks *int){
	var i int
	*indeks = -1

	for i = 0; i < A.mahasiswaCount; i++ {
		if A.mahasiswa[i].NIM == nim {
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
		fmt.Println("NIM  :", A.mahasiswa[indeks].NIM)
		fmt.Println("Nama :", A.mahasiswa[indeks].Nama)
	}
}

func cari_binary(A *Database) {
	var nim string
	var indeks int

	fmt.Print("Masukkan NIM Mahasiswa yang akan dicari: ")
	fmt.Scan(&nim)
	sort(A)
	binary(A, nim, &indeks)
	if indeks == -1 {
		fmt.Println("== Mahasiswa Tidak Ditemukan ==")
	} else {
		fmt.Println("NIM  :", A.mahasiswa[indeks].NIM)
		fmt.Println("Nama :", A.mahasiswa[indeks].Nama)
	}	
}

func binary(A *Database, nim string, indeks *int) {
	var kanan, kiri, tengah int
	

	*indeks = -1
	kiri = 0
	kanan = A.mahasiswaCount - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		
		if A.mahasiswa[tengah].NIM == nim {
			*indeks = tengah
			kiri = kanan + 1
		} else if A.mahasiswa[tengah].NIM < nim {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
}


func sort(A *Database) {
	var i, j int
	var sementara mahasiswa

	for i = 1; i < A.mahasiswaCount-1; i++ {
		sementara = A.mahasiswa[i]
		j = i - 1
		
		for j >= 0 && A.mahasiswa[j].NIM > sementara.NIM {
			A.mahasiswa[j+1] = A.mahasiswa[j]
			j = j - 1
		}
		A.mahasiswa[j+1] = sementara
	}
}

func tampil(A *Database) {

	var i int

	for i = 0; i < A.mahasiswaCount; i++ {
		fmt.Println("-----")
		fmt.Println("Nama :", A.mahasiswa[i].Nama)
		fmt.Println("NIM  :", A.mahasiswa[i].NIM)
		fmt.Println("-----")
	}
}

func bayar(A *Database) {
	var nim string
	var indeks int

	fmt.Print("Masukkan NIM Mahasiswa yang akan melakukan pembayaran: ")
	fmt.Scan(&nim)
	squential(A, nim, &indeks)

	if indeks == -1 {
		fmt.Println("== Mahasiswa Tidak Ditemukan ==")
	}else{
		fmt.Print("Masukan Tanggal Pembayaran (DD/MM/YYYY): "
		fmt.Scan(&transaksi.tanggal)

		
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
		fmt.Println("5. Cari Sequential")
		fmt.Println("6. Cari Binary")
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

			cari_squential(&DB)

		} else if pilih == 6 {

			cari_binary(&DB)

		} else if pilih == 0 {

			fmt.Println("Selesai")

		} else {

			fmt.Println("Menu salah")
		}
	}
}
	


	