package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task1/database"
	"task1/model"

	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetEmployee(c *gin.Context) {
	db := database.Setup()
	id := c.Param("id")

	var employee models.Employee
	if err := db.Preload("Contact.Address").First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	employeeMap := make(map[string]interface{})
	employeeJSON, _ := json.Marshal(employee)
	_ = json.Unmarshal(employeeJSON, &employeeMap)

	c.JSON(http.StatusOK, employeeMap)
	fmt.Print(employeeMap)
}

func DeleteEmployee(c *gin.Context) {
	db := database.Setup()
	id := c.Param("id")

	var employee models.Employee
	if err := db.Preload("Contact.Address").First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	if err := db.Delete(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

func CreateEmployee(c *gin.Context) {
	db:= database.Setup()
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&employee)

	c.JSON(http.StatusCreated, employee)
}
func GetAllEmployees(c *gin.Context) {
	db := database.Setup()
	var employees []models.Employee
	if err := db.Preload("Contact.Address").Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, employees)

}
