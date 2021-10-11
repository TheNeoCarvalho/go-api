package controllers

import (
	"strconv"

	"github.com/TheNeoCarvalho/go-api/database"
	"github.com/TheNeoCarvalho/go-api/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find book" + err.Error(),
		})
		return
	}

	c.JSON(200, book)

}
