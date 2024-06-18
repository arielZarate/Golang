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



/*


func Get(ctx *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := services.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func Created(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func Updated(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.UpdateUser(id, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func Deleted(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := services.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted User By Id " + id})
}

*/