package main

import (
	"net/http"

	"log"

	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	log.SetPrefix("locations: ")
	router.GET("/locations", getLocations)
	router.GET("/locations/:id", getLocation)
	router.PUT("/locations/:id", updateLocation)

	router.Run("localhost:8080")
}

func getLocations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, locations)
}

func getLocation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	loc := GetLocation(id, locations)
	c.IndentedJSON(http.StatusOK, loc)

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "location not found"})
}

func updateLocation(c *gin.Context) {
	id, idErr := strconv.Atoi(c.Param("id"))
	if idErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	var updatedLocation location

	if err := c.BindJSON(&updatedLocation); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
	}

	log.Printf("%+v/n", updatedLocation)

	_, err := UpdateLocation(id, updatedLocation, locations)
	if err == nil {
		c.JSON(http.StatusOK, updatedLocation)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "location not found"})
}
