package dto

type  HealthPassportUpdateRequest struct {
	Id             int64     `param:"id" validate:"required"`
	Status         string    `json:"status" form:"status"`          //'Active', 'Expired', 'Suspended'
}

type GetAllHealthPassportRequest struct {
	Page      int64  `query:"page"`
	Limit     int64  `query:"limit"`
	Search    string `query:"search"`
	Sort      string `query:"sort"`
	Order     string `query:"order"`
	Status    string `query:"status"`
}

type HealthPassportByIdRequest struct {
	Id int64 `param:"id" validate:"required"`
}

type HealthPassportByUserIdRequest struct {
	UserId int64 `param:"user_id" validate:"required"`
}
