package main
import "fmt"

func main() {
	var name1 = "Arya"
	var name2 = "Arya"

	var result bool = name1 == name2
	fmt.Println(result)

	result = name1 != name2
	fmt.Println(result)

	var name3 = "Pratama"

	result = name1 == name3
	fmt.Println(result)
}