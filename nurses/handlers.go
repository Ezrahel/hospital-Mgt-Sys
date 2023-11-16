package nurses

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func createDoctor(c *gin.Context) {
	var nurse nursesDetails
	c.BindJSON(&nurse)
	db.Create(&nurse)
	c.JSON(200, nurse)
}

func getDoctors(c *gin.Context) {
	var nurse []nursesDetails
	db.Find(&nurse)
	c.JSON(200, nurse)
}

func getDoctor(c *gin.Context) {
	id := c.Params.ByName("id")
	var nurse nursesDetails
	if err := db.First(&nurse, id).Error; err != nil {
		c.AbortWithStatus(404)
		log.Fatal(err)
	} else {
		c.JSON(200, nurse)
	}
}

func updateDoctor(c *gin.Context) {
	var nurse nursesDetails
	id := c.Params.ByName("id")
	if err := db.First(&nurse, id).Error; err != nil {
		c.AbortWithStatus(404)
		log.Fatal(err)
	}
	c.BindJSON(&nurse)
	db.Save(&nurse)
	c.JSON(200, nurse)
}

// Delete a user by ID
func deleteDoctor(c *gin.Context) {
	id := c.Params.ByName("id")
	var nurse nursesDetails
	d := db.Where("id = ?", id).Delete(&nurse)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
