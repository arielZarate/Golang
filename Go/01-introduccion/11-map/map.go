package main

import "fmt"


func main() {

mapa1:=map[int]string{

1:"Lunes",
2:"Martes",
3:"Miercoles",
4:"Jueves",
5:"Viernes",
6:"Sabado",
7:"Domingo",
}



///mapa de nuemros
mapa2:=map[string]int{
"uno":1,
"dos":2,
"tres":3,
"cuatro":4,
"cinco":5,
"seis":6,
}


//quiero acceder a un valor 

fmt.Println(mapa1[4])

delete(mapa1, 1)
fmt.Println(mapa1)



fmt.Println(mapa2)



//mapa de manera generica s
mapa3:=make(map[int]string)



mapa3[1]="Lunes"
mapa3[2]="Martes"
mapa3[3]="Miercoles"
mapa3[4]="Jueves"



fmt.Println(mapa3[1])
}