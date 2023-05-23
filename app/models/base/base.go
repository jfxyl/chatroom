package base

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel 模型基类
type BaseIDModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
}

type BaseCreatedAtModel struct {
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;not null" json:"created_at"`
}

type BaseTimeModel struct {
	BaseCreatedAtModel
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
}
