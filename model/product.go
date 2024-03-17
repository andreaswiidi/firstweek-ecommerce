package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id    uuid.UUID `gorm:"column:id;type:uuid;primary_key;not null;default:uuid_generate_v4()"`
	UPI   string    `gorm:"column:UPI;not null" json:"UPI"`
	Title string    `gorm:"column:title;not null" json:"title"`
	Price float64   `gorm:"column:price;type:numeric(12, 2);not null" json:"price"`
	// BrandId   uuid.UUID `gorm:"column:brand_id;not null" json:"brand_id"`
	IsDeleted bool      `gorm:"column:isdeleted;not null;default:false"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()"`
	CreatedBy string    `gorm:"column:created_by;not null;default:SYSTEM"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:now()"`
	UpdatedBy string    `gorm:"column:updated_by;"`
}
