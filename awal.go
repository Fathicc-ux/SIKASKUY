package main

import "fmt"

type mahasiswa struct {
	nama, nim string
}
type transkasi struct {
	nim     string
	nominal int
	tanggal int
	bulan   int
	tahun   int
	status  string
}



func main() {
	fmt.Print("== SIKAS ==")
	fmt.Println()

	var user, pass string
	var inuser, inpass string

	for {
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("0. Keluar")

		fmt.Println()

		var menu int
		var admin int

		fmt.Print("Pilih menu: ")
		fmt.Scan(&menu)

		switch menu {
		case 1:
			fmt.Println("== Registrasi ==")
			fmt.Print("User: ")
			fmt.Scan(&user)
			fmt.Print("Password: ")
			fmt.Scan(&pass)
			fmt.Println("== Registrasi Berhasil ==")
			fmt.Println()

		case 2:
			fmt.Println("== Login ==")

			fmt.Print("Masukan user: ")
			fmt.Scan(&inuser)
			fmt.Print("Masukan Password: ")
			fmt.Scan(&inpass)

			if inuser == user && inpass == pass {
				fmt.Println("== Login berhasil ==")
				fmt.Println()
			}
			for {
				fmt.Println("== ADMIN ==")
				fmt.Println("1. Mahasiswa")
				fmt.Println(" 1.1 Tambah Mahasiswa")
				fmt.Println(" 1.2 Hapus Mahasiswa")
				fmt.Println(" 1.3 Ubah Mahasiswa")

				fmt.Print("Pilih menu: ")
				fmt.Scan(&admin)
			}
		case 0:
			break
		}
	}
}

func tambahMahasiswa() {
	var mhs mahasiswa
	fmt.Print("Nama: ")
	fmt.Scan(&mhs.nama)
	fmt.Print("NIM: ")
	fmt.Scan(&mhs.nim)
	fmt.Println("== Mahasiswa Berhasil Ditambahkan ==")
}
