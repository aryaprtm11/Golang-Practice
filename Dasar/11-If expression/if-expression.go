package main
import "fmt"

func main() {
	name := "Pratama"

	if name == "Arya" {
		fmt.Println("ini benar")
	} else if name == "Pratama" {
		fmt.Println("oke")
	} else {
		fmt.Println("salah")
	}

	// If short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlaalu panjang")
	} else {
		fmt.Println("sudah benar")
	}
}