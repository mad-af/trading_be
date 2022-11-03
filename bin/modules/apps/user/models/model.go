package models

import (
	"time"
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
		RoleID   int    `json:"role_id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	ReqLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
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

	Balances struct {
		ID     string `json:"id"`
		UserID string `json:"user_id"`
		Value  int64  `json:"value"`
	}

	UserGrades struct {
		UserID    string    `json:"user_id"`
		GradeID   int       `json:"grade_id"`
	}
)

// COMMON
type (
	ResLogin struct {
		ID    string `json:"id"`
		Token string `json:"token"`
	}

	UserDetail struct {
		Users
		GradeID      int    `json:"grade_id"`
		GradeName    string `json:"grade_name"`
		RoleName     string `json:"role_name"`
		BalanceValue int    `json:"balance_value"`
	}

	UserList struct {
		Users
		GradeName string `json:"grade_name"`
		RoleName  string `json:"role_name"`
	}
)
