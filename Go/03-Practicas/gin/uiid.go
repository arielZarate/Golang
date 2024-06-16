package main

import (
	"github.com/google/uuid"

	"fmt"
)


func generateUUID()string{
u,err:=uuid.NewRandom()
if err != nil{
	fmt.Println("Error generating UUID")
}
return u.String()
}

