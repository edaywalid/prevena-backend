package dto

import "github.com/edaywalid/pinktober-hackathon-backend/internal/models"

type PaginatedResponse struct {
	Products   []models.Product `json:"products"`
	TotalCount int64            `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	TotalPages int              `json:"total_pages"`
}
