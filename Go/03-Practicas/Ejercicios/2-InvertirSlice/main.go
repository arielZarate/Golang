package main

import "fmt"


func InvertirSlice(array[]int) []int {

	//creao un nuevo slice don le indico el tipo y el tamaño
	reversed:=make([]int,len(array))	


    for i :=len(array)-1; i>=0;i--{
        
	reversed[(len(array)-1) -i]=array[i]
    }
  
return reversed
}


func main(){

// Invertir un Slice

var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
    fmt.Println(array)

	aux:=fmt.Sprintf("%v",InvertirSlice(array))
	
	fmt.Println("Array invertido")
	fmt.Println(aux)




}