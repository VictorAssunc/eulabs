package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"eulabs/pkg/service"
)

type Handler struct {
	service *service.Product
}

func NewHandler(service *service.Product) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateProduct(c echo.Context) error {
	product := new(ProductMutation)
	if err := c.Bind(product); err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "invalid request body")
	}

	p := product.Entity()
	if err := h.service.Create(c.Request().Context(), p); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetProduct(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}

	currencyCode := c.QueryParam("currency")
	if len(currencyCode) != 3 {
		if currencyCode == "" {
			return c.String(http.StatusBadRequest, "currency parameter is required")
		}

		return c.String(http.StatusBadRequest, "invalid currency")
	}

	product, err := h.service.Get(c.Request().Context(), id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	result := NewProductFromEntity(product, currencyCode)
	if result == nil {
		return c.String(http.StatusNotFound, "product not found")
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Handler) UpdateProduct(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}

	oldProduct, err := h.service.Get(c.Request().Context(), id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if oldProduct == nil {
		return c.String(http.StatusNotFound, "product not found")
	}

	product := new(ProductMutation)
	if err = c.Bind(product); err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "invalid request body")
	}

	p := product.Entity()
	p.ID = id
	if err = h.service.Update(c.Request().Context(), oldProduct, p); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) DeleteProduct(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}

	if err := h.service.Delete(c.Request().Context(), id); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
