package main

import "fmt";


func main() {


/*
// primero programa escrito
   var age int;

  fmt.Println("ingrese su edad");

  _,error :=fmt.Scan(&age)

  if error!=nil{
fmt.Println("Error",error);
  }

   if age >=18{
	fmt.Println("Eres mayor de edad")
   }else{
	fmt.Println("Eres menor de edad")
   }

*/


fmt.Println("==============dias de la semana con switch========================")

fmt.Println("Ingrese un dia de la semana")
fmt.Printf("Ingrese algunas de las siguientes opciones\n"+
"1- Para Lunes \n"+
"2- Para Martes \n"+
"3- Para Miercoles \n"+
"4- Para Jueves \n"+
"5- Para Viernes \n"+
"6- Para Sabado\n"+
"7- Para domingo\n")

println("Esperando su opcion ...")


var(
   dia int
   salida string;
)

fmt.Scanln(&dia)

switch dia{
case 1:salida ="Lunes";
case 2:salida ="Martes";
case 3:salida ="Miercoles";
case 4:salida ="Jueves";
case 5:salida ="Viernes";
case 6:salida ="Sabado";
case 7:salida ="Domingo";
default:salida ="Opcion no valida";
}

fmt.Printf("eL dia seleccionado es %s \n",salida)

fmt.Println("Programa terminado")

}