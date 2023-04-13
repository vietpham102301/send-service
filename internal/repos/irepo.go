package repos

import "booking-service/internal/models"

type IRepo interface {
	SendService() ISendServiceRepo
}

type ISendServiceRepo interface {
	GetAvaiableEmpl() (*models.Employee, error)
}
