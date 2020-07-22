package main

import (
    "github.com/gin-gonic/gin"
    "github.com/sathishachilles/gin-bookstore/models"
    "github.com/sathishachilles/gin-bookstore/controllers"
)

func main() {
    r := gin.Default()
    models.ConnectDatabase()
    r.GET("/books", controllers.FindBooks)
    r.POST("/books", controllers.CreateBook)
    r.Run()
}