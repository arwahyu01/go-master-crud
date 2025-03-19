package pagination

import (
	"math"

	"gorm.io/gorm"
)

type Page struct {
	TotalRecords int64       `json:"total_records"`
	TotalPages   int         `json:"total_pages"`
	CurrentPage  int         `json:"current_page"`
	PageSize     int         `json:"page_size"`
	Data         interface{} `json:"data"`
}

func Paginate(db *gorm.DB, dataModel interface{}, page, pageSize int) (*Page, error) {
	var totalRecords int64
	query := db.Model(dataModel)

	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize

	result := query.Limit(pageSize).Offset(offset).Find(dataModel)
	if result.Error != nil {
		return nil, result.Error
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	return &Page{
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		CurrentPage:  page,
		PageSize:     pageSize,
		Data:         dataModel,
	}, nil
}
