package handlers

import (
	"item/application"
	"item/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	itemService *application.ItemService
}

func NewItemHandler(itemService *application.ItemService) *ItemHandler {
	return &ItemHandler{
		itemService: itemService,
	}
}

func (h *ItemHandler) CreateItem(c echo.Context) error {
	var item domain.Item
	if err := c.Bind(&item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.itemService.Create(c.Request().Context(), &item); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, item)
}

func (h *ItemHandler) UpdateItem(c echo.Context) error {
	var item domain.Item
	if err := c.Bind(&item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.itemService.Update(c.Request().Context(), &item); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) DeleteItem(c echo.Context) error {
	var item domain.Item
	if err := c.Bind(&item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.itemService.Delete(c.Request().Context(), &item); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) GetItems(c echo.Context) error {
	items, err := h.itemService.GetAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) GetItem(c echo.Context) error {
	var item *domain.Item
	if err := c.Bind(&item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item, err := h.itemService.GetByID(c.Request().Context(), item.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) GetItemsByUserID(c echo.Context) error {
	var item *domain.Item
	if err := c.Bind(&item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	items, err := h.itemService.GetByUserID(c.Request().Context(), item.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}
