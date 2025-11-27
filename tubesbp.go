package main 

import "fmt"

const nmax int = 10

type konten struct {
	daftarKonten string
	tglPosting string
	kategori string
	jamPosting string
	jmlInteraksi int
}

type tabKonten [nmax]konten

func bacaDaftarKonten(A *tabKonten, n *int) {
	var i int
	
	fmt.Print("jumlah konten: ")
	fmt.Scan(n)
	for i = 0; i < *n && i < nmax; i++ {
		fmt.Println("masukan daftarKonten, dan kategori (jangan di spasi)")
		fmt.Scan(&A[i].daftarKonten, &A[i].tglPosting, &A[i].kategori)
	}
}

func jadwalkanKonten(A *tabKonten, n int) {
	var idx int
	fmt.Print("\nMasukkan indeks konten yang ingin dijadwalkan (0 - ", n-1, "): ")
	fmt.Scan(&idx)

	if idx >= 0 && idx < n {
		fmt.Print("Masukkan tanggal posting (yyyy-mm-dd): ")
		fmt.Scan(&A[idx].tglPosting)
		fmt.Print("Masukkan jam posting (hh:mm): ")
		fmt.Scan(&A[idx].jamPosting)
		fmt.Println("Konten berhasil dijadwalkan.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func ideKonten() {
    type daftarIde = []string{
        "Tutorial Photoshop untuk pemula",
        "Review smartphone terbaru",
        "Tutorial membuat makeup natural",
        "Vlog liburan ke Bali",
        "Tips fotografi dengan HP",
        "Panduan editing video dasar",
    }

    for {
        fmt.Print("\nCari ide untuk konten kamu yuk: ")
        fmt.Println("\n(ketik 'selesai' bila sudah selesai)")
        var input string
        fmt.Scanln(&input)

        if input == "selesai" {
            break
        }

        fmt.Println("\nHasil pencarian:")
        ditemukan := false
        
        for _, ide := range daftarIde {
            cocok := true
            for j := 0; j < len(input); j++ {
                if j >= len(ide) || input[j] != ide[j] {
                    cocok = false
                    break
                }
            }
            
            if cocok {
                fmt.Printf(ide)
                ditemukan = true
            }
        }

        if !ditemukan {
            fmt.Println("\nWaduh, belum ada ide yang cocok")
        }
    }

    fmt.Println("\nTerima kasih")
}

func main() {
	var data tabKonten
	var n, i int

	bacaDaftarKonten(&data, &n)
	jadwalkanKonten(&data, n)
	
	for i = 0; i < n; i++ {
		fmt.Printf("Konten: %s, Tanggal: %s, Kategori: %s, Jam: %s, Interaksi: %d\n", 
			data[i].daftarKonten, data[i].tglPosting, data[i].kategori, data[i].jamPosting, data[i].jmlInteraksi)
	}
}
