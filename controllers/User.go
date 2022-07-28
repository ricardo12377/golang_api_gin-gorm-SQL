package controllers

import (
	"apigin/dtos"
	"apigin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func FindAllUsers (c *gin.Context){
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateBook (c *gin.Context) {
	var input dtos.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Password: input.Password, Email: input.Email}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindUser(c *gin.Context) {  // Get model if exist
  var user models.User

  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
  // Get model if exist
  var user models.User
  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  // Validate input
  var input dtos.UpdateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&user).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": user})
}


func DeleteUser(c *gin.Context) {
  // Get model if exist
  var user models.User
  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&user)

  c.JSON(http.StatusOK, gin.H{"data": true})
}