package main

import "fmt"



func main() {

slices1:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13}

slices2:=[]int{14,15,16,17,18,19,20,21,22,23,24,25,26}


fmt.Println(slices1)
fmt.Println(slices2)


nuevoSlcies:= append(slices1,slices2... )

fmt.Printf("Nuevo slices ")
fmt.Println(nuevoSlcies)

}
