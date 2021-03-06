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

// GET/books/:id
// Get a single book by ID
func FindBook(c *gin.Context) {
    var book models.Book
    if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Record Not Found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": book})
}

// POST /books
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

// PATCH /books/:id
// Update Book
func UpdateBook(c *gin.Context) {

    var book models.Book
    if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    var input UpdateBookInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Model(&book).Updates(input)
    c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete Book by ID
func DeleteBook(c *gin.Context) {
    var book models.Book
    if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.DB.Delete(&book)
    c.Status(http.StatusNoContent)
}

// input form schemes

type CreateBookInput struct {
    Title string `json:"title" binding:"required"`
    Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
    Title string `json:"title"`
    Author string `json:"author"`
}