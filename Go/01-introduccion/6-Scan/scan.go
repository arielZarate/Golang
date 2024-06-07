package main

import "fmt"


func main() {

var name,lastName string	
fmt.Println("ingrese un nombre")
fmt.Scanln(&name)
fmt.Println("ingrese un apellido")
fmt.Scanln(&lastName)

fmt.Printf("Ud se llama %s %s \n",name,lastName)

}
