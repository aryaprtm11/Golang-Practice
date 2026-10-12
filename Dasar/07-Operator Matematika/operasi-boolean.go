package main
import "fmt"

func main() {
	//Operator &&
	var nilaiakhir = 90
	var absensi = 70

	var lulusNilaiAkhir bool = nilaiakhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)

	//Operator ||
	lulus = lulusNilaiAkhir || lulusAbsensi
	fmt.Println(lulus) 

}