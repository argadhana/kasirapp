package handler

import (
	"api-kasirapp/customers"
	"api-kasirapp/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	customerService customers.Service
}

func NewCustomerHandler(customerService customers.Service) *customerHandler {
	return &customerHandler{
		customerService,
	}
}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	var input customers.CustomerInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create customer failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCustomer, err := h.customerService.CreateCustomer(input)
	if err != nil {
		response := helper.APIResponse("Create customer failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create customer", http.StatusCreated, "success", customers.FormatCustomer(newCustomer))
	c.JSON(http.StatusCreated, response)
}

func (h *customerHandler) GetCustomers(c *gin.Context) {
	fetchCustomers, err := h.customerService.FindAll()
	if err != nil {
		if err.Error() == "record not found" {
			response := helper.APIResponse("Get customers failed", http.StatusNotFound, "error", nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponse("Get customers failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get customers", http.StatusOK, "success", customers.FormatCustomers(fetchCustomers))
	c.JSON(http.StatusOK, response)
}

func (h *customerHandler) GetCustomerById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getCustomer, err := h.customerService.FindByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			response := helper.APIResponse("Get customer failed", http.StatusNotFound, "error", nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponse("Get customer failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get customer", http.StatusOK, "success", customers.FormatCustomer(getCustomer))
	c.JSON(http.StatusOK, response)
}

func (h *customerHandler) UpdateCustomer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input customers.CustomerInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update customer failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateCustomer, err := h.customerService.Update(id, input)
	if err != nil {
		response := helper.APIResponse("Update customer failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success update customer", http.StatusOK, "success", customers.FormatCustomer(updateCustomer))
	c.JSON(http.StatusOK, response)
}

func (h *customerHandler) DeleteCustomer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deleteCustomer, err := h.customerService.Delete(id)
	if err != nil {
		response := helper.APIResponse("Delete customer failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete customer", http.StatusOK, "success", customers.FormatCustomer(deleteCustomer))
	c.JSON(http.StatusOK, response)
}
