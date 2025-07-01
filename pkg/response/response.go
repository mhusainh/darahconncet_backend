package response

import "net/http"

type Response struct {
	Meta           Meta        `json:"meta"`
	Data           interface{} `json:"data"`
	Pagination     *Pagination `json:"pagination,omitempty"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Pagination struct {
	Page       int64 `json:"page"`
	PerPage    int64 `json:"per_page"`
	TotalItems int64 `json:"total_items"`
	TotalPages int64 `json:"total_pages"`
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Meta: Meta{Code: http.StatusOK, Message: message},
		Data: data,
	}
}

func SuccessResponseWithPagi(message string, data interface{}, page, perPage, totalItems int64) Response {
	totalPages := calculateTotalPages(totalItems, perPage)
	return Response{
		Meta:           Meta{Code: http.StatusOK, Message: message},
		Data:           data,
		Pagination: &Pagination{
			Page:       page,
			PerPage:    perPage,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
	}
}

func calculateTotalPages(totalItems, perPage int64) int64 {
	if perPage <= 0 {
		return 0
	}
	totalPages := totalItems / perPage
	if totalItems % perPage > 0 {
		totalPages++
	}
	return totalPages
}

func ErrorResponse(code int, message string) Response {
	return Response{
		Meta: Meta{Code: code, Message: message},
		Data: nil,
	}
}
