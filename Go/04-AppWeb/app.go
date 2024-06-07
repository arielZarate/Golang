package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//===========funciones=========
 func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res,"bienvenido a la api ")
}



type User struct {

Name string
Skills string
Age int
}


//puedo crear un tenplate 
func Web1(res http.ResponseWriter, req *http.Request){
  template,err:= template.ParseFiles("./index.html")
	user:=User{"Ariel","React",37}
  if err!=nil {
panic(err)
  }else{

template.Execute(res,user)
  
  }
}


func main() {
 //fmt.Println("conectado al localhost:3000" )
http.HandleFunc("/",Index)
http.HandleFunc("/web",Web1)

fmt.Println("run server : http:localhost:3000")
	http.ListenAndServe("localhost:3000",nil) 
		
}