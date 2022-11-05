package models

import (
	m "trading_be/bin/middleware"
)

// REQUEST
type (
	ReqTopUp struct {
		TransactionID string     `json:"transaction_id"`
		Options       m.JwtClaim `json:"opts"`
	}
)

// REPOSITORY
type (
	Transactions struct {
		ID                string  `json:"id"`
		UserID            string  `json:"user_id"`
		BankID            int     `json:"bank_id"`
		TransactionTypeID int     `json:"transaction_type_id"`
		Value             float64 `json:"value"`
		Status            string  `json:"status"`
		Description       string  `json:"description"`
	}

	Balances struct {
		ID     string  `json:"id,omitempty"`
		UserID string  `json:"user_id"`
		Value  float64 `json:"value"`
	}
)

// COMMON
type ()
