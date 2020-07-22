package controllers

import (
    "net/http"
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

// Create Book
func CreateBook(c *gin.Context) {
    var input CreateBookInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

   book := models.Book{Title: input.Title, Author: input.Author}
   models.DB.Create(&book)
   c.JSON(http.StatusOK, gin.H{"data": book})
}

// input form schemes

type CreateBookInput struct {
    Title string `json:"title" binding:"required"`
    Author string `json:"author" binding:"required"`
}