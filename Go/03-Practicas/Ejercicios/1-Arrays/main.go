package main

import (
	"fmt"
	"strconv"
)



func main() {


//Ejercicio 1: Iterar sobre un Array y realizar aciones
//Crea un array de enteros y escribe una función que itere sobre él e imprima cada elemento.


//var array[3] string;// primera forma

var numeros = [11]int{1,-12,789,100,45,98,23,0,-45,8748967,12}

for i:=0; i< len(numeros);i++{

fmt.Print( strconv.Itoa(numeros[i]) +" \n\n")	
}



//Crea un slice de enteros y escribe una función que itere sobre él e imprima cada elemento.

numeros2:=[]int{1,-12,999,100,45,98,23,0,-45,20,12}
fmt.Println("Recorriendo slice []int{1,2...,3,4,5,6,7}")
IterarSlice(numeros2)



//Sumar elementos de un array

SumarElementosdeArray(numeros2)

// mayor de los elementos del array

mayor:=EncontrarElMayor(numeros2)

fmt.Println("\n El mayor de los elementos del array :" + strconv.Itoa(mayor))


// encontrar todos los numeros pares y devolverlos en un array

pares:=EncontrarNumerosPares(numeros2)
//pasar primero a cadena el array

strvpares:=fmt.Sprintf("%v",pares)
fmt.Println("\n Los numeros pares del array son: " + strvpares);

}

