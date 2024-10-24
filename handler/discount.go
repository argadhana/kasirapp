package handler

import (
	"api-kasirapp/discount"
	"api-kasirapp/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type discountHandler struct {
	discountService discount.Service
}

func NewDiscountHandler(discountService discount.Service) *discountHandler {
	return &discountHandler{discountService}
}

func (h *discountHandler) CreateDiscount(c *gin.Context) {
	var input discount.DiscountInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create discount failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDiscount, err := h.discountService.Create(input)
	if err != nil {
		response := helper.APIResponse("Create discount failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create discount", http.StatusCreated, "success", discount.FormatDiscount(newDiscount))
	c.JSON(http.StatusCreated, response)
}

func (h *discountHandler) GetDiscounts(c *gin.Context) {
	discounts, err := h.discountService.GetAll()
	if err != nil {
		response := helper.APIResponse("Get discounts failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get discounts", http.StatusOK, "success", discount.FormatDiscounts(discounts))
	c.JSON(http.StatusOK, response)
}

func (h *discountHandler) GetDiscountById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getDiscount, err := h.discountService.GetByID(id)
	if err != nil {
		response := helper.APIResponse("Get discount failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get discount", http.StatusOK, "success", discount.FormatDiscount(getDiscount))
	c.JSON(http.StatusOK, response)
}

func (h *discountHandler) UpdateDiscount(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input discount.DiscountInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update discount failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateDiscount, err := h.discountService.Update(id, input)
	if err != nil {
		response := helper.APIResponse("Update discount failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success update discount", http.StatusOK, "success", discount.FormatDiscount(updateDiscount))
	c.JSON(http.StatusOK, response)
}

func (h *discountHandler) DeleteDiscount(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deleteSupplier, err := h.discountService.Delete(id)
	if err != nil {
		response := helper.APIResponse("Delete discount failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete discount", http.StatusOK, "success", discount.FormatDiscount(deleteSupplier))
	c.JSON(http.StatusOK, response)
}
