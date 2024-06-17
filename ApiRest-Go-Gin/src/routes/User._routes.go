package routes

import (
	"apirest/src/controller"

	"github.com/gin-gonic/gin"
)


func UserRoutes(api *gin.RouterGroup){
	

router:=api.Group("/users")
{
    router.GET("/",controller.Get)

	router.POST("/",controller.Created)
	
	router.GET("/:id",controller.GetById)

	router.PUT("/:id",controller.Updated)

	router.DELETE("/:id",controller.Deleted)

}

}