package main

import "fmt"

func main() {

	var array1 [5]string
	array1[0] = "Posição 1"

	fmt.Println(array1)

	array2 := [5]string{"Posição 0", "Posição 1", "Posição 2", "Posição 3", "Posição 4"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	array3[0] = array3[len(array3)-1]
	fmt.Println(array3)

	slice := []int{20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(slice)

	slice2 := array2[1:4]
	fmt.Println(slice2)

	array4 := array2[1:4]

	array2[2] = "Posição Alterada"
	fmt.Println(slice2)
	fmt.Println(array4)

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 6)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
