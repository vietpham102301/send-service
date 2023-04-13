package sendservice

import (
	"booking-service/internal/models"
	"booking-service/internal/repos"
	"log"
)

type ISendService interface {
	GetAvailableEmpl() (*models.Employee, error)
}

type Employee struct {
	mgRepo repos.IRepo
}

func NewEmployee(repo repos.IRepo) ISendService {
	return &Employee{
		mgRepo: repo,
	}
}

func (e Employee) GetAvailableEmpl() (*models.Employee, error) {
	record, err := e.mgRepo.SendService().GetAvaiableEmpl()
	if err != nil {
		log.Printf("get availale employee fail with err %v", err)
		return nil, err
	}
	return record, nil
}
