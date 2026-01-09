package main

import (
	"fmt"
)

func main () {
	var array1 [3]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	fmt.Println(array1)
	
	array2 := [4]string{"Posição1","Posição2","Posição3","Posiçã4"}
	fmt.Println(array2)
	
	array3 := [...]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(array3)
	
	
	slice := []int{11,12,13,14,15,16,17,18,19,20}
	slice = append(slice, 21)
	slice2 := array3[:]
	slice2 = append(slice2, slice...)

	fmt.Println(slice2)
	
	slice3 := array2[1:3]

	fmt.Println(slice3)

}