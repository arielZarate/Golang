package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)





func Get(ctx *gin.Context){

ctx.JSON(http.StatusOK,gin.H{"message":"Get Users"} )
}


func GetById(ctx *gin.Context){

id:=ctx.Param("id")	
ctx.JSON(http.StatusOK,gin.H{"message":"Get User By Id " + id } )
}




func Created(ctx *gin.Context){

ctx.JSON(http.StatusOK,gin.H{"message":"New User "  } )
}


func Updated(ctx *gin.Context){

id:=ctx.Param("id")	
ctx.JSON(http.StatusOK,gin.H{"message":"Upated User By Id " + id } )
}



func Deleted(ctx *gin.Context){

id:=ctx.Param("id")	
ctx.JSON(http.StatusOK,gin.H{"message":"Deleted User By Id " + id } )
}



