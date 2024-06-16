package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)




func getTodos(c *gin.Context){

	c.JSON(http.StatusOK,todos)
}



func getTodosById(c *gin.Context){
id:=c.Param("id")
//busco en el for
for _,todo:=range todos{
	if todo.ID==id{
		c.JSON(http.StatusOK,todo)
		return
	}
}
c.JSON(http.StatusNotFound,gin.H{"Error":"Todo not found"})

}



func createTodos(c *gin.Context){

 var todo Todo

 err:=c.ShouldBindJSON(&todo)

 if(err != nil){c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})}


 todo.ID=generateUUID()
 //agrego todo al array 
 todos = append(todos,todo) 

 c.JSON(http.StatusCreated,todo)

}


func updateTodos(c *gin.Context){
id:=c.Param(":id")

var updatedTodo Todo

 err:=c.ShouldBindJSON(&updatedTodo)
   if err != nil {c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()}) 
return
}

for i,todo:=range todos{
 if todo.ID==id{
     todos[i].Title=updatedTodo.Title
	 todos[i].Status=updatedTodo.Status
     c.JSON(http.StatusOK,todos[i])
     return
 }
}
}

func deleteTodos(c *gin.Context){

	id:=c.Param("id")

for i,todo:=range todos{
	if todo.ID==id{
        todos=append(todos[:i],todos[i+1:]...)
		c.JSON(http.StatusOK,gin.H{"message":"Todo Deleted"})
        return
    }

	c.JSON(http.StatusNotFound,gin.H{"Error":"Todo not found"})
}
	
}