package routes

import (
	"task1/controlllers"
	"github.com/gin-gonic/gin"

)

func Routes() {

	router := gin.Default()

	router.GET("/employees/:id", controllers.GetEmployee)
	router.POST("/employees", controllers.CreateEmployee)
	router.GET("/employees", controllers.GetAllEmployees)
	router.DELETE("/employees/:id", controllers.DeleteEmployee)

	router.Run()
}
