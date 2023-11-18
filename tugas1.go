package main

import "fmt"

type soal struct {
	soalU string
	soal1 string
	soal2 string
	soal3 string
	soal4 string
}

var nilai_benar []int
var nilai_salah []int

var x int
var y int
var z int
var q int

var nama string

func main() {

	fmt.Println("input nama:")
	fmt.Scanln(&nama)

	var soal_pertama = []soal{
		{
			soalU: "siapa sih dosen terbaik?",
			soal1: "1) yoga",
			soal2: "2) widyasa",
			soal3: "3) PAK HELMY",
			soal4: "4) mas lana",
		},
	}
	var soal_kedua = []soal{
		{
			soalU: "siapakah yang membuat GTA V??",
			soal1: "1) Fromsofware",
			soal2: "2) Activision",
			soal3: "3) PAK HELMY",
			soal4: "4) Rockstar",
		},
	}
	var soal_ketiga = []soal{
		{
			soalU: "siapakah pengembang mincraft???",
			soal1: "1) valve",
			soal2: "2) mojang",
			soal3: "3) PAK HELMY",
			soal4: "4) epic",
		},
	}
	var soal_keempat = []soal{
		{
			soalU: "genre permainan 'fornite' adalah????",
			soal1: "1) battle royale",
			soal2: "2) MOBA",
			soal3: "3) PAK HELMY",
			soal4: "4) survival horror",
		},
	}

	for _, s := range soal_pertama {
		fmt.Println(s.soalU)
		fmt.Println(s.soal1)
		fmt.Println(s.soal2)
		fmt.Println(s.soal3)
		fmt.Println(s.soal4)
	}
	fmt.Println("==========")
	fmt.Println("jawaban (1,2,3,4)")
	fmt.Println("masukkan jawaban:")
	fmt.Scanln(&x)

	for _, s := range soal_kedua {
		fmt.Println(s.soalU)
		fmt.Println(s.soal1)
		fmt.Println(s.soal2)
		fmt.Println(s.soal3)
		fmt.Println(s.soal4)
	}
	fmt.Println("==========")
	fmt.Println("jawaban (1,2,3,4)")
	fmt.Println("masukkan jawaban:")
	fmt.Scanln(&y)

	for _, s := range soal_ketiga {
		fmt.Println(s.soalU)
		fmt.Println(s.soal1)
		fmt.Println(s.soal2)
		fmt.Println(s.soal3)
		fmt.Println(s.soal4)
	}
	fmt.Println("==========")
	fmt.Println("jawaban (1,2,3,4)")
	fmt.Println("masukkan jawaban:")
	fmt.Scanln(&z)

	for _, s := range soal_keempat {
		fmt.Println(s.soalU)
		fmt.Println(s.soal1)
		fmt.Println(s.soal2)
		fmt.Println(s.soal3)
		fmt.Println(s.soal4)
	}
	fmt.Println("==========")
	fmt.Println("jawaban (1,2,3,4)")
	fmt.Println("masukkan jawaban:")
	fmt.Scanln(&q)

	if x == 3 {
		nilai_benar = append(nilai_benar, 1)
	} else {
		nilai_salah = append(nilai_salah, 1)
	}
	if y == 4 {
		nilai_benar = append(nilai_benar, 1)
	} else {
		nilai_salah = append(nilai_salah, 1)
	}
	if z == 2 {
		nilai_benar = append(nilai_benar, 1)
	} else {
		nilai_salah = append(nilai_salah, 1)
	}
	if q == 1 {
		nilai_benar = append(nilai_benar, 1)
	} else {
		nilai_salah = append(nilai_salah, 1)
	}

	var total = (len(nilai_benar) * 25)

	fmt.Println("nama	:", nama)
	fmt.Println("total	:", total)
	fmt.Println("benar	:", len(nilai_benar))
	fmt.Println("salah	:", len(nilai_salah))

}
