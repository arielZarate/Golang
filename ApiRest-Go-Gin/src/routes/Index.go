package routes

import (
	"github.com/gin-gonic/gin"
)


func IndexRoutes(){

router:=gin.Default()

api:=router.Group("/api")
{
   Hello(api)
 UserRoutes(api)
}


  err :=router.Run(":3000")
  if err != nil{
	panic(err)
  }


}








