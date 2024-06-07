package main

import (
	"fmt"
)


func main() {

const genero string= "masculino"
var  nombre string="Ariel";
var edad int=37;
var estado string="Soltero"

//otra forma

ciudad:="Cordoba Capital";




//fmt.Println(name, edad, estado, ciudad)
//fmt.Println("Yo me llamo " + name+ " tengo " +strconv.Itoa(edad) + " y mi estado civil es "+ estado+ " vivo en "+ ciudad    )







/*
Aquí hay una tabla con algunos especificadores de formato comunes en Go:

%s - Cadena de texto (string)
%d - Entero (int)
%f - Número de punto flotante (float)
%t - Booleano (bool)
%v - Valor de cualquier tipo (versátil)

*/


fmt.Printf("yo me llamo %s , tengo %d años , mi estado civil es %s , vivo en %s y mi genero es %s \n",nombre, edad, estado,ciudad,genero)
}