package omsApi

import "math"

// Pagination constants
const (
	paginationDefaultLimit = 100
	paginationMaxLimit     = 1000
)

// PaginationReq holds pagination http fields and tags
type PaginationReq struct {
	Limit int `query:"limit"`
	Page  int `query:"page" validate:"min=0"`
}

// Transform checks and converts http pagination into database pagination model
func (p *PaginationReq) Transform() *Pagination {
	if p.Limit < 1 {
		p.Limit = paginationDefaultLimit
	}

	if p.Limit > paginationMaxLimit {
		p.Limit = paginationMaxLimit
	}

	if p.Page == 0 {
		p.Page = 1
	}

	return &Pagination{Limit: p.Limit, Offset: (p.Page - 1) * p.Limit, CurrPage: p.Page}
}

func Page(limit, page, total int) *Pagination {
	result := &Pagination{
		Limit:     limit,
		CurrPage:  page,
		TotalPage: int(math.Ceil(float64(total) / float64(limit))),
		TotalRows: total,
	}

	return result
}

// Pagination holds paginations data
type Pagination struct {
	Limit     int `json:"limit,omitempty"`
	Offset    int `json:"-"`
	CurrPage  int `json:"curr_page,omitempty"`
	TotalPage int `json:"total_page,omitempty"`
	TotalRows int `json:"total_rows,omitempty"`
}
