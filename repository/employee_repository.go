package repository

import (
	"github.com/employee/model"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

// CreateEmployee creates a new employee record in the database
func (repo EmployeeRepository) CreateEmployee(employee *model.Employee) error {
	// Implement logic to create an employee record in the database

	err := repo.DB.Create(employee).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo EmployeeRepository) GetAllEmployees() ([]model.Employee, error) {
	employees := []model.Employee{}
	err := repo.DB.Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, err

}
func (repo EmployeeRepository) GetEmployeeByID(id int) (model.Employee, error) {
	employee := model.Employee{}
	err := repo.DB.First(&employee, id).Error
	if err != nil {
		return model.Employee{}, err
	}
	return employee, err

}

func (repo EmployeeRepository) UpdateEmployee(id int, employee *model.Employee) (model.Employee, error) {
	employ := model.Employee{}
	err := repo.DB.Model(&employ).Where("id = ?", id).Updates(employee).Error
	if err != nil {
		return employ, err
	}
	return employ, err
}

func (repo EmployeeRepository) DeleteEmployee(id int) error {
	employee := model.Employee{}
	err := repo.DB.Delete(employee, id).Error

	if err != nil {
		return err
	}
	return nil

}
