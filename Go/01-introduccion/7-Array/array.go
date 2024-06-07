package main

import "fmt"

func main() {

var array [2]string  //2 opcion array := [2]string{"Hello", "World"}


array[0]="Ariel"
array[1]="Juan"
	



for i := 0; i < len(array); i++ {
	
	fmt.Println(array[i])
}



//==========================================

arraData:=[2]string{"Ariel", "Juan", }

fmt.Println(arraData)






}

