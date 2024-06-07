package main

import "fmt"



type Person struct {

Name string; 

Age int;
}



func main() {


	per1:=Person{Name: "John",Age: 37};

	fmt.Println(per1)


}




