package mysql

import (
	"gorm.io/gorm"
	"maoim/pkg/id"
	"time"
)

type BaseModel struct {
	ID string `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = id.Generate()
	return nil
}

func UnDeletedScope(tx *gorm.DB) *gorm.DB {
	return tx.Where("DeletedAt IS NULL")
}
