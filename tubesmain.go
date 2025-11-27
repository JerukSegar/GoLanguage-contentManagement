package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const nmax int = 10

type konten struct {
	kodeUnik     string
	daftarKonten string
	tglPosting   string
	kategori     string
	jamPosting   string
	views        int
}

type tabKonten [nmax]konten

func isKodeUnikSudahAda(A *tabKonten, n int, kode string) bool {
	var i int
	for i = 0; i < n; i++ {
		if A[i].kodeUnik == kode {
			return true
		}
	}
	return false
}

func bacaDaftarKonten(A *tabKonten, n *int) {
	var i int
	var jumlah int
	var scanner = bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan jumlah konten: ")
	fmt.Scan(&jumlah)

	// Membersihkan newline setelah input jumlah
	scanner.Scan()

	if jumlah > nmax {
		jumlah = nmax
		fmt.Printf("Jumlah konten dibatasi maksimal %d\n", nmax)
	}

	for i = 0; i < jumlah && i < nmax; i++ {
		fmt.Printf("\nKonten ke-%d:\n", i+1)

		// Input kode unik
		var kode string
		for {
			fmt.Print("Kode Unik (3 huruf): ")
			fmt.Scan(&kode)
			if len(kode) != 3 {
				fmt.Println("Kode unik harus terdiri dari 3 huruf!")
				continue
			}
			if isKodeUnikSudahAda(A, *n, kode) {
				fmt.Println("Kode unik sudah ada, gunakan kode lain!")
				continue
			}
			break
		}
		A[*n].kodeUnik = kode

		// Membersihkan newline setelah input kode
		scanner.Scan()

		// Input nama konten
		fmt.Print("Nama konten: ")
		if scanner.Scan() {
			A[*n].daftarKonten = scanner.Text()
		}

		// Input kategori
		fmt.Print("Kategori: ")
		if scanner.Scan() {
			A[*n].kategori = scanner.Text()
		}

		A[*n].views = 0
		*n = *n + 1
	}
}

