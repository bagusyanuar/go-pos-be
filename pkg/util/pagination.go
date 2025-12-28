package util

type QueryPagination struct {
	Page     int `json:"page" query:"page"`
	PageSize int `json:"page_size" query:"page_size"`
}

type QuerySort struct {
	Sort  string `json:"sort" query:"sort"`
	Order string `json:"order" query:"order"`
}
