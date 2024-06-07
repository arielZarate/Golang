package main

import "fmt"

func main() {

array:=[]string{}

//lo guardo en el mismo array
array= append(array,"Pedro");
array= append(array,"Ariel","Juan","Nadia","Johana");

fmt.Print(array)

}