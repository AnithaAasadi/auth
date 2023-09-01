package factories

import (
	// "task1/model"
)

type EmployeeFactory interface{
	CreateEmployee() string
	DeleteEmployee() string
	GetEmployee() string
	GetAllEmployees() string
}