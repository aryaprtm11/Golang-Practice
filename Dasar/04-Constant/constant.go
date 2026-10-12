package main
import "fmt"

func main() {
	const firstname = "Arya"
	const lastname = "Pratama"

	fmt.Println(firstname)
	fmt.Println(lastname)

	//error
	//firstname = "Pratama"
	//lastname = "Arya"

	//Multiple Constant
	const (
		lokasi = "Ryacudu"
		umur = 24
	)

	fmt.Println(lokasi)
	fmt.Println(umur)

	//error
	//lokasi = "Bandung"
	//umur = "25"
}