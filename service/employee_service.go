package service

import (
	"fmt"
	"strconv"

	"github.com/employee/model"
	"github.com/employee/repository"

	//"strconv"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type EmployeeService struct {
	Repo *repository.EmployeeRepository
}

// CreateEmployee creates a new employee
func (service *EmployeeService) CreateEmployee(context *fiber.Ctx) error {
	// Parse request body to extract employee data
	var employee model.Employee // Use the Employee type from the entity package
	if err := context.BodyParser(&employee); err != nil {
		return err
	}
	// Call repository method to save employee data to the database
	err := service.Repo.CreateEmployee(&employee)
	// Return success response
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create employee"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "emplyee has been added"})
	return nil
}

func (service *EmployeeService) GetAllEmployees(context *fiber.Ctx) error {

	employees, err := service.Repo.GetAllEmployees()

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get employees"})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":   "All employees fetch succesfully",
		"employees": employees,
	})

	return nil

}

func (service *EmployeeService) GetEmployeeByID(context *fiber.Ctx) error {
	id := context.Params("id")
	employeeID, err := strconv.Atoi(id)
	if err != nil {
		// ... handle error
		panic(err)
	}
	fmt.Println("user id", employeeID)

	employee, err := service.Repo.GetEmployeeByID(employeeID)

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Employee did not found",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":  " Eployee fetch successfully",
		"Employee": employee,
	})

	return nil

}

func (service *EmployeeService) UpdateEmployee(context *fiber.Ctx) error {
	employee := model.Employee{}
	if err := context.BodyParser(&employee); err != nil {
		return err
	}
	id := context.Params("id")
	employeeID, err := strconv.Atoi(id)
	if err != nil {
		// ... handle error
		panic(err)
	}
	fmt.Println("user id", employeeID)
	employee, err = service.Repo.UpdateEmployee(employeeID, &employee)

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Employee did not found",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":  " Eployee update successfully",
		"Employee": employee,
	})

	return nil

}

func (service *EmployeeService) DeleteEmployee(context *fiber.Ctx) error {

	id := context.Params("id")
	employeeID, err := strconv.Atoi(id)
	if err != nil {
		// ... handle error
		panic(err)
	}
	fmt.Println("user id", employeeID)
	err = service.Repo.DeleteEmployee(employeeID)
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Employee did not found",
		})

		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": " Eployee deleted successfully",
	})

	return nil

}
