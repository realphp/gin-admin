package utils

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page" validate:"required"`
	PageSize int `json:"pageSize" form:"pageSize" validate:"required"`
}
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}
