package helpers

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	TotalRecords int64       `json:"total_records"`
	TotalPages   int         `json:"total_pages"`
	CurrentPage  int         `json:"page"`
	PageSize     int         `json:"page_size"`
	Data         interface{} `json:"data"`
}

func Paginate(db *gorm.DB, model interface{}, page int, pageSize int) (*Pagination, error) {
	var totalRecords int64
	query := db.Model(model)

	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize

	result := db.Limit(pageSize).Offset(offset).Find(model)
	if result.Error != nil {
		return nil, result.Error
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	pagination := &Pagination{
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		CurrentPage:  page,
		PageSize:     pageSize,
		Data:         model,
	}

	return pagination, nil
}
