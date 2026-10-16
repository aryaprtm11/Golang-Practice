package main

import (
	"fmt"
)

func main() {
	names := [...]string{"Arya","Pratama","Boru","Opung","Kakek","Nenek","Prabowo"}
	
	slice := names[4:6]
	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	//Function Slice
	days := [...]string{"senin","selasa","rabu","kamis","jumat","sabtu","minggu"}
	dayslice1 := days[5:]
	dayslice1[0] = "sabtu baru"
	dayslice1[1] = "minggu baru"
	fmt.Println(dayslice1)

	dayslice2 := append(dayslice1,"libur baru")
	dayslice2[0] = "Ups"
	fmt.Println(dayslice2)
	fmt.Println(days)
	fmt.Println(cap(dayslice2))

	//Function Slice (make)
	newslice := make([]string,2,5)
	newslice[0] = "Arya"
	newslice[1] = "Pratama"
	//newslice[2] = "Ben" -.bakal error karna kita udah nentuin panjang arraynya 2, harusnya menggunakan append
	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	newslice2 := append(newslice,"Ben")
	fmt.Println(newslice2)
	fmt.Println(len(newslice2))
	fmt.Println(cap(newslice2))

	//Function Slice (copy)
	fromslice := days[:]
	toslice := make([]string,len(fromslice),cap(fromslice))
	copy(toslice,fromslice)
	fmt.Println(toslice)

	//Perbedaan Array dan Slice
	iniArray := [...]int{1,2,3,4,5}
	inislice := []int{1,2,3,4,5} //kalau dihilangkan titik tiganya jadi slice

	fmt.Println(iniArray)
	fmt.Println(inislice)
}