package main

import "fmt"

//En go no existe el el poo, se crean estructuras
//que son clases en otros lenguajes

//asi se crea un nuevo tipo de dato en go
type cadena string;

type Person struct {

Name cadena; 

Age int;
}



func main() {


	per1:=Person{Name: "John",Age: 37};

	fmt.Println(per1)


}




