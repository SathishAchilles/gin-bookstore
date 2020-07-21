package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/sathishachilles/gin-bookstore/models"
)

// GET /books
// GET all books
func FindBooks(c *gin.Context) {
    var books []models.Book

    models.DB.Find(&books)

    c.JSON(http.StatusOK, gin.H{"data": books})
}