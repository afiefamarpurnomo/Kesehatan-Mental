package main

import (
	"fmt"
	"time"
)

type Penilaian struct {
	ID         int
	IDPengguna int
	Tanggal    time.Time
	Respon     []int
	TotalSkor  int
	Kategori   string
}

var penilaians []Penilaian

// Pertanyaan DASS-10 yang disesuaikan
var pertanyaanDASS10 = []string{
	"Saya merasa sulit untuk tenang setelah sesuatu yang mengganggu",
	"Saya tidak dapat merasakan perasaan positif sama sekali",
	"Saya merasa sulit untuk memulai melakukan sesuatu",
	"Saya cenderung bereaksi berlebihan terhadap situasi",
	"Saya merasa menggunakan banyak energi mental untuk khawatir",
	"Saya merasa tidak ada yang bisa saya nantikan",
	"Saya merasa gelisah",
	"Saya merasa sedih dan tertekan",
	"Saya tidak bisa antusias tentang apa pun",
	"Saya merasa hidup tidak bermakna",
}

// Skala penilaian 1-5
var skalaJawaban = []string{
	"1 - Tidak Pernah",
	"2 - Kadang-kadang",
	"3 - Sering",
	"4 - Sangat Sering",
	"5 - Selalu",
}

// untuk menghitung total skor dari array respon
func hitungTotal(respon []int) int {
	jumlah := 0
	for _, nilai := range respon {
		jumlah += nilai
	}
	return jumlah
}

// menentukan kategori berdasarkan total skor
func tentukanKategori(totalSkor int) string {
	if totalSkor >= 1 && totalSkor <= 10 {
		return "Sangat ringan"
	} else if totalSkor >= 11 && totalSkor <= 20 {
		return "Ringan"
	} else if totalSkor >= 21 && totalSkor <= 30 {
		return "Sedang"
	} else if totalSkor >= 31 && totalSkor <= 40 {
		return "Tinggi"
	} else if totalSkor >= 41 && totalSkor <= 50 {
		return "Sangat tinggi/kritis"
	} else {
		return "Skor tidak valid"
	}
}

// menampilkan pertanyaan DASS-10 dan mengambil respon
func tampilkanPertanyaan() []int {
	respon := make([]int, len(pertanyaanDASS10))

	fmt.Println("\n=== KUESIONER DASS-10 ===")
	fmt.Println("Jawab setiap pertanyaan berdasarkan pengalaman Anda dalam 2 minggu terakhir:")
	fmt.Println("\nSkala Jawaban:")
	for _, skala := range skalaJawaban {
		fmt.Println(skala)
	}
	fmt.Println()

	for i, pertanyaan := range pertanyaanDASS10 {
		fmt.Printf("%d. %s\n", i+1, pertanyaan)
		fmt.Print("Jawaban (1-5): ")

		var jawaban int
		for {
			fmt.Scan(&jawaban)
			if jawaban >= 1 && jawaban <= 5 {
				respon[i] = jawaban
				break
			} else {
				fmt.Print("Jawaban harus antara 1-5. Coba lagi: ")
			}
		}
		fmt.Println()
	}

	return respon
}

// menambah penilaian baru ke dalam slice penilaians
func tambahPenilaian(id, idPengguna int) {
	respon := tampilkanPertanyaan()
	totalSkor := hitungTotal(respon)
	kategori := tentukanKategori(totalSkor)

	penilaians = append(penilaians, Penilaian{
		ID:         id,
		IDPengguna: idPengguna,
		Tanggal:    time.Now(),
		Respon:     respon,
		TotalSkor:  totalSkor,
		Kategori:   kategori,
	})

	fmt.Printf("\n=== HASIL PENILAIAN ===\n")
	fmt.Printf("Total Skor: %d\n", totalSkor)
	fmt.Printf("Kategori: %s\n", kategori)
	fmt.Println("Penilaian berhasil ditambahkan.")

	// Rekomendasi berdasarkan kategori
	tampilkanRekomendasi(kategori)
}

