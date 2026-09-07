package rest

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	r.GET("/health", h.Health)
	r.POST("/transactions", h.CreateTransaction)
	r.GET("/merchants/:id/daily-totals", h.DailyTotals)
}
