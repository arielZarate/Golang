package main

import (
	"fmt"
	"strconv"
	// Importa el paquete desde la carpeta actual
)

//funciones

func mod(a int ,  b int )int {
return a%b
}


func main() {

fmt.Println("El resultado del mod es:" +  strconv.Itoa(mod(4,3)))
fmt.Printf("El resultado del mod  es %d \n",  mod(4,3) )


//mt.Println(Multiplicar(4,5))

}