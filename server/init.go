package server

import (
	handler "booking-service/handler"
	"booking-service/internal/repos"
	"booking-service/internal/services/sendservice"
)

func (s *Server) initServices(repo repos.IRepo) *ServiceList {
	sendService := sendservice.NewEmployee(repo)
	return &ServiceList{
		sendService: sendService,
	}
}

func (s *Server) initRouters(serviceList *ServiceList) {
	handler := handler.NewHandler(serviceList.sendService)

	handler.ConfigAPIRoute(s.router)
}
