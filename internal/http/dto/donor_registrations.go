package dto

type DonorRegistrationCreateRequest struct {
	UserId    int64  `json:"user_id" form:"user_id" validate:"required"`
	RequestId int64  `json:"request_id" form:"request_id" validate:"required"`
	Notes     string `json:"notes" form:"notes"` // Additional notes for the registration
}

type DonorRegistrationUpdateRequest struct {
	Id     int64  `param:"id" validate:"required"`
	Status string `json:"status" form:"status"` //'Registered', 'Completed', 'Cancelled', 'No-show'
	Notes  string `json:"notes" form:"notes"`  // Additional notes for the registration
}

type DonorRegistrationByIdRequest struct {
	Id int64 `param:"id" validate:"required"`
}

type DonorRegistrationByUserIdRequest struct {
	UserId int64 `param:"user_id" validate:"required"`
}

type GetAllDonorRegistrationRequest struct {
	UserId int64  `query:"user_id"`
	Page   int64  `query:"page"`
	Limit  int64  `query:"limit"`
	Search string `query:"search"`
	Sort   string `query:"sort"`
	Order  string `query:"order"`
	Status string `query:"status"`
}
