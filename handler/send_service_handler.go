package handler

import (
	"booking-service/handler/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) getAvailableEmployee() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("start find available employee for job")

		record, err := h.sendService.GetAvailableEmpl()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		var res models.EmployeeResponse
		if record != nil {
			res = models.EmployeeResponse{
				ID:          record.ID,
				Name:        record.Name,
				DateOfBirth: record.DateOfBirth,
			}
			ctx.JSON(http.StatusOK, res)
		} else {
			ctx.JSON(http.StatusOK, map[string]string{
				"status":          "200",
				"employee_status": "No one available",
			})
		}

	}
}
