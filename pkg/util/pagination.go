package util

import "gorm.io/gorm"

const MaxPageSize = 100

type QueryPagination struct {
	Page     int `json:"page" query:"page"`
	PageSize int `json:"page_size" query:"page_size"`
}

type QuerySort struct {
	Sort  string `json:"sort" query:"sort"`
	Order string `json:"order" query:"order"`
}

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	TotalPage   int   `json:"total_page"`
	TotalItems  int64 `json:"total_items"`
	PageSize    int   `json:"page_size"`
}

func Paginate(db *gorm.DB, page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		switch {
		case pageSize > MaxPageSize:
			pageSize = MaxPageSize
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetTotalPages(totalRows int64, pageSize int) int {
	return int((totalRows + int64(pageSize) - 1) / int64(pageSize))
}

func MakePagination(page, pageSize int, totalRows int64) PaginationMeta {
	return PaginationMeta{
		CurrentPage: page,
		PageSize:    pageSize,
		TotalItems:  totalRows,
		TotalPage:   GetTotalPages(totalRows, pageSize),
	}
}
