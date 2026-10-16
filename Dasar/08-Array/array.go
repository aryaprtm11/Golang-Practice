package main
import "fmt"

func main() {
	var names [3]string
	names[0] = "Arya"
	names[1] = "Pratama"
	names[2] = "Ryatama"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//Array 2 -> menentukan array secara langsung
	var values = [3]int{90, 75, 60}

	fmt.Println(values)

	//Function Array
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}