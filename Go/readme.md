# Lenguaje de Programacion Go

## Para la instalacion de Go

#### Descarga

vaya al siguiente link `https://go.dev/dl/`

#### Instalacion en Linux

```bash
1 -Elimine cualquier instalación anterior de Go eliminando la carpeta /usr/local/go (si existe), luego extraiga el archivo que acaba de descargar en /usr/local, creando un árbol de Go nuevo en /usr/local/go:
$  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz

(Es posible que deba ejecutar el comando como root o mediante sudo).

No descomprima el archivo en un árbol /usr/local/go existente. Se sabe que esto produce instalaciones Go rotas.

2 -Agregue /usr/local/go/bin a la PATHvariable de entorno.
Puede hacer esto agregando la siguiente línea a su $HOME/.profile o /etc/profile (para una instalación en todo el sistema):

export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin


Nota: aplicar los cambios
source ~/.profile


3 -Crear el Espacio de Trabajo de Go:

Aunque no se crea automáticamente, puedes crear un directorio llamado go en tu directorio home para tu espacio de trabajo:

mkdir -p $HOME/go/{bin,src,pkg}


4 -Verifique que haya instalado Go abriendo un símbolo del sistema y escribiendo el siguiente comando:
$ ir versión

4.2 -verificar que las variables de entorno estén configuradas correctamente con:
echo $GOPATH

Esto debería mostrar /home/tu_usuario/go.

```

## Estructura del Repositorio

En este repositorio esta todo el proceso de aprendizaje con Go
Se separa la informacion en varios folder

```bash

01-introduccion
02-Poo
03-Api
04-AppWeb
...
```

## Documentacion Oficial

`https://go.dev/`
`https://go.dev/doc/effective_go` Go Efectivo Aca esta la documentacion con ejemplos ❤️
`https://pkg.go.dev/std` Libreria Estandar donde estn lso paquetes como `fmt`

![]()

## Documentacion de youutube

Existen varios canales uno es

- `Roelcode` tiene uan lista de unos 20 videos 😀
- `Digital Ocean`

### Tips

- Si uds viene de otro lenguaje como `C++` `C` `Java` o `C#` o `Javascript` la curva sera suave
- Tenga en cuenta el que no hace falta poner `;`
- Mucho cuidaod con el identando de llaves ejemplo `if condicion {}` la primera llave va en la declaracion del if y la ultima debe terminar junto a la ultima linea de codigo
- No hace falta usar `()` ejemplo `if num<=8`, esto causara error ==> `if (num <=8)`
- declaracion de variables

  ```go
      var num string = ""
      num:=""   // esta forma es mas sencilla

  ```

- declaracion de arrays

  ```go
  var arra=[2]string
  array[0]="Ariel"
  array[1]="Nahuel"

  //delcarcion y asigancion en una sola linea
  array:=[2]string{"Ariel","Nahuel"}
  ```

- declaracion de slicen

  ```go
    array:=[]string{}

    //lo guardo en el mismo array
    array= append(array,"Pedro");
    array= append(array,"Ariel","Juan","Nadia","Johana");

    fmt.Print(array)
  ```

- clases

  ```go
  type Person struct {
  Name string;
  Age int;
  }

  func main() {
  per1:=Person{Name: "John",Age: 37};
   fmt.Println(per1)
  }
  ```

- Funciones Publicas y Privadas , las publicas tienen la primera palabra en mayucula y podran ser usas en otros paquetes

### Creacion de Package y Modulos

Este tema suele ser facil pero en go por ahi hace cosas raras

- el go.mod es un package.json donde podemos configurar todas las dependencias o un pom.xml en java
- La idea es poder compartir archivos entre distintos directorios vamos a suponer algo asi:

```bash

├── Proyecto
│   ├── main.go
│   └── compartido/
│   |    └── mensaje.go
|   └──go.mod

```

#### go.mod

Creacion de un modulo ` go mod init <nombre de module>`

```bash
  module GoTutorial
  go 1.22.4
```

Aca lo que vemos es que el go.mod esta en la raiz entonces todos los archivos pueden ser accedidos de todas partes

#### configuracion de package

Siguiendo la estructura tendriamos el folder compartido con un archivo llamado `mensaje.go`

```go
package compartido

// Función para obtener un mensaje compartido
func ObtenerMensaje() string {
    return "¡Hola desde el paquete compartido!"
}

```

#### Y en el archivo main.go

```go
package main

import (
    "fmt"
    "GoTutorial/compartido" // Importa el paquete compartido a taves del nombre del modulo
)

func main() {
    mensaje := compartido.ObtenerMensaje()
    fmt.Println(mensaje)
}

```

- Luego ejecutamos go run main.go y saldra una salida por consola como esta

```go
ariel@Linux-Mint:/Programacion/Golang/Go/02-Poo/2-paquetes$ go run main.go
Hola desde main.go ¡Hola desde saluda!
```

## Tutorial: Desarrollo de una API RESTful con Go y Gin

- La documentacion en Go esta en `https://go.dev/doc/tutorial/web-service-gin`

```go
//continuara ...

```
