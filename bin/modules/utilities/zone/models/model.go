package models

// REQUEST
type (
	ReqGetList struct {
		Param struct {
			Zone string `param:"zone" validate:"required,oneof=province district subdistrict village"`
			ID   string `param:"id"`
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
