package services

import (
	"github.com/Kahffi/go-sample-pos/models"
	"github.com/Kahffi/go-sample-pos/repositories"
)

type EmployeeService interface {
	CreateEmployee(employee *models.Employee) error
	GetAllEmployees() ([]models.Employee, error)
	GetEmployeeByID(id string) (*models.Employee, error)
	UpdateEmployee(id string, employee *models.Employee) error
	DeleteEmployee(id string) error
}

type employeeServiceImpl struct {
	Repo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeService {
	return &employeeServiceImpl{repo}
}
func (s *employeeServiceImpl) CreateEmployee(employee *models.Employee) error {
	return s.Repo.Create(employee)
}

func (s *employeeServiceImpl) GetAllEmployees() ([]models.Employee, error) {
	return s.Repo.GetAll()
}

func (s *employeeServiceImpl) GetEmployeeByID(id string) (*models.Employee, error) {
	return s.Repo.GetByID(id)
}

func (s *employeeServiceImpl) UpdateEmployee(id string, employee *models.Employee) error {
	return s.Repo.Update(id, employee)
}

func (s *employeeServiceImpl) DeleteEmployee(id string) error {
	return s.Repo.Delete(id)
}
