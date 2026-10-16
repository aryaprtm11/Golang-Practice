package main
import "fmt"

func main() {
	name := "Arya"

	switch name {
	case "Arya":
		fmt.Println("Halooo")
	case "Pratama":
		fmt.Println("Haiiii")
	default:
		fmt.Println("Default")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("terlalu panjang")
	case false:
		fmt.Println("benar")
	}
}