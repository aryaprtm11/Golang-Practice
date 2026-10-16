package main
import "fmt"

func main() {
	mahasiswa := map[string]string{
		"name": "Arya",
		"address": "Tebet",
	}

	fmt.Println(mahasiswa["name"])
	fmt.Println(mahasiswa["address"])
	fmt.Println(mahasiswa)

	// Functional map
	book := make(map[string]string)
	book["title"] = "Harry Potter"
	book["author"] = "Arya"
	book["ups"] = "salah"

	fmt.Println(book)
	delete(book,"ups")
	fmt.Println(book)
}