// menampilkan rekomendasi berdasarkan kategori
func tampilkanRekomendasi(kategori string) {
	fmt.Printf("\n=== REKOMENDASI ===\n")
	switch kategori {
	case "Sangat ringan":
		fmt.Println(" Kondisi mental Anda sangat baik. Pertahankan pola hidup sehat!")
		fmt.Println(" Tips: Terus lakukan aktivitas positif, olahraga teratur, dan jaga pola tidur.")
		fmt.Println(" Saran: Bisa menjadi mentor atau membantu orang lain yang membutuhkan dukungan.")
	case "Ringan":
		fmt.Println(" Kondisi mental Anda cukup baik dengan sedikit tekanan normal.")
		fmt.Println(" Tips: Luangkan waktu untuk relaksasi, hobi yang menyenangkan, atau bertemu teman.")
		fmt.Println(" Saran: Coba teknik pernapasan dalam atau meditasi ringan 10-15 menit sehari.")
	case "Sedang":
		fmt.Println(" Anda mengalami tingkat stres/depresi/kecemasan yang perlu perhatian.")
		fmt.Println(" Tips: Kurangi beban kerja, atur prioritas, dan cari dukungan dari orang terdekat.")
		fmt.Println(" Saran: Pertimbangkan konseling dengan konselor atau psikolog untuk bantuan profesional.")
	case "Tinggi":
		fmt.Println(" Kondisi mental Anda memerlukan perhatian serius dan tindakan segera.")
		fmt.Println(" Tips: Segera kurangi stressor, ambil cuti jika perlu, dan hindari keputusan besar.")
		fmt.Println(" Saran: Sangat disarankan untuk segera berkonsultasi dengan psikolog atau psikiater.")
	case "Sangat tinggi/kritis":
		fmt.Println("KONDISI KRITIS - Memerlukan bantuan profesional segera!")
		fmt.Println(" Tips: Jangan menghadapi sendirian, segera hubungi tenaga kesehatan mental.")
		fmt.Println(" Saran: Hubungi hotline kesehatan mental atau kunjungi rumah sakit terdekat.")
		fmt.Println("Darurat: Jika ada pikiran menyakiti diri, segera hubungi 119 atau IGD terdekat.")
	}
}

// Untuk memperbarui penilaian berdasarkan ID
func perbaruiPenilaian(id int) bool {
	for i := range penilaians {
		if penilaians[i].ID == id {
			respon := tampilkanPertanyaan()
			totalSkor := hitungTotal(respon)
			kategori := tentukanKategori(totalSkor)

			penilaians[i].Respon = respon
			penilaians[i].TotalSkor = totalSkor
			penilaians[i].Kategori = kategori
			penilaians[i].Tanggal = time.Now()

			fmt.Printf("\n=== HASIL PENILAIAN TERBARU ===\n")
			fmt.Printf("Total Skor: %d\n", totalSkor)
			fmt.Printf("Kategori: %s\n", kategori)
			tampilkanRekomendasi(kategori)
			return true
		}
	}
	return false
}

// Untuk menghapus penilaian berdasarkan ID
func hapusPenilaian(id int) bool {
	for i, p := range penilaians {
		if p.ID == id {
			penilaians = append(penilaians[:i], penilaians[i+1:]...)
			return true
		}
	}
	return false
}

// Selection sort berdasarkan skor
func urutkanSelectionBySkor() {
	n := len(penilaians)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if penilaians[j].TotalSkor < penilaians[minIndex].TotalSkor {
				minIndex = j
			}
		}
		penilaians[i], penilaians[minIndex] = penilaians[minIndex], penilaians[i]
	}
	fmt.Println("Data telah diurutkan menggunakan Selection Sort berdasarkan skor.")
	tampilkanSemua()
}

// Insertion sort berdasarkan tanggal (terbaru ke terlama)
func urutkanInsertionByTanggal() {
	n := len(penilaians)
	for i := 1; i < n; i++ {
		kunci := penilaians[i]
		j := i - 1
		// Urutkan dari tanggal terbaru ke terlama
		for j >= 0 && penilaians[j].Tanggal.Before(kunci.Tanggal) {
			penilaians[j+1] = penilaians[j]
			j--
		}
		penilaians[j+1] = kunci
	}
	fmt.Println("Data telah diurutkan menggunakan Insertion Sort berdasarkan tanggal (terbaru ke terlama).")
	tampilkanSemua()
}

// mengurutkan penilaian berdasarkan ID (ascending) untuk binary search
func urutkanByID() {
	n := len(penilaians)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if penilaians[j].ID < penilaians[minIndex].ID {
				minIndex = j
			}
		}
		penilaians[i], penilaians[minIndex] = penilaians[minIndex], penilaians[i]
	}
}

