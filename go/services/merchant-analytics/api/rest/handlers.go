package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"merchant-analytics/internal/analytics"
)

type Handler struct {
	analytics analytics.Service
}

func New(analyticsService analytics.Service) *Handler {
	return &Handler{analytics: analyticsService}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) CreateTransaction(c *gin.Context) {
	var body struct {
		MerchantID string    `json:"merchantId"`
		Amount     float64   `json:"amount"`
		Status     string    `json:"status"`
		CreatedAt  time.Time `json:"createdAt"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := h.analytics.CreateTransaction(c.Request.Context(), &analytics.Transaction{
		MerchantID: body.MerchantID,
		Amount:     body.Amount,
		Status:     body.Status,
		CreatedAt:  body.CreatedAt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tx)
}

func (h *Handler) DailyTotals(c *gin.Context) {
	start, err := time.ParseInLocation("2006-01-02", c.Query("date"), time.UTC)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date, expected YYYY-MM-DD"})
		return
	}

	result, err := h.analytics.DailyTotals(c.Request.Context(), c.Param("id"), start, start.Add(24*time.Hour))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
