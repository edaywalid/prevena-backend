package dto

type PaginatedResponse struct {
	Products   []ProductDTO `json:"products"`
	TotalCount int64        `json:"total_count"`
	Page       int          `json:"page"`
	PerPage    int          `json:"per_page"`
	TotalPages int          `json:"total_pages"`
}
