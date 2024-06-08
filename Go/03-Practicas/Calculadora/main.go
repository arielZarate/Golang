package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main() {

   scanner:= bufio.NewScanner(os.Stdin)

      for{ 
		fmt.Println("Menú de Operaciones Matemáticas")
		fmt.Println("1. Sumar")
		fmt.Println("2. Restar")
		fmt.Println("3. Multiplicar")
		fmt.Println("4. Dividir")
		fmt.Println("5. Salir")
		fmt.Print("Elige una opción: ")
		
		//sino detecta nada en el scanner
		if !scanner.Scan() {
			fmt.Println("Error al leer la entrada. Saliendo...")
			break
		}
		  
		// puede elegir el valor y un error opcional
		//si ingresa  error o fuera de las ocpiones
         choice, err := strconv.Atoi(scanner.Text())
          if err != nil || choice < 1 || choice > 5{
		   fmt.Println("Opcion no valida, por favor elija una opcion correcta")
           continue
		}
       
		//si ingresa 5
		if(choice == 5){
			fmt.Println("Saliendo...")
            break
		}
    		//========valor a========
		// pide el valor y da error sino ingresa 
		fmt.Print("Introduce el primer número: ");
		if!scanner.Scan() {
            fmt.Println("Error al leer la entrada. Saliendo...")
            break
        }

		a,err:=strconv.ParseFloat(strings.TrimSpace(scanner.Text()),64)

		if err != nil {
			fmt.Println("Entrada no valida.Por favor ingresa un numero")
            continue
		}

          ///=====valor b=====
		  fmt.Print("Ingresa el segundo número: ")
		if !scanner.Scan() {
			fmt.Println("Error al leer la entrada. Saliendo...")
			break
		}
			b, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			fmt.Println("Entrada no válida. Por favor, ingresa un número.")
			continue
		}

		switch choice{
		   case 1:
			fmt.Printf("Resultado: %.2f\n\n", sumar(a, b))
		   case 2:
			fmt.Printf("Resultado: %.2f\n\n", restar(a, b))
		   case 3:
			fmt.Printf("Resultado: %.2f\n\n", multiplicar(a, b))
		   case 4:
			if b != 0 {
				fmt.Printf("Resultado: %.2f\n\n", dividir(a, b))
				fmt.Println("==============================================")
			} else {
				//no se peude dividir por 0 da error
				fmt.Println("Error: División por cero no permitida.")
			}
		  }
	
    }//fin de ciclo for

}// fin del main


//funciones fuera del main 
func sumar(a, b float64) float64 {
	return a + b
}

func restar(a, b float64) float64 {
	return a - b
}

func multiplicar(a, b float64) float64 {
	return a * b
}

func dividir(a, b float64) float64 {
	return a / b
}