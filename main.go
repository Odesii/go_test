package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var items = []Item{}

func main() {
	r := gin.Default()

	r.GET("/items", getItems)
	r.GET("/items/:id", getItemsById)
	r.POST("/items", createItem)
	r.PUT("/items/:id", updateItem)
	r.DELETE("/items/:id", deleteItem)


	r.Run(":3000")
}

func getItems(c *gin.Context){
	c.JSON(http.StatusOK, items)
}

func getItemsById(c *gin.Context){
	id := c.Param("id")

	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func createItem(c *gin.Context){
	var newItem Item

	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	items = append(items, newItem)
	c.JSON(http.StatusCreated, newItem)
}

func updateItem(c *gin.Context){
	id := c.Param("id")
	var updatedItem Item

	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, item := range items {
		if item.ID == id {
			items[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func deleteItem(c *gin.Context){
	id := c.Param("id")

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}