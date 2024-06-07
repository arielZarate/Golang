# Lenguaje Go

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

exportar RUTA=$RUTA:/usr/local/go/bin

Nota: Es posible que los cambios realizados en un archivo de perfil no se apliquen hasta la próxima vez que inicie sesión en su computadora. Para aplicar los cambios inmediatamente, simplemente ejecute los comandos del shell directamente o ejecútelos desde el perfil usando un comando como source $HOME/.profile.

3 -Verifique que haya instalado Go abriendo un símbolo del sistema y escribiendo el siguiente comando:
$ ir versión

4 -Confirme que el comando imprima la versión instalada de Go.

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

## Documentacion de youutube

Existen varios canales uno es

- `Roelcode` tiene uan lista de unos 20 videos 😀
- `Digital Ocean`
-

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
