package handlers

import (
	"product-ms/apps/models/services"
	"product-ms/apps/views"
	"product-ms/libs/helpers"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	productService services.ProductService
}

func NewHandler(productService services.ProductService) *Handler {
	return &Handler{productService: productService}
}

func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	var request views.CreateProductRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(views.Status{Code: fiber.StatusBadRequest, Message: err.Error()})
	}

	result, err := h.productService.CreateProduct(request)
	if err != nil {
		if err == helpers.ErrProductAlreadyExists {
			return c.Status(fiber.StatusConflict).JSON(views.Status{Code: fiber.StatusConflict, Message: err.Error()})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(views.Status{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(views.DefaultResponse{Status: views.Status{Code: fiber.StatusCreated, Message: "product created successfully"}, Data: result})
}

func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var request views.UpdateProductRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(views.Status{Code: fiber.StatusBadRequest, Message: err.Error()})
	}

	result, err := h.productService.UpdateProductByID(id, request)
	if err != nil {
		var statusCode int
		if err == helpers.ErrProductNotFound {
			statusCode = fiber.StatusNotFound
		} else {
			statusCode = fiber.StatusInternalServerError
		}

		return c.Status(statusCode).JSON(views.Status{Code: statusCode, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(views.DefaultResponse{Status: views.Status{Code: fiber.StatusOK, Message: "product updated successfully"}, Data: result})
}

func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.productService.DeleteProductByID(id)
	if err != nil {
		var statusCode int
		if err == helpers.ErrProductNotFound {
			statusCode = fiber.StatusNotFound
		} else {
			statusCode = fiber.StatusInternalServerError
		}

		return c.Status(statusCode).JSON(views.Status{Code: statusCode, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(views.Status{Code: fiber.StatusOK, Message: "product deleted successfully"})
}

func (h *Handler) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.productService.GetProductByID(id)
	if err != nil {
		var statusCode int
		if err == helpers.ErrProductNotFound {
			statusCode = fiber.StatusNotFound
		} else {
			statusCode = fiber.StatusInternalServerError
		}

		return c.Status(statusCode).JSON(views.Status{Code: statusCode, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(views.DefaultResponse{Status: views.Status{Code: fiber.StatusOK, Message: "product retrieved"}, Data: result})
}

func (h *Handler) GetProducts(c *fiber.Ctx) error {
	page := c.QueryInt("page", 0)
	limit := c.QueryInt("limit", 0)

	result, err := h.productService.GetProducts(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(views.Status{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(views.DefaultResponse{Status: views.Status{Code: fiber.StatusOK, Message: "products retrieved"}, Data: result})
}
