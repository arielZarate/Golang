package main

import (
	"fmt"
	"strconv"
)

//=========funciones==========================

func IterarSlice(array[]int){
 for i,num:= range array{
fmt.Println( "En el indice " +strconv.Itoa(i)+":" + strconv.Itoa(num) +" \n\n")
}
}

func SumarElementosdeArray(array[]int){
sum:=0;
for _,num:= range array{
sum+=num	
}

fmt.Println("La suma de los elementos del array es: " + strconv.Itoa(sum))


}


func EncontrarElMayor(array[]int)int{

mayor:=-99999

for _,num:= range array{

if(num >= mayor){
	mayor=num
 }	
}

return mayor
}




//encontrar pares 
func EncontrarNumerosPares(array []int)[]int{

  arrayPares:=[]int{}

  for _,num:= range array{
	if num%2==0{
		arrayPares=append(arrayPares, num)
	}
	}


	return arrayPares
}