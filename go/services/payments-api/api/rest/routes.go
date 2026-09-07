package rest

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	r.GET("/health", h.Health)
	r.POST("/accounts", h.CreateAccount)
	r.GET("/accounts/:id", h.GetAccount)
	r.POST("/accounts/:id/debit", h.Debit)
	r.POST("/payments", h.CreatePayment)
	r.POST("/api/transfer", h.Transfer)
}
