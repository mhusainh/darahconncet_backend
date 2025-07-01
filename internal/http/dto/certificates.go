package dto

type CertificateGetByIdRequest struct {
	Id int64 `param:"id" validate:"required"`
}

type CertificateGetByDonationIdRequest struct {
	DonationId int64 `param:"donation_id" validate:"required"`
}

type CertificateGetByUserIdRequest struct {
	UserId int64 `param:"user_id" validate:"required"`
}

type GetAllCertificateRequest struct {
	Page       int64  `query:"page"`
	Limit      int64  `query:"limit"`
	Search     string `query:"search"`
	Sort       string `query:"sort"`
	Order      string `query:"order"`
	UserId     string  `query:"user_id"`
	DonationId string  `query:"donation_id"`
}
