package main

import "fmt"


func Sumar(a float32 , b float32 )float32  {
	return a +b
}

func Restar(a float32 , b float32 )float32  {
	return a-b
}

func Dividir(a float32 , b float32 )float32  {
	
	if a==0 || b==0{
		 fmt.Println("No se puede dividir por 0")
	}

	return a/b;
}

func Multiplicar(a float32 , b float32 )float32  {
	return a*b;
}


