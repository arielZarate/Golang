package main

import "fmt"

//Los metodos se crean con func
type cadena string

type Curso struct{
	nombre cadena
	catedra cadena
}

func (c *Curso)Incripcion(nom cadena, cat cadena) {
c.nombre=nom
c.catedra=cat
}



//struc carrera
type Carrera struct{
	nombreCarrera string;
	duracion int
	Curso  //de esta forma hace la herenncia
}




//aca ejectuto no mas
func main() {

	c1:=new(Carrera)

	c1.nombreCarrera="Ingenieria en Sistemas";
	c1.duracion=4;
	c1.nombre="AED";
	c1.catedra="Programacion";

    c1.Incripcion("analisis","Electronica")


	fmt.Println(c1)


}