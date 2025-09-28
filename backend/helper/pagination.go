package helper

import "math"

type Pagination struct {
	Page      int `json:"page"`
	Size      int `json:"Size"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

type PaginationMeta struct {
	Pagination Pagination `json:"pagination"`
}

func (p Pagination) TotalPages() int {
	if p.Size <= 0 {
		return 0
	}
	return int(math.Ceil(float64(p.Total) / float64(p.Size)))
}

func BuildPagination(page, size, total int) PaginationMeta {
	paging := Pagination{
		Page:  page,
		Size:  size,
		Total: total,
	}

	paging.TotalPage = paging.TotalPages()

	return PaginationMeta{
		Pagination: paging,
	}
}
