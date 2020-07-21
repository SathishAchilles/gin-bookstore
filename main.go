package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/sathishachilles/gin-bookstore/models"
)

func main() {
    r := gin.Default()
    models.ConnectDatabase()
    r.Run()
}