package main

import (
	"fmt"
	"math/rand"
	"time"
)

var n, l, p, k, x int
var parameter, namaParameter string
var params = make([]int, 0)

var dataMahasiswa = []string{
	"Faarhurrahman Arief",
	"Freddy Lorenzo",
	"Habib Marzuqi",
	"Intan Permatasari",
	"Irvan Nizuar",
	"Jeffri Maulana Saputra",
	"Jerry Olbinson",
	"Lubi Arsada",
	"Miarinda Yulianti Dewi",
	"Mohammad Husnul Rahmadi",
	"Muhamad Ghufron",
	"Muhammad Ahsan Fuady",
	"Octaviani Hutapea",
	"Ongky Sulistyo Wibowo",
	"Razky Satria Ganesha",
	"Reky Rolen Kencana",
	"Reynaldo Mohammad Gozzal",
	"Rizka Khairunnisah",
	"Saifurrahim Widya Prawira",
	"Shesia Rizki Damara",
	"Nabilah Aziz",
	"Santi Dwi Ratnasari",
	"Agung Satria",
	"Agus Ramdani",
	"Burhan Mafazi",
	"Faisal Akbar",
	"Reza Fitrah Amirul Mukmin",
	"Riyan Purnama",
	"Vincent Vallensky Maulana",
	"Tri Endah Sari",
}

// var n int

func main() {
	// fmt.Print("Masukan Jumlah Orang : ")
	// _, err := fmt.Scanf("%v\n", &n)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	fmt.Print("Masukan Jumlah Kelompok : ")
	_, err := fmt.Scanf("%v\n", &k)
	if err != nil {
		fmt.Println(err.Error())
	}

	// fmt.Print("Berapa Kali Acak : ")
	// _, err = fmt.Scanf("%v\n", &x)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	fmt.Print("Butuh Parameter : ")
	_, err = fmt.Scanf("%v\n", &parameter)
	if err != nil {
		fmt.Println(err.Error())
	}

	if parameter == "y" || parameter == "Y" {
		fmt.Print("Nama Parameter : ")
		_, err = fmt.Scanf("%v\n", &namaParameter)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("Masukan Jumlah %s : ", namaParameter)
		_, err = fmt.Scanf("%v\n", &p)
		if err != nil {
			fmt.Println(err.Error())
		}

		for i := 0; i < p; i++ {
			var temp int
			fmt.Printf("Masukan absen %s - %d : ", namaParameter, i+1)
			_, err = fmt.Scanf("%v\n", &temp)
			if err != nil {
				fmt.Println(err.Error())
			}
			params = append(params, temp)
		}
	}

	rand.Seed(time.Now().UnixNano())
	n := len(dataMahasiswa)

	data := rand.Perm(n)

	for i := 0; i < x; i++ {
		data = rand.Perm(n)
	}
	totalAnggota := n / k
	totalParams := len(params) / k
	sisa := n % k
	sisaParams := totalParams % k
	// fmt.Println(sisaParams)
	tmp := 0
	tmpParams := 0
	tmpDataIndex := 0

	kelompok := make([][]int, k)
	for i := 0; i < k; i++ {
		kelompok[i] = make([]int, 0, k)
		anggota := make([]int, totalAnggota+1)
		for j := 0; j < len(data); j++ {
			if sisa > 0 {
				if j > totalAnggota {
					break
				}
				if tmpParams < totalParams {
					for y := 0; y < len(params); y++ {
						if tmpParams < totalParams {
							for z := 0; z < len(data); z++ {
								if data[z]+1 == params[y] {
									anggota[j] = data[z] + 1
									params = append(params[:y], params[y+1:]...)
									tmpDataIndex = z
									tmpParams++
									break
								}
							}
						}
					}
				} else {
					if tmpDataIndex == 0 {
						anggota[j] = data[tmp] + 1
					} else {
						for _, val := range data {
							if val+1 == data[tmpDataIndex]+1 {
								continue
							}
							check := false
							for _, valAnggota := range anggota {
								if val+1 == valAnggota {
									check = true
									break
								}
							}
							checkKelompok := false
							for _, valKelompok := range kelompok {
								if checkKelompok {
									check = true
									break
								}
								for _, valAnggota := range valKelompok {
									if val+1 == valAnggota {
										checkKelompok = true
										break
									}
								}
							}
							for _, valParams := range params {
								if val+1 == valParams {
									if sisaParams <= 0 {
										check = true
										break
									}
									sisaParams--
									break
								}
							}
							if check {
								continue
							}
							anggota[j] = val + 1
							break
						}
					}
				}
				kelompok[i] = append(kelompok[i], anggota[j])
			} else {
				if j > totalAnggota-1 {
					break
				}
				if tmpParams < totalParams {
					for y := 0; y < len(params); y++ {
						if tmpParams < totalParams {
							for z := 0; z < len(data); z++ {
								if data[z]+1 == params[y] {
									anggota[j] = data[z] + 1
									params = append(params[:y], params[y+1:]...)
									tmpDataIndex = z
									tmpParams++
									break
								}
							}
						}
					}
				} else {
					if tmpDataIndex == 0 {
						anggota[j] = data[tmp] + 1
					} else {
						for _, val := range data {
							if val+1 == data[tmpDataIndex]+1 {
								continue
							}
							check := false
							for _, valAnggota := range anggota {
								if val+1 == valAnggota {
									check = true
									break
								}
							}
							checkKelompok := false
							for _, valKelompok := range kelompok {
								if checkKelompok {
									check = true
									break
								}
								for _, valAnggota := range valKelompok {
									if val+1 == valAnggota {
										checkKelompok = true
										break
									}
								}
							}
							for _, valParams := range params {
								if val+1 == valParams {
									if sisaParams <= 0 {
										check = true
										break
									}
									sisaParams--
									break
								}
							}
							if check {
								continue
							}
							anggota[j] = val + 1
							break
						}
					}
				}
				kelompok[i] = append(kelompok[i], anggota[j])
			}
			tmp++
		}
		tmpParams = 0
		sisa--
	}

	for index, dataKelompok := range kelompok {
		fmt.Println("================================")
		fmt.Println("Kelompok : ", index+1)
		fmt.Println("================================")
		for no, anggota := range dataKelompok {
			fmt.Printf("%d. %s\n", no+1, dataMahasiswa[anggota-1])
		}
		fmt.Println()
		// fmt.Println(dataKelompok)
	}

	tahan := 0
	_, err = fmt.Scanf("%v", &tahan)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func checkInsert() int {
	return 0
}
