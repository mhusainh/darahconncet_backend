package dto

type HospitalCreateRequest struct {
	Name      string  `json:"name" form:"name" validate:"required"`
	Address   string  `json:"address" form:"address" validate:"required"`
	City      string  `json:"city" form:"city" validate:"required"`
	Province  string  `json:"province" form:"province" validate:"required"`
	Latitude  float64 `json:"latitude" form:"latitude"`
	Longitude float64 `json:"longitude" form:"longitude"`
}

type HospitalUpdateRequest struct {
	Id        int64   `param:"id" form:"id" validate:"required"`
	Name      string  `json:"name" form:"name"`
	Address   string  `json:"address" form:"address"`
	City      string  `json:"city" form:"city"`
	Province  string  `json:"province" form:"province"`
	Latitude  float64 `json:"latitude" form:"latitude"`
	Longitude float64 `json:"longitude" form:"longitude"`
}

type HospitalDeleteRequest struct {
	Id int64 `param:"id" validate:"required"`
}

type HospitalGetByIdRequest struct {
	Id int64 `param:"id" validate:"required"`
}

type GetAllHospitalRequest struct {
	Page   int64  `query:"page" `
	Limit  int64  `query:"limit" `
	Search string `query:"search"`
	Sort   string `query:"sort"`
	Order  string `query:"order"`
}
