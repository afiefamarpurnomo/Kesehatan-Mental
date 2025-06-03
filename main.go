package main

import (
	"fmt"
	"time"
)

// nyimpen data penilaian
type penilaian struct {
	id         int
	idpengguna int
	tanggal    time.Time
	jawaban    []int
	skor       int
	level      string
}

var datapenilaian []penilaian

// pertanyaan
var soaldass = []string{
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

// opsi jawaban
var pilihanjawabann = []string{
	"1 - Tidak Pernah",
	"2 - Kadang-kadang",
	"3 - Sering",
	"4 - Sangat Sering",
	"5 - Selalu",
}

func hitungskor(jawaban []int) int {
	total := 0
	for _, nilai := range jawaban {
		total += nilai
	}
	return total
}

func ceklevel(skor int) string {
	switch {
	case skor <= 10:
		return "Normal"
	case skor <= 20:
		return "Ringan"
	case skor <= 30:
		return "Sedang"
	case skor <= 40:
		return "Berat"
	default:
		return "Sangat Berat"
	}
}

func isikuesioner() []int {
	jawaban := make([]int, len(soaldass))

	fmt.Println("\n=== KUESIONER DASS-10 ===")
	fmt.Println("Jawab berdasarkan kondisi 2 minggu terakhir:")
	fmt.Println()
	for _, pilihan := range pilihanjawabann {
		fmt.Println(pilihan)
	}
	fmt.Println()

	for i, soal := range soaldass {
		fmt.Printf("%d. %s\n", i+1, soal)
		fmt.Print("Pilih (1-5): ")

		var input int
		for {
			fmt.Scan(&input)
			if input >= 1 && input <= 5 {
				jawaban[i] = input
				break
			} else {
				fmt.Print("Input salah, coba lagi (1-5): ")
			}
		}
		fmt.Println()
	}

	return jawaban
}

func tambahdata(id, userid int) {
	jawaban := isikuesioner()
	skortotal := hitungskor(jawaban)
	levelstress := ceklevel(skortotal)

	//tambah ke slice
	datapenilaian = append(datapenilaian, penilaian{
		id:         id,
		idpengguna: userid,
		tanggal:    time.Now(),
		jawaban:    jawaban,
		skor:       skortotal,
		level:      levelstress,
	})

	fmt.Printf("\n=== HASIL ===\n")
	fmt.Printf("Skor: %d\n", skortotal)
	fmt.Printf("Level: %s\n", levelstress)

	kasihsaran(levelstress)
}

func kasihsaran(level string) {
	fmt.Printf("\n=== SARAN ===\n")
	switch level {
	case "Normal":
		fmt.Println("Kondisi mental baik, pertahankan!")
		fmt.Println("  Tips: Tetap olahraga dan jaga pola tidur")
	case "Ringan":
		fmt.Println("Ada sedikit tekanan, tapi masih normal")
		fmt.Println("  Tips: Coba relaksasi atau hobi yang disukai")
	case "Sedang":
		fmt.Println("Perlu perhatian lebih")
		fmt.Println("  Tips: Kurangi beban, cari dukungan keluarga/teman")
		fmt.Println("  Saran: Pertimbangkan konseling")
	case "Berat":
		fmt.Println("Kondisi serius, butuh tindakan")
		fmt.Println("  Tips: Kurangi stress, ambil cuti jika perlu")
		fmt.Println("  Saran: Segera konsultasi psikolog/psikiater")
	case "Sangat Berat":
		fmt.Println("KONDISI KRITIS!")
		fmt.Println("  Jangan sendirian, hubungi tenaga medis segera")
		fmt.Println("  Darurat: Pikiran menyakiti diri -> hubungi 119")
	}
}

// update data
func updatedata(id int) bool {
	for i := range datapenilaian {
		if datapenilaian[i].id == id {
			jawaban := isikuesioner()
			skor := hitungskor(jawaban)
			level := ceklevel(skor)

			datapenilaian[i].jawaban = jawaban
			datapenilaian[i].skor = skor
			datapenilaian[i].level = level
			datapenilaian[i].tanggal = time.Now()

			fmt.Printf("\n=== HASIL UPDATE ===\n")
			fmt.Printf("Skor: %d\n", skor)
			fmt.Printf("Level: %s\n", level)
			kasihsaran(level)
			return true
		}
	}
	return false
}

func hapusdata(id int) bool {
	for i, data := range datapenilaian {
		if data.id == id {
			datapenilaian = append(datapenilaian[:i], datapenilaian[i+1:]...)
			return true
		}
	}
	return false
}

// selection sort berdasarkan skor
func urutberdasarkanskor() {
	n := len(datapenilaian)
	for i := 0; i < n-1; i++ {
		minidx := i
		for j := i + 1; j < n; j++ {
			if datapenilaian[j].skor < datapenilaian[minidx].skor {
				minidx = j
			}
		}
		datapenilaian[i], datapenilaian[minidx] = datapenilaian[minidx], datapenilaian[i]
	}
	fmt.Println("Data sudah diurutkan berdasarkan skor (rendah ke tinggi)")
	tampilkansemua()
}

// insertion sort berdasarkan tanggal
func urutberdasarkantanggal() {
	n := len(datapenilaian)
	for i := 1; i < n; i++ {
		key := datapenilaian[i]
		j := i - 1

		for j >= 0 && datapenilaian[j].tanggal.Before(key.tanggal) {
			datapenilaian[j+1] = datapenilaian[j]
			j--
		}
		datapenilaian[j+1] = key
	}
	fmt.Println("Data sudah diurutkan berdasarkan tanggal (terbaru dulu)")
	tampilkansemua()
}

func urutkan() {
	n := len(datapenilaian)
	for i := 0; i < n-1; i++ {
		minidx := i
		for j := i + 1; j < n; j++ {
			if datapenilaian[j].id < datapenilaian[minidx].id {
				minidx = j
			}
		}
		datapenilaian[i], datapenilaian[minidx] = datapenilaian[minidx], datapenilaian[i]
	}
}

func caribinary(id int) *penilaian {
	urutkan()

	low := 0
	high := len(datapenilaian) - 1

	for low <= high {
		mid := (low + high) / 2
		if datapenilaian[mid].id == id {
			return &datapenilaian[mid]
		} else if datapenilaian[mid].id < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

func carisequential(id int) *penilaian {
	for i := range datapenilaian {
		if datapenilaian[i].id == id {
			return &datapenilaian[i]
		}
	}
	return nil
}

// cari semua data
func caridatauser(userid int) []penilaian {
	var hasil []penilaian
	for _, data := range datapenilaian {
		if data.idpengguna == userid {
			hasil = append(hasil, data)
		}
	}
	return hasil
}

func tampilkandetail(p *penilaian) {
	fmt.Printf("\n=== DETAIL PENILAIAN ===\n")
	fmt.Printf("ID: %d\n", p.id)
	fmt.Printf("User ID: %d\n", p.idpengguna)
	fmt.Printf("Tanggal: %s\n", p.tanggal.Format("02/01/2006 15:04"))
	fmt.Printf("Skor: %d\n", p.skor)
	fmt.Printf("Level: %s\n", p.level)
	kasihsaran(p.level)
}

// laporan untuk user
func laporanuser(userid int) {
	userdata := caridatauser(userid)

	if len(userdata) == 0 {
		fmt.Println("Ga ada data untuk user ini")
		return
	}

	// urutkan tanggal
	for i := 1; i < len(userdata); i++ {
		key := userdata[i]
		j := i - 1
		for j >= 0 && userdata[j].tanggal.Before(key.tanggal) {
			userdata[j+1] = userdata[j]
			j--
		}
		userdata[j+1] = key
	}

	fmt.Printf("\n=== LAPORAN USER %d ===\n", userid)
	fmt.Println("\n5 Penilaian Terakhir:")

	limit := len(userdata)
	if limit > 5 {
		limit = 5
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("ID:%d | Skor:%d | %s | %s\n",
			userdata[i].id, userdata[i].skor,
			userdata[i].level, userdata[i].tanggal.Format("02/01/06"))
	}

	// hitung rata2 30 hari
	cutoff := time.Now().AddDate(0, 0, -30)
	totalskor := 0
	count := 0

	for _, data := range userdata {
		if data.tanggal.After(cutoff) {
			totalskor += data.skor
			count++
		}
	}

	if count > 0 {
		avg := float64(totalskor) / float64(count)
		fmt.Printf("\nRata-rata 30 hari terakhir: %.1f\n", avg)
		fmt.Printf("Level rata-rata: %s\n", ceklevel(int(avg)))
	} else {
		fmt.Println("\nGa ada data 30 hari terakhir")
	}
}

func tampilkansemua() {
	if len(datapenilaian) == 0 {
		fmt.Println("Belum ada data")
		return
	}

	fmt.Println("\n=== SEMUA DATA ===")
	fmt.Printf("%-4s %-8s %-5s %-12s %s\n", "ID", "UserID", "Skor", "Level", "Tanggal")
	fmt.Println("------------------------------------------")

	for _, data := range datapenilaian {
		fmt.Printf("%-4d %-8d %-5d %-12s %s\n",
			data.id, data.idpengguna, data.skor,
			data.level, data.tanggal.Format("02/01/06"))
	}
}

func tampilkanstats() {
	if len(datapenilaian) == 0 {
		fmt.Println("Belum ada data untuk statistik")
		return
	}

	// hitung distribusi
	levelcount := make(map[string]int)
	totalskor := 0

	for _, data := range datapenilaian {
		levelcount[data.level]++
		totalskor += data.skor
	}

	fmt.Printf("\n=== STATISTIK ===\n")
	fmt.Printf("Total data: %d\n", len(datapenilaian))
	fmt.Printf("Rata-rata skor: %.1f\n", float64(totalskor)/float64(len(datapenilaian)))

	fmt.Printf("\nDistribusi Level:\n")
	for level, count := range levelcount {
		persen := float64(count) * 100 / float64(len(datapenilaian))
		fmt.Printf("- %s: %d (%.1f%%)\n", level, count, persen)
	}
}

func main() {
	// data dummy
	datapenilaian = append(datapenilaian,
		penilaian{
			id: 1, idpengguna: 1001,
			tanggal: time.Now().AddDate(0, 0, -10),
			jawaban: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			skor:    20, level: "Ringan",
		},
		penilaian{
			id: 2, idpengguna: 1002,
			tanggal: time.Now().AddDate(0, 0, -5),
			jawaban: []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
			skor:    30, level: "Sedang",
		},
		penilaian{
			id: 3, idpengguna: 1001,
			tanggal: time.Now().AddDate(0, 0, -2),
			jawaban: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			skor:    10, level: "Normal",
		},
	)

	fmt.Println("=================== SISTEM PENILAIAN KESEHATAN MENTAL ===================")
	fmt.Println("Selamat datang!")

	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Tambah Penilaian")
		fmt.Println("2. Update Penilaian")
		fmt.Println("3. Hapus Penilaian")
		fmt.Println("4. Cari (Binary Search)")
		fmt.Println("5. Cari (Sequential Search)")
		fmt.Println("6. Cari by User ID")
		fmt.Println("7. Urutkan by Skor")
		fmt.Println("8. Urutkan by Tanggal")
		fmt.Println("9. Laporan User")
		fmt.Println("10. Lihat Semua")
		fmt.Println("11. Statistik")
		fmt.Println("0. Exit")

		var pilihan int
		fmt.Print("\nPilih (0-11): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var id, userid int
			fmt.Print("ID: ")
			fmt.Scan(&id)
			fmt.Print("User ID: ")
			fmt.Scan(&userid)
			tambahdata(id, userid)

		case 2:
			var id int
			fmt.Print("ID yang mau diupdate: ")
			fmt.Scan(&id)
			if updatedata(id) {
				fmt.Println("Berhasil diupdate")
			} else {
				fmt.Println("ID tidak ditemukan")
			}

		case 3:
			var id int
			fmt.Print("ID yang mau dihapus: ")
			fmt.Scan(&id)
			if hapusdata(id) {
				fmt.Println("Berhasil dihapus")
			} else {
				fmt.Println("ID tidak ada")
			}

		case 4:
			var id int
			fmt.Print("Cari ID: ")
			fmt.Scan(&id)
			if hasil := caribinary(id); hasil != nil {
				tampilkandetail(hasil)
			} else {
				fmt.Println("Tidak ketemu")
			}

		case 5:
			var id int
			fmt.Print("Cari ID: ")
			fmt.Scan(&id)
			if hasil := carisequential(id); hasil != nil {
				tampilkandetail(hasil)
			} else {
				fmt.Println("Tidak ada")
			}

		case 6:
			var userid int
			fmt.Print("User ID: ")
			fmt.Scan(&userid)
			hasil := caridatauser(userid)
			if len(hasil) > 0 {
				fmt.Printf("\nKetemu %d data untuk user %d:\n", len(hasil), userid)
				for _, data := range hasil {
					fmt.Printf("ID:%d | Skor:%d | %s | %s\n",
						data.id, data.skor, data.level, data.tanggal.Format("02/01/06"))
				}
			} else {
				fmt.Println("User tidak punya data")
			}

		case 7:
			urutberdasarkanskor()

		case 8:
			urutberdasarkantanggal()

		case 9:
			var userid int
			fmt.Print("User ID untuk laporan: ")
			fmt.Scan(&userid)
			laporanuser(userid)

		case 10:
			tampilkansemua()

		case 11:
			tampilkanstats()

		case 0:
			fmt.Println("\nTerima kasih!")
			fmt.Println("Jaga kesehatan mental ya!")
			return

		default:
			fmt.Println("Pilihan salah, coba lagi")
		}
	}
}
