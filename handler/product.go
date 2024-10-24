package handler

import (
	"api-kasirapp/helper"
	"api-kasirapp/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{
		productService,
	}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var input product.ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create product failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProduct, err := h.productService.CreateProduct(input)
	if err != nil {
		if err.Error() == "product code already exists" {
			response := helper.APIResponse("product code already exists", http.StatusConflict, "error", nil)
			c.JSON(http.StatusConflict, response)
			return
		}
		response := helper.APIResponse("Create product failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create product", http.StatusCreated, "success", product.FormatProduct(newProduct))
	c.JSON(http.StatusCreated, response)
}

func (h *productHandler) GetProducts(c *gin.Context) {
	products, err := h.productService.FindAll()
	if err != nil {
		response := helper.APIResponse("Get products failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get products", http.StatusOK, "success", product.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProductById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getProduct, err := h.productService.FindProductByID(id)
	if err != nil {
		response := helper.APIResponse("Get product failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get product", http.StatusOK, "success", product.FormatProduct(getProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input product.ProductInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update product failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateProduct, err := h.productService.UpdateProduct(id, input)
	if err != nil {
		response := helper.APIResponse("Update product failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success update product", http.StatusOK, "success", product.FormatProduct(updateProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := helper.APIResponse("Invalid ID format", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deleteProduct, err := h.productService.DeleteProduct(id)
	if err != nil {
		response := helper.APIResponse("Delete product failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete product", http.StatusOK, "success", product.FormatProduct(deleteProduct))
	c.JSON(http.StatusOK, response)
}
