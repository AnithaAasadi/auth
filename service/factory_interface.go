package factories

import (
	// "task1/model"
)

type EmployeeFactory interface{
	CreateEmployee(employee model.Employee) error
    DeleteEmployee(employeeID string) error
    GetEmployee(employeeID string) (model.Employee, error)
    GetAllEmployees() ([]model.Employee, error)
	
}