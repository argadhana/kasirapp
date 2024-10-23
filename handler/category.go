package handler

import (
	"api-kasirapp/category"
	"api-kasirapp/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService category.Service
}

func NewCategoryHandler(categoryService category.Service) *categoryHandler {
	return &categoryHandler{
		categoryService,
	}
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input category.CategoryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create category failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCategory, err := h.categoryService.Save(input)
	if err != nil {
		response := helper.APIResponse("Create category failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create category", http.StatusOK, "success", category.FormatCategory(newCategory))
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetCategories(c *gin.Context) {
	categories, err := h.categoryService.FindAll()
	if err != nil {
		response := helper.APIResponse("Get categories failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatedCategories := category.FormatCategories(categories)

	response := helper.APIResponse("Success get categories", http.StatusOK, "success", formatedCategories)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetCategoryById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getCategory, err := h.categoryService.FindByID(id)
	if err != nil {
		response := helper.APIResponse("Get category failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatedCategory := category.FormatCategory(getCategory)

	response := helper.APIResponse("Success get category", http.StatusOK, "success", formatedCategory)
	c.JSON(http.StatusOK, response)
}
