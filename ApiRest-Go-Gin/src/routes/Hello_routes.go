package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func Hello(api *gin.RouterGroup){


api.GET("/",
func (ctx *gin.Context) {
ctx.JSON(http.StatusOK,gin.H{"Message":"Bienvenido a la api con Go y Gin"})
})


}