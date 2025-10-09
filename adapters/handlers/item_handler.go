package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"item/application"
	"item/domain"
	"item/internal/logger"

	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	itemService *application.ItemService
	logger      domain.Logger
}

func NewItemHandler(itemService *application.ItemService, log domain.Logger) *ItemHandler {
	if log == nil {
		log = logger.GetGlobalLogger()
	}

	return &ItemHandler{
		itemService: itemService,
		logger:      log,
	}
}

func (h *ItemHandler) CreateItem(c echo.Context) error {
	log := h.logger.WithContext(c.Request().Context())

	var item domain.Item
	if err := c.Bind(&item); err != nil {
		log.With("error", err.Error()).Warn("invalid request payload")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := h.itemService.Create(c.Request().Context(), &item); err != nil {
		log.Error("failed to create item", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create item")
	}

	return c.JSON(http.StatusCreated, item)
}

func (h *ItemHandler) UpdateItem(c echo.Context) error {
	log := h.logger.WithContext(c.Request().Context())

	var item domain.Item
	if err := c.Bind(&item); err != nil {
		log.With("error", err.Error()).Warn("invalid update payload")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := h.itemService.Update(c.Request().Context(), &item); err != nil {
		log.Error("failed to update item", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update item")
	}

	return c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) DeleteItem(c echo.Context) error {
	log := h.logger.WithContext(c.Request().Context())

	id, err := getIDParam(c)
	if err != nil {
		log.With("error", err.Error()).Warn("invalid item id")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid item ID")
	}

	if err := h.itemService.Delete(c.Request().Context(), &domain.Item{ID: id}); err != nil {
		log.Error("failed to delete item", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete item")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *ItemHandler) GetItems(c echo.Context) error {
	log := h.logger.WithContext(c.Request().Context())

	items, err := h.itemService.GetAll(c.Request().Context())
	if err != nil {
		log.Error("failed to fetch items", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch items")
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) GetItem(c echo.Context) error {
	log := h.logger.WithContext(c.Request().Context())

	id, err := getIDParam(c)
	if err != nil {
		log.With("error", err.Error()).Warn("invalid item id")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid item ID")
	}

	item, err := h.itemService.GetByID(c.Request().Context(), id)
	if err != nil {
		log.Error("failed to fetch item", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch item")
	}

	return c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) GetItemsByUserID(c echo.Context) error {
	log := h.logger.WithContext(c.Request().Context())

	userID := c.Param("userID")
	if userID == "" {
		log.Warn("user id is required")
		return echo.NewHTTPError(http.StatusBadRequest, "userID is required")
	}

	items, err := h.itemService.GetByUserID(c.Request().Context(), userID)
	if err != nil {
		log.Error("failed to fetch items by user id", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch items")
	}

	return c.JSON(http.StatusOK, items)
}

// getIDParam extracts and validates the ID parameter from the request
func getIDParam(c echo.Context) (int64, error) {
	idStr := c.Param("id")
	if idStr == "" {
		return 0, fmt.Errorf("id parameter is required")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid id format: %v", err)
	}

	if id <= 0 {
		return 0, fmt.Errorf("id must be a positive integer")
	}

	return id, nil
}