func jadwalkanKonten(A *tabKonten, n int) {
	var idx int
	if n == 0 {
		fmt.Println("Belum ada konten. Silakan input konten terlebih dahulu.")
		return
	}

	tampilkanKonten(A, n)
	fmt.Print("\nMasukkan indeks konten yang ingin dijadwalkan (1 - ", n, "): ")
	fmt.Scan(&idx)
	idx = idx - 1

	if idx >= 0 && idx < n {
		fmt.Print("Masukkan tanggal posting (yyyy-mm-dd): ")
		fmt.Scan(&A[idx].tglPosting)
		fmt.Print("Masukkan jam posting (hh:mm): ")
		fmt.Scan(&A[idx].jamPosting)
		fmt.Println("Konten berhasil dijadwalkan!")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func tampilkanKonten(A *tabKonten, n int) {
	var i int
	fmt.Println("\n=== DAFTAR KONTEN ===")
	for i = 0; i < n; i++ {
		fmt.Printf("%d. Kode: %s | Konten: %s\n", i+1, A[i].kodeUnik, A[i].daftarKonten)
		fmt.Printf("   Kategori: %s\n", A[i].kategori)
		fmt.Printf("   Tanggal Posting: %s\n", A[i].tglPosting)
		fmt.Printf("   Jam Posting: %s\n", A[i].jamPosting)
		fmt.Printf("   Views: %d\n", A[i].views)
		fmt.Println("-----------------------")
	}
}

func hapusKonten(A *tabKonten, n *int) {
	var idx int
	var i int

	if *n == 0 {
		fmt.Println("Tidak ada konten untuk dihapus.")
		return
	}

	tampilkanKonten(A, *n)
	fmt.Print("\nMasukkan indeks konten yang ingin dihapus (1 - ", *n, "): ")
	fmt.Scan(&idx)
	idx = idx - 1

	if idx >= 0 && idx < *n {
		fmt.Printf("Konten '%s' dengan kode %s berhasil dihapus!\n", A[idx].daftarKonten, A[idx].kodeUnik)
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n = *n - 1
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func tambahViews(A *tabKonten, n int) {
	var idx int
	var jumlahViews int

	if n == 0 {
		fmt.Println("Tidak ada konten untuk input views.")
		return
	}

	tampilkanKonten(A, n)
	fmt.Print("\nMasukkan indeks konten yang ingin anda input views (1 - ", n, "): ")
	fmt.Scan(&idx)
	idx = idx - 1

	if idx >= 0 && idx < n {
		fmt.Print("Masukkan jumlah views yang ingin and input: ")
		fmt.Scan(&jumlahViews)
		A[idx].views = A[idx].views + jumlahViews
		fmt.Printf("Views berhasil diinput! Total views konten '%s': %d\n", A[idx].daftarKonten, A[idx].views)
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func urutkanBerdasarkanViews(A *tabKonten, n int) {
	var i int
	var j int
	var temp konten

	if n == 0 {
		fmt.Println("Tidak ada konten untuk diurutkan.")
		return
	}

	B := *A

	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			if B[j].views < B[j+1].views {
				temp = B[j]
				B[j] = B[j+1]
				B[j+1] = temp
			}
		}
	}

	fmt.Println("\n=== DAFTAR KONTEN TERURUT BERDASARKAN VIEWS (TERBANYAK -> TERENDAH) ===")
	for i = 0; i < n; i++ {
		fmt.Printf("%d. Kode: %s | Konten: %s\n", i+1, B[i].kodeUnik, B[i].daftarKonten)
		fmt.Printf("   Kategori: %s\n", B[i].kategori)
		fmt.Printf("   Tanggal Posting: %s\n", B[i].tglPosting)
		fmt.Printf("   Jam Posting: %s\n", B[i].jamPosting)
		fmt.Printf("   Views: %d\n", B[i].views)
		fmt.Println("-----------------------")
	}
}

func cariBerdasarkanKode(A *tabKonten, n int) {
	var kode string
	var i int
	var ditemukan bool

	if n == 0 {
		fmt.Println("Tidak ada konten untuk dicari.")
		return
	}

	fmt.Print("Masukkan kode unik konten: ")
	fmt.Scan(&kode)

	ditemukan = false
	for i = 0; i < n; i++ {
		if A[i].kodeUnik == kode {
			fmt.Println("\n=== KONTEN DITEMUKAN ===")
			fmt.Printf("Kode: %s | Konten: %s\n", A[i].kodeUnik, A[i].daftarKonten)
			fmt.Printf("Kategori: %s\n", A[i].kategori)
			fmt.Printf("Tanggal Posting: %s\n", A[i].tglPosting)
			fmt.Printf("Jam Posting: %s\n", A[i].jamPosting)
			fmt.Printf("Views: %d\n", A[i].views)
			ditemukan = true
			break
		}
	}

	if !ditemukan {
		fmt.Println("Konten dengan kode", kode, "tidak ditemukan.")
	}
}

func ideKonten() {
	var daftarIde = [6]string{
		"Tutorial Photoshop untuk pemula",
		"Review smartphone terbaru",
		"Tutorial membuat makeup natural",
		"Vlog liburan ke Bali",
		"Tips fotografi dengan HP",
		"Panduan editing video dasar",
	}

	var input string
	var ditemukan bool
	var i int

	fmt.Println("\n=== PENCARIAN IDE KONTEN ===")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Kata kunci: tutorial, review, vlog, tips, panduan")
		fmt.Print("Masukkan kata kunci (ketik 'selesai' untuk keluar): ")

		if scanner.Scan() {
			input = scanner.Text()
		}

		if input == "selesai" {
			break
		}

		fmt.Println("\nHasil pencarian:")
		ditemukan = false

		for i = 0; i < len(daftarIde); i++ {
			if strings.Contains(strings.ToLower(daftarIde[i]), strings.ToLower(input)) {
				fmt.Printf("- %s\n", daftarIde[i])
				ditemukan = true
			}
		}

		if !ditemukan {
			fmt.Println("Waduh, belum ada ide yang cocok dengan kata kunci tersebut.")
		}
	}

	fmt.Println("\nTerima kasih telah menggunakan fitur pencarian ide!")
}

func main() {
	var data tabKonten
	var n int
	var pilihan int

	fmt.Println("=== APLIKASI MANAJEMEN KONTEN ===")

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Input Daftar Konten")
		fmt.Println("2. Jadwalkan Konten")
		fmt.Println("3. Tampilkan Semua Konten")
		fmt.Println("4. Hapus Konten")
		fmt.Println("5. Input Views")
		fmt.Println("6. Urutkan Berdasarkan Views")
		fmt.Println("7. Cari Konten Berdasarkan Kode")
		fmt.Println("8. Cari Ide Konten")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu (1-9): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			bacaDaftarKonten(&data, &n)
		case 2:
			jadwalkanKonten(&data, n)
		case 3:
			if n > 0 {
				tampilkanKonten(&data, n)
			} else {
				fmt.Println("Belum ada konten untuk ditampilkan.")
			}
		case 4:
			hapusKonten(&data, &n)
		case 5:
			tambahViews(&data, n)
		case 6:
			urutkanBerdasarkanViews(&data, n)
		case 7:
			cariBerdasarkanKode(&data, n)
		case 8:
			ideKonten()
		case 9:
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih 1-9.")
		}
	}
}
