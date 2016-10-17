package main

import (
"github.com/TheInsideMan/reverseGeoCode/log"
"github.com/gin-gonic/gin"
"net/http"
"fmt"
)

func main() {
gin.SetMode(gin.ReleaseMode)
r := gin.Default()
//	r.Use(controllers.GetURL())
    //router := gin.Default()
//log.Log.Critical("ErrMsg: %s",http.ListenAndServe(":8080", nil))
fmt.Println("Lets see if this prints")
    // This handler will match /user/john but will not match neither /user/ or /user
    r.GET("/test", func(c *gin.Context) {
       // name := c.Param("name")
       // c.String(http.StatusOK, "Hello World")
c.JSON(200, gin.H{"data":"test"}) 
   })
r.NoRoute(func(c *gin.Context){
c.JSON(200,gin.H{"data":"test 2"})
})
r.Run(":8080")
log.Log.Critical("ErrMsg: %s",http.ListenAndServe(":8080", nil))
}
