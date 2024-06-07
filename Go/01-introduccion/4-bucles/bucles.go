package main

import "fmt"


func main()  {

	
	for i:= 0; i <10; i++ {
    
		//uso de continue
		/*
		if i==6{

		fmt.Print("es 6!!")
		continue
		} 
		
		*/
		if(i==7){
			fmt.Println(i)
			fmt.Print("es 7 hacemos break")
            break
		}
        
		fmt.Println(i)

	}


	//ciclo while simulacion

  	j:=0;

	 for j<=20{
		fmt.Println(j)
		j++;
	 }
	 


	 //ciclo do-while

	  z:=0;
	  
	  for{
		fmt.Println(z)
		z++;

		if z==5{
			fmt.Println(z)
            break;
        }
	  } 
	
}