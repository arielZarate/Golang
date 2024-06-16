package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
)





func main() {

 //PORT :=os.Getenv("PORT");	
app:= gin.Default();


app.GET("/",func(res *gin.Context){

	res.JSON(200, gin.H{
        "message": "Hola mundo",
    })
} )


//grupo
router:=app.Group("/api")


router.GET("/todos",getTodos)






router.GET("/todos/:id",getTodosById)


router.POST("/todos",createTodos)



router.PUT("/todos/:id",updateTodos)


router.DELETE("/todos/:id",deleteTodos)


//servidor 

fmt.Println("http://localhost:4000")
 err:=app.Run(":4000")

  if(err != nil){
		panic(err)
}

}

