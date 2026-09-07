package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"payments-api/internal/account"
	"payments-api/internal/payment"
	"payments-api/internal/transfer"
)

type Handler struct {
	accounts  account.Service
	payments  payment.Service
	transfers transfer.Service
}

func New(accounts account.Service, payments payment.Service, transfers transfer.Service) *Handler {
	return &Handler{
		accounts:  accounts,
		payments:  payments,
		transfers: transfers,
	}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var body struct {
		ID         string  `json:"id"`
		OwnerName  string  `json:"ownerName"`
		Balance    float64 `json:"balance"`
		MerchantID string  `json:"merchantId"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a, err := h.accounts.Create(c.Request.Context(), &account.Account{
		ID:         body.ID,
		OwnerName:  body.OwnerName,
		Balance:    body.Balance,
		MerchantID: body.MerchantID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, a)
}

func (h *Handler) GetAccount(c *gin.Context) {
	a, err := h.accounts.Get(c.Request.Context(), c.Param("id"))
	if errors.Is(err, account.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, a)
}

func (h *Handler) Debit(c *gin.Context) {
	var body struct {
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance, err := h.accounts.Debit(c.Request.Context(), c.Param("id"), body.Amount)
	if errors.Is(err, account.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	if errors.Is(err, account.ErrInsufficientFunds) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Insufficient funds"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

func (h *Handler) CreatePayment(c *gin.Context) {
	var body struct {
		AccountID string  `json:"accountId"`
		Amount    float64 `json:"amount"`
		Currency  string  `json:"currency"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := h.payments.Create(c.Request.Context(), body.AccountID, body.Amount, body.Currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (h *Handler) Transfer(c *gin.Context) {
	var body struct {
		FromAccountID string  `json:"fromAccountId"`
		ToAccountID   string  `json:"toAccountId"`
		Amount        float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance, err := h.transfers.Transfer(c.Request.Context(), body.FromAccountID, body.ToAccountID, body.Amount)
	if errors.Is(err, transfer.ErrInsufficientFunds) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient funds"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "newBalance": balance})
}
