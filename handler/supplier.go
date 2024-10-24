package handler

import (
	"api-kasirapp/helper"
	"api-kasirapp/supplier"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type supplierHandler struct {
	supplierService supplier.Service
}

func NewSupplierHandler(supplierService supplier.Service) *supplierHandler {
	return &supplierHandler{supplierService}
}

func (h *supplierHandler) CreateSupplier(c *gin.Context) {
	var input supplier.SupplierInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create supplier failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSupplier, err := h.supplierService.CreateSupplier(input)
	if err != nil {
		response := helper.APIResponse("Create supplier failed", http.StatusBadRequest, "error", gin.H{"message": err.Error()})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create supplier", http.StatusCreated, "success", supplier.FormatSupplier(newSupplier))
	c.JSON(http.StatusCreated, response)
}

func (h *supplierHandler) GetSupplierById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getSupplier, err := h.supplierService.GetByID(id)
	if err != nil {
		response := helper.APIResponse("Get supplier failed", http.StatusNotFound, "error", gin.H{"message": err.Error()})
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Success get supplier", http.StatusOK, "success", supplier.FormatSupplier(getSupplier))
	c.JSON(http.StatusOK, response)
}

func (h *supplierHandler) GetSuppliers(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 4
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	suppliers, err := h.supplierService.GetAll(limit, offset)
	if err != nil {
		response := helper.APIResponse("Get suppliers failed", http.StatusBadRequest, "error", gin.H{"message": err.Error()})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get suppliers", http.StatusOK, "success", supplier.FormatSuppliers(suppliers))
	c.JSON(http.StatusOK, response)
}

func (h *supplierHandler) UpdateSupplier(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input supplier.SupplierInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update supplier failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateSupplier, err := h.supplierService.Update(id, input)
	if err != nil {
		response := helper.APIResponse("Update supplier failed", http.StatusBadRequest, "error", gin.H{"message": err.Error()})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success update supplier", http.StatusOK, "success", supplier.FormatSupplier(updateSupplier))
	c.JSON(http.StatusOK, response)
}

func (h *supplierHandler) DeleteSupplier(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deleteSupplier, err := h.supplierService.Delete(id)
	if err != nil {
		response := helper.APIResponse("Delete supplier failed", http.StatusBadRequest, "error", gin.H{"message": err.Error()})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete supplier", http.StatusOK, "success", supplier.FormatSupplier(deleteSupplier))
	c.JSON(http.StatusOK, response)
}
