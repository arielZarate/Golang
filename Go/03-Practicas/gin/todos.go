package main

//seria como mi interface
type Todo struct {

ID string `json:"id"`
Title string `json:"title"`
Status string `json:"status"`

}

 var todos =[]Todo{
    {ID: "1", Title: "Comprar leche", Status: "pendiente"},
    {ID: "2", Title: "Llamar al médico", Status: "completado"},
    {ID: "3", Title: "Enviar correo electrónico", Status: "pendiente"},
}