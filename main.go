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
}

var penilaians []Penilaian

// Untuk menghitung total skor dari array respon
func hitungTotal(respon []int) int {
	jumlah := 0
	for _, nilai := range respon {
		jumlah += nilai
	}
	return jumlah
}

// Menambah penilaian baru ke dalam slice penilaians
func tambahPenilaian(id, idPengguna int, respon []int) {
	penilaians = append(penilaians, Penilaian{
		ID:         id,
		IDPengguna: idPengguna,
		Tanggal:    time.Now(),
		Respon:     respon,
		TotalSkor:  hitungTotal(respon),
	})
	fmt.Println("Penilaian berhasil ditambahkan.")
}

// Untuk memperbarui penilaian berdasarkan ID
func perbaruiPenilaian(id int, respon []int) bool {
	for i := range penilaians {
		if penilaians[i].ID == id {
			penilaians[i].Respon = respon
			penilaians[i].TotalSkor = hitungTotal(respon)
			penilaians[i].Tanggal = time.Now()
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

// Selection sort
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
}

// Insertion sort
func urutkanInsertionBySkor() {
	n := len(penilaians)
	for i := 1; i < n; i++ {
		kunci := penilaians[i]
		j := i - 1
		for j >= 0 && penilaians[j].TotalSkor > kunci.TotalSkor {
			penilaians[j+1] = penilaians[j]
			j--
		}
		penilaians[j+1] = kunci
	}
	fmt.Println("Data telah diurutkan menggunakan Insertion Sort berdasarkan skor.")
}

// Binary search untuk mengurutkan penilaian berdasarkan ID (ascending)
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

// Fungsi pencarian binary search berdasarkan ID penilaian
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

// Pencarian sequential search berdasarkan ID penilaian
func cariSequential(id int) *Penilaian {
	for i := range penilaians {
		if penilaians[i].ID == id {
			return &penilaians[i]
		}
	}
	return nil
}

// Menampilkan laporan
func tampilkanLaporan(idPengguna int) {
	var penilaianUser []Penilaian
	for _, p := range penilaians {
		if p.IDPengguna == idPengguna {
			penilaianUser = append(penilaianUser, p)
		}
	}

	for i := 1; i < len(penilaianUser); i++ {
		kunci := penilaianUser[i]
		j := i - 1
		for j >= 0 && penilaianUser[j].Tanggal.Before(kunci.Tanggal) {
			penilaianUser[j+1] = penilaianUser[j]
			j--
		}
		penilaianUser[j+1] = kunci
	}

	fmt.Println("\nLima Penilaian Terakhir:")
	for i := 0; i < len(penilaianUser) && i < 5; i++ {
		fmt.Printf("ID:%d Skor:%d Tgl:%s\n",
			penilaianUser[i].ID, penilaianUser[i].TotalSkor, penilaianUser[i].Tanggal.Format("2006-01-02"))
	}

	// Hitung rata-rata skor 30 hari terakhir
	batasTanggal := time.Now().AddDate(0, 0, -30)
	totalSkor, jumlah := 0, 0
	for _, p := range penilaianUser {
		if p.Tanggal.After(batasTanggal) {
			totalSkor += p.TotalSkor
			jumlah++
		}
	}
	if jumlah > 0 {
		fmt.Printf("Rata-rata skor 30 hari terakhir: %.2f\n", float64(totalSkor)/float64(jumlah))
	} else {
		fmt.Println("Tidak ada data dalam 30 hari terakhir.")
	}
}

// Menampilkan semua data penilaian
func tampilkanSemua() {
	if len(penilaians) == 0 {
		fmt.Println("Tidak ada data penilaian.")
		return
	}
	fmt.Println("ID   IDPengguna  Skor  Tanggal")
	for _, p := range penilaians {
		fmt.Printf("%-4d %-11d %-5d %s\n",
			p.ID, p.IDPengguna, p.TotalSkor, p.Tanggal.Format("2006-01-02"))
	}
}

func main() {

	penilaians = append(penilaians,
		Penilaian{ID: 1, IDPengguna: 1001, Tanggal: time.Now().AddDate(0, 0, -10), Respon: []int{4, 3, 5, 2, 4}, TotalSkor: 18},
		Penilaian{ID: 2, IDPengguna: 1002, Tanggal: time.Now().AddDate(0, 0, -5), Respon: []int{5, 5, 5, 4, 5}, TotalSkor: 24},
		Penilaian{ID: 3, IDPengguna: 1001, Tanggal: time.Now().AddDate(0, 0, -2), Respon: []int{3, 2, 4, 3, 3}, TotalSkor: 15},
	)

	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Tambah Penilaian")
		fmt.Println("2. Perbarui Penilaian")
		fmt.Println("3. Hapus Penilaian")
		fmt.Println("4. Cari Penilaian (Binary Search)")
		fmt.Println("5. Cari Penilaian (Sequential Search)")
		fmt.Println("6. Urutkan Penilaian (Selection Sort berdasarkan skor)")
		fmt.Println("7. Tampilkan Laporan")
		fmt.Println("8. Tampilkan Semua Data")
		fmt.Println("9. Urutkan Penilaian (Insertion Sort berdasarkan skor)")
		fmt.Println("0. Keluar")

		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var id, idPengguna int
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scan(&idPengguna)
			respon := make([]int, 5)
			for i := 0; i < 5; i++ {
				fmt.Printf("Masukkan skor ke-%d (1-5): ", i+1)
				fmt.Scan(&respon[i])
			}
			tambahPenilaian(id, idPengguna, respon)
		case 2:
			var id int
			fmt.Print("Masukkan ID Penilaian yang akan diperbarui: ")
			fmt.Scan(&id)
			respon := make([]int, 5)
			for i := 0; i < 5; i++ {
				fmt.Printf("Masukkan skor ke-%d (1-5): ", i+1)
				fmt.Scan(&respon[i])
			}
			if perbaruiPenilaian(id, respon) {
				fmt.Println("Penilaian berhasil diperbarui.")
			} else {
				fmt.Println("Penilaian tidak ditemukan.")
			}
		case 3:
			var id int
			fmt.Print("Masukkan ID Penilaian yang akan dihapus: ")
			fmt.Scan(&id)
			if hapusPenilaian(id) {
				fmt.Println("Penilaian berhasil dihapus.")
			} else {
				fmt.Println("Penilaian tidak ditemukan.")
			}
		case 4:
			var id int
			fmt.Print("Masukkan ID Penilaian yang dicari: ")
			fmt.Scan(&id)
			if p := cariBinary(id); p != nil {
				fmt.Printf("Ditemukan: ID %d, ID Pengguna %d, Skor %d\n", p.ID, p.IDPengguna, p.TotalSkor)
			} else {
				fmt.Println("Penilaian tidak ditemukan.")
			}
		case 5:
			var id int
			fmt.Print("Masukkan ID Penilaian yang dicari: ")
			fmt.Scan(&id)
			if p := cariSequential(id); p != nil {
				fmt.Printf("Ditemukan: ID %d, ID Pengguna %d, Skor %d\n", p.ID, p.IDPengguna, p.TotalSkor)
			} else {
				fmt.Println("Penilaian tidak ditemukan.")
			}
		case 6:
			urutkanSelectionBySkor()
			tampilkanSemua()
		case 7:
			var idPengguna int
			fmt.Print("Masukkan ID Pengguna untuk laporan: ")
			fmt.Scan(&idPengguna)
			tampilkanLaporan(idPengguna)
		case 8:
			tampilkanSemua()
		case 9:
			urutkanInsertionBySkor()
			tampilkanSemua()
		case 0:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
