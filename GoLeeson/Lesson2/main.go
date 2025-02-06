package main

import (
	"net/http" // Стандартный пакет для работы с HTTP
	"github.com/gin-gonic/gin"
)


func main() {
    
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
       
		c.JSON(http.StatusOK, gin.H{
			"message": "Привет, мир!",
		})
	})

	router.POST("/post", func(c *gin.Context) {
		var jsonData map[string]interface{}
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Данные успешно получены!",
			"data":    jsonData,
		})
	})

   	router.Run(":8080")
}
