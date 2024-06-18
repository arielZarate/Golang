package routes

import (
	"github.com/gin-gonic/gin"
)


func IndexRoutes(router *gin.RouterGroup){



api:=router.Group("/api")
{
Hello(api)
 UserRoutes(api)
}


}








