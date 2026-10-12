package main
import "fmt"

func main() {
	var name string

	name = "Arya Pratama"
	fmt.Println(name)

	name = "Pratama Arya"
	fmt.Println(name)

	var lokasi = "Ryacudu" //Go akan otomatis mengetahui tipe data nya adalah string dari valuenya tanpa perlu kita initial dengan type data string
	fmt.Println(lokasi)

	umur := 24 //: hanya digunakan di awal initialisasi (tidak boleh dua kali)
	fmt.Println(umur)

	umur = 30
	fmt.Println(umur)
}