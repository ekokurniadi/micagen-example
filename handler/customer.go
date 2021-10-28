package handler

import (
	"net/http"
	"tesss/formatter"
	"tesss/helper"
	"tesss/input"
	"tesss/service"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) *customerHandler {
	return &customerHandler{service}
}
func (h *customerHandler) GetCustomer(c *gin.Context) {
	var input input.InputIDCustomer
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Customer", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	customerDetail, err := h.service.CustomerServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Campaign Detail", http.StatusOK, "success", formatter.FormatCustomer(customerDetail))
	c.JSON(http.StatusOK, response)

}
