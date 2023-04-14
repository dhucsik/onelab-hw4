package handler

import (
	"net/http"

	"github.com/dhucsik/onelab-hw4/models"
	"github.com/labstack/echo"
)

func (h Manager) CreateTransaction(c echo.Context) error {
	var req models.Transaction
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.srv.Transaction.Create(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h Manager) UpdateTransaction(c echo.Context) error {
	id := c.Param("id")

	var req models.Transaction
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.srv.Transaction.Update(c.Request().Context(), id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h Manager) ListTransactions(c echo.Context) error {
	resp, err := h.srv.Transaction.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) GetTransaction(c echo.Context) error {
	id := c.Param("id")

	resp, err := h.srv.Transaction.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) DeleteTransaction(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.Transaction.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
