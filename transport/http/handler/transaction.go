package handler

import (
	"net/http"

	"github.com/dhucsik/onelab-hw4/models"
	"github.com/labstack/echo/v4"
)

// @Summary Create Transaction
// @Tags transaction
// @Description create transaction
// @ID create-transaction
// @Accept json
// @Produce json
// @Param input body models.Transaction true "transaction info"
// @Success 201 {string} string "id"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /transaction [post]
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

// @Summary Update Transaction
// @Tags transaction
// @Description update transaction
// @ID update-transaction
// @Accept json
// @Produce json
// @Param transaction_id path string true "transaction id"
// @Param input body models.Transaction true "transaction info"
// @Success 200
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /transaction/{transaction_id} [put]
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

// @Summary List Transactions
// @Tags transaction
// @Description list transactions
// @ID list-transactions
// @Accept json
// @Produce json
// @Success 200 {object} []models.Transaction
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /transaction [get]
func (h Manager) ListTransactions(c echo.Context) error {
	resp, err := h.srv.Transaction.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary Get Transaction
// @Tags transaction
// @Description get transaction by ID
// @ID get-transaction
// @Accept json
// @Produce json
// @Param transaction_id path string true "transaction id"
// @Success 200 {object} models.Transaction
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /transaction/{transaction_id} [get]
func (h Manager) GetTransaction(c echo.Context) error {
	id := c.Param("id")

	resp, err := h.srv.Transaction.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary Delete Transaction
// @Tags transaction
// @Description delete transcation by ID
// @ID delete-transaction
// @Accept json
// @Produce json
// @Param transaction_id path string true "transaction id"
// @Success 200
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /transaction/{transaction_id} [delete]
func (h Manager) DeleteTransaction(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.Transaction.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
