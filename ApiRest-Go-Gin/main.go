package main

import (
	"apirest/src/lib"
	"apirest/src/routes"
	"log"

	"github.com/gin-gonic/gin"
)


func main() {

 // Crear una instancia de gin.Default() para el enrutador
    app := gin.Default()

    // Conectar a la base de datos
  lib.ConnectToDatabase()
   
 	defer lib.CloseDB()

    // Configurar rutas
    routes.IndexRoutes(app.Group(""))

    // Iniciar el servidor Gin en el puerto 3000
      if err:=app.Run(":3000"); err != nil {
     log.Fatalf("Error iniciando el servidor: %v", err)
    }



}




