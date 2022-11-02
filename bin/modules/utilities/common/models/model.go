package models

// REQUEST
type (
	ReqCreate struct {
		Param struct {
			Table string `param:"table" validate:"required,oneof=banks roles grades transaction_types"`
		}
		Data interface{} `json:"data"`
	}

	ReqGetList struct {
		Param struct {
			Table string `param:"table" validate:"required,oneof=banks roles grades transaction_types"`
		}
	}
)

// REPOSITORY
type (
	Zone struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