// fungsi pencarian binary search berdasarkan ID penilaian
func cariBinary(id int) *Penilaian {
	urutkanByID()
	low, high := 0, len(penilaians)-1
	for low <= high {
		mid := (low + high) / 2
		if penilaians[mid].ID == id {
			return &penilaians[mid]
		} else if penilaians[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

// pencarian sequential search berdasarkan ID penilaian
func cariSequential(id int) *Penilaian {
	for i := range penilaians {
		if penilaians[i].ID == id {
			return &penilaians[i]
		}
	}
	return nil
}

// pencarian berdasarkan ID Pengguna
func cariByIDPengguna(idPengguna int) []Penilaian {
	var hasil []Penilaian
	for _, p := range penilaians {
		if p.IDPengguna == idPengguna {
			hasil = append(hasil, p)
		}
	}
	return hasil
}

// menampilkan detail penilaian
func tampilkanDetailPenilaian(p *Penilaian) {
	fmt.Printf("\n=== DETAIL PENILAIAN ===\n")
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("ID Pengguna: %d\n", p.IDPengguna)
	fmt.Printf("Tanggal: %s\n", p.Tanggal.Format("2006-01-02 15:04:05"))
	fmt.Printf("Total Skor: %d\n", p.TotalSkor)
	fmt.Printf("Kategori: %s\n", p.Kategori)
	tampilkanRekomendasi(p.Kategori)
}

// menampilkan laporan
func tampilkanLaporan(idPengguna int) {
	var penilaianUser []Penilaian
	for _, p := range penilaians {
		if p.IDPengguna == idPengguna {
			penilaianUser = append(penilaianUser, p)
		}
	}

	if len(penilaianUser) == 0 {
		fmt.Println("Tidak ada data penilaian untuk pengguna ini.")
		return
	}

	// Urutkan berdasarkan tanggal (terbaru ke terlama)
	for i := 1; i < len(penilaianUser); i++ {
		kunci := penilaianUser[i]
		j := i - 1
		for j >= 0 && penilaianUser[j].Tanggal.Before(kunci.Tanggal) {
			penilaianUser[j+1] = penilaianUser[j]
			j--
		}
		penilaianUser[j+1] = kunci
	}

	fmt.Printf("\n=== LAPORAN PENGGUNA ID: %d ===\n", idPengguna)
	fmt.Println("\nLima Penilaian Terakhir:")
	for i := 0; i < len(penilaianUser) && i < 5; i++ {
		fmt.Printf("ID:%d Skor:%d Kategori:%s Tgl:%s\n",
			penilaianUser[i].ID, penilaianUser[i].TotalSkor,
			penilaianUser[i].Kategori, penilaianUser[i].Tanggal.Format("2006-01-02"))
	}

	// rata-rata skor 30 hari terakhir
	batasTanggal := time.Now().AddDate(0, 0, -30)
	totalSkor, jumlah := 0, 0
	for _, p := range penilaianUser {
		if p.Tanggal.After(batasTanggal) {
			totalSkor += p.TotalSkor
			jumlah++
		}
	}
	if jumlah > 0 {
		rataRata := float64(totalSkor) / float64(jumlah)
		fmt.Printf("\nRata-rata skor 30 hari terakhir: %.2f\n", rataRata)
		fmt.Printf("Kategori rata-rata: %s\n", tentukanKategori(int(rataRata)))
	} else {
		fmt.Println("\nTidak ada data dalam 30 hari terakhir.")
	}
}

// tampilkan semua data penilaian
func tampilkanSemua() {
	if len(penilaians) == 0 {
		fmt.Println("Tidak ada data penilaian.")
		return
	}
	fmt.Println("\n=== SEMUA DATA PENILAIAN ===")
	fmt.Println("ID   IDPengguna  Skor  Kategori           Tanggal")
	fmt.Println("--------------------------------------------------------")
	for _, p := range penilaians {
		fmt.Printf("%-4d %-11d %-5d %-18s %s\n",
			p.ID, p.IDPengguna, p.TotalSkor, p.Kategori, p.Tanggal.Format("2006-01-02"))
	}
}

// tampilkan statistik umum
func tampilkanStatistik() {
	if len(penilaians) == 0 {
		fmt.Println("Tidak ada data untuk statistik.")
		return
	}

	// Hitung kategori
	kategoriBaik := 0
	kategoriHiburan := 0
	kategoriPendampingan := 0
	kategoriPerawatan := 0
	totalSkor := 0

	for _, p := range penilaians {
		totalSkor += p.TotalSkor
		switch p.Kategori {
		case "Sangat ringan":
			kategoriBaik++
		case "Ringan":
			kategoriHiburan++
		case "Sedang":
			kategoriPendampingan++
		case "Tinggi":
			kategoriPerawatan++
		case "Sangat tinggi/kritis":
			kategoriPerawatan++
		}
	}

	fmt.Printf("\n=== STATISTIK UMUM ===\n")
	fmt.Printf("Total Penilaian: %d\n", len(penilaians))
	fmt.Printf("Rata-rata Skor: %.2f\n", float64(totalSkor)/float64(len(penilaians)))
	fmt.Printf("\nDistribusi Kategori:\n")
	fmt.Printf("Sangat ringan: %d (%.1f%%)\n", kategoriBaik, float64(kategoriBaik)*100/float64(len(penilaians)))
	fmt.Printf("Ringan: %d (%.1f%%)\n", kategoriHiburan, float64(kategoriHiburan)*100/float64(len(penilaians)))
	fmt.Printf("Sedang: %d (%.1f%%)\n", kategoriPendampingan, float64(kategoriPendampingan)*100/float64(len(penilaians)))
	fmt.Printf("Tinggi & Sangat tinggi/kritis: %d (%.1f%%)\n", kategoriPerawatan, float64(kategoriPerawatan)*100/float64(len(penilaians)))
}

func main() {
	// Data contoh
	penilaians = append(penilaians,
		Penilaian{ID: 1, IDPengguna: 1001, Tanggal: time.Now().AddDate(0, 0, -10),
			Respon:    []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			TotalSkor: 20, Kategori: "Ringan"},
		Penilaian{ID: 2, IDPengguna: 1002, Tanggal: time.Now().AddDate(0, 0, -5),
			Respon:    []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
			TotalSkor: 30, Kategori: "Sedang"},
		Penilaian{ID: 3, IDPengguna: 1001, Tanggal: time.Now().AddDate(0, 0, -2),
			Respon:    []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			TotalSkor: 10, Kategori: "Sangat ringan"},
		Penilaian{ID: 4, IDPengguna: 1003, Tanggal: time.Now().AddDate(0, 0, -1),
			Respon:    []int{4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
			TotalSkor: 40, Kategori: "Tinggi"},
	)

	fmt.Println("=== SISTEM PENILAIAN KESEHATAN MENTAL DASS-10 ===")
	fmt.Println("Selamat datang di sistem penilaian kesehatan mental!")

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah Penilaian Baru")
		fmt.Println("2. Perbarui Penilaian")
		fmt.Println("3. Hapus Penilaian")
		fmt.Println("4. Cari Penilaian (Binary Search)")
		fmt.Println("5. Cari Penilaian (Sequential Search)")
		fmt.Println("6. Cari Berdasarkan ID Pengguna")
		fmt.Println("7. Urutkan Berdasarkan Skor (Selection Sort)")
		fmt.Println("8. Urutkan Berdasarkan Tanggal (Insertion Sort)")
		fmt.Println("9. Tampilkan Laporan Pengguna")
		fmt.Println("10. Tampilkan Semua Data")
		fmt.Println("11. Tampilkan Statistik Umum")
		fmt.Println("0. Keluar")

		var pilihan int
		fmt.Print("\nPilih menu (0-11): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var id, idPengguna int
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scan(&idPengguna)
			tambahPenilaian(id, idPengguna)

		case 2:
			var id int
			fmt.Print("Masukkan ID yang akan diperbarui: ")
			fmt.Scan(&id)
			if perbaruiPenilaian(id) {
				fmt.Println("Penilaian berhasil diperbarui.")
			} else {
				fmt.Println("Penilaian tidak ditemukan.")
			}

		case 3:
			var id int
			fmt.Print("Masukkan ID yang akan dihapus: ")
			fmt.Scan(&id)
			if hapusPenilaian(id) {
				fmt.Println("Penilaian berhasil dihapus.")
			} else {
				fmt.Println("Penilaian tidak ditemukan.")
			}

		case 4:
			var id int
			fmt.Print("Masukkan ID yang dicari: ")
			fmt.Scan(&id)
			if p := cariBinary(id); p != nil {
				tampilkanDetailPenilaian(p)
			} else {
				fmt.Println("Penilaian tidak ditemukan.")
			}

		case 5:
			var id int
			fmt.Print("Masukkan ID yang dicari: ")
			fmt.Scan(&id)
			if p := cariSequential(id); p != nil {
				tampilkanDetailPenilaian(p)
			} else {
				fmt.Println("Penilaian tidak ditemukan.")
			}

		case 6:
			var idPengguna int
			fmt.Print("Masukkan ID Pengguna yang dicari: ")
			fmt.Scan(&idPengguna)
			hasil := cariByIDPengguna(idPengguna)
			if len(hasil) > 0 {
				fmt.Printf("\nDitemukan %d penilaian untuk pengguna ID %d:\n", len(hasil), idPengguna)
				for _, p := range hasil {
					fmt.Printf("ID:%d Skor:%d Kategori:%s Tgl:%s\n",
						p.ID, p.TotalSkor, p.Kategori, p.Tanggal.Format("2006-01-02"))
				}
			} else {
				fmt.Println("Tidak ada penilaian untuk pengguna ini.")
			}

		case 7:
			urutkanSelectionBySkor()

		case 8:
			urutkanInsertionByTanggal()

		case 9:
			var idPengguna int
			fmt.Print("Masukkan ID Pengguna untuk laporan: ")
			fmt.Scan(&idPengguna)
			tampilkanLaporan(idPengguna)

		case 10:
			tampilkanSemua()

		case 11:
			tampilkanStatistik()

		case 0:
			fmt.Println("\nTerima kasih telah menggunakan sistem penilaian kesehatan mental!")
			fmt.Println("Jaga kesehatan mental Anda. Sampai jumpa!")
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih 0-11.")
		}
	}
}
