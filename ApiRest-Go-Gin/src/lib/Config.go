package lib

//ACA CONFUGURO LA BD

// database/database.go

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Esta será la variable global para la conexión a la base de datos


func ConnectToDatabase(){

  //  dbhost := os.Getenv("DB_HOST")
  //  dbPort := os.Getenv("DB_PORT")
  //  dbUser := os.Getenv("DB_USER")
  //  dbPassword := os.Getenv("DB_PASSWORD")
 //   dbName := os.Getenv("DB_NAME")

	//SN := "host= usuario=gorm contraseña=gorm dbname=gorm puerto=9920 sslmode=disable TimeZone=Asia/Shanghai"

   // Construir DSN (Data Source Name)
   // DSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Argentina/Buenos_Aires",dbhost , dbPort, dbUser, dbPassword, dbName)
   
 

  
  


DSN := "host=localhost user=postgres password=admin dbname=db_go port=5432 sslmode=disable TimeZone=America/Argentina/Buenos_Aires"
   _, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})


	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		os.Exit(1)
	}



	// Migrar modelos si es necesario
    // AutoMigrateModels()


	fmt.Println("Database connected!")
	//return DB,nil
}



func CloseDB() {
    db, err := DB.DB()
    if err != nil {
        log.Fatal("Error getting DB instance: ", err)
    }
    db.Close();
}