package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/sathishachilles/gin-bookstore/models"
    "github.com/sathishachilles/gin-bookstore/controllers"
)

func main() {
    r := gin.Default()
    models.ConnectDatabase()
    r.GET("/books", controllers.FindBooks)
    r.Run()
}