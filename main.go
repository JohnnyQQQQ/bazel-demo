package main

import "github.com/gin-gonic/gin"

func main() {
  r := gin.Default()
  r.GET("/" func(c *gin.Context){
    c.Status(200)
  })
  r.Run("8080")
}
