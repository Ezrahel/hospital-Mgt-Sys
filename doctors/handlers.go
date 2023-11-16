package doctors

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func CreateDoctor(c *gin.Context) {
	var doc Doctor
	c.BindJSON(&doc)
	db.Create(&doc)
	c.JSON(200, doc)
}

func GetDoctors(c *gin.Context) {
	var doc []Doctor
	db.Find(&doc)
	c.JSON(200, doc)
}

func GetDoctor(c *gin.Context) {
	id := c.Params.ByName("id")
	var doc Doctor
	if err := db.First(&doc, id).Error; err != nil {
		c.AbortWithStatus(404)
		log.Fatal(err)
	} else {
		c.JSON(200, doc)
	}
}

func UpdateDoctor(c *gin.Context) {
	var user Doctor
	id := c.Params.ByName("id")
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatus(404)
		log.Fatal(err)
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(200, user)
}

// Delete a user by ID
func DleteDoctor(c *gin.Context) {
	id := c.Params.ByName("id")
	var doc Doctor
	d := db.Where("id = ?", id).Delete(&doc)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
