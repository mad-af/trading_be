package models

import (
	"time"
	m "trading_be/bin/middleware"
)

type (
	Get struct {
		Query string `query:"query"`
	}

	Pagination struct {
		Page     int      `query:"page" validate:"required"`
		Quantity int      `query:"quantity" validate:"required"`
		Sort     []string `query:"sort"`
		Search   string   `query:"search"`
	}
)

// REQUEST
type (
	ReqCreate struct {
		BankID            int        `json:"bank_id"`
		TransactionTypeID int        `json:"transaction_type_id"`
		Value             float64    `json:"value"`
		Options           m.JwtClaim `json:"opts"`
	}

	ReqUpdate struct {
		Param struct {
			ID string `param:"id" validate:"required"`
		}
		Type        string     `json:"type" validate:"required,oneof=status"`
		Status      string     `json:"status" validate:"oneof=pending rejected canceled transfered checked finalized used"`
		Description string     `json:"description"`
		Options     m.JwtClaim `json:"opts"`
	}

	ReqGetList struct {
		Query struct {
			Pagination
			Get
		}
	}

	ReqGetDetail struct {
		Param struct {
			ID string `param:"id"`
		}
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

	TransactionStatus struct {
		ID            string    `json:"id,omitempty"`
		TransactionID string    `json:"transaction_id,omitempty"`
		Status        string    `json:"status,omitempty"`
		CreatedBy     string    `json:"created_by,omitempty"`
		CreatedAt     time.Time `json:"created_at,omitempty"`
	}

	Users struct {
		ID        string    `json:"id"`
		RoleID    int       `json:"role_id"`
		Name      string    `json:"name"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		Phone     string    `json:"phone"`
		Password  string    `json:"-"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

// COMMON
type (
	TransactionDetail struct {
		TransactionData
		TransactionStatus   []TransactionStatus `json:"transaction_status"`
	}

	TransactionData struct {
		Transactions
		UserName            string    `json:"user_name"`
		BankName            string    `json:"bank_name"`
		TransactionTypeName string    `json:"transaction_type_name"`
		CreatedAt           time.Time `json:"created_at"`
		UpdatedAt           time.Time `json:"updated_at"`
	}
)
