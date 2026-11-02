package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// Collection 藏品model
type Collection struct {
	ID                uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name              string          `gorm:"type:varchar(512)" json:"name"`
	Cover             string          `gorm:"type:varchar(512)" json:"cover"`
	ClassID           string          `gorm:"type:varchar(128)" json:"class_id"`
	Price             decimal.Decimal `gorm:"type:decimal(20,3)" json:"price"`
	Quantity          uint64          `gorm:"type:bigint" json:"quantity"`
	Text              string          `gorm:"type:text" json:"text"`
	SaleableInventory uint64          `gorm:"type:bigint" json:"saleable_inventory"`
	Identifier        string          `gorm:"type:varchar(128)" json:"identifier"`
	OccupiedInventory uint64          `gorm:"type:bigint" json:"occupied_inventory"`
	Frozen_inventory  uint64          `gorm:"type:bigint" json:"frozen_inventory"`
	State             string          `gorm:"type:varchar(128)" json:"state"`
	SaleTime          time.Time       `gorm:"type:datetime" json:"sale_time"`       // 藏品发售时间
	SyncChainTime     time.Time       `gorm:"type:datetime" json:"sync_chain_time"` // 藏品上链时间
	BookStartTime     time.Time       `gorm:"type:datetime" json:"book_start_time"` // 预约开始时间
	BookEndTime       time.Time       `gorm:"type:datetime" json:"book_end_time"`   // 预约结束时间
	CanBook           bool            `gorm:"type:tinyint(1)" json:"can_book"`      // 是否可预约
	LockVersion       uint32          `gorm:"type:int" json:"lock_version"`
	CreatorID         string          `gorm:"type:varchar(128)" json:"creator_id"`
	Version           uint32          `gorm:"type:int" json:"version"` // 修改版本
	CreatedAt         time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}
