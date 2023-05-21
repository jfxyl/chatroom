package base

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel 模型基类
type BaseIDModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`
}

type BaseCreatedAtModel struct {
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;not null"`
}

type BaseTimeModel struct {
	BaseCreatedAtModel
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime;not null"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
