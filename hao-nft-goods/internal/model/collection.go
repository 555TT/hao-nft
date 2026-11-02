package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// Collection 藏品model
type Collection struct {
	ID                uint64          `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name              string          `gorm:"column:name;type:varchar(512)" json:"name"`
	Cover             string          `gorm:"column:cover;type:varchar(512)" json:"cover"`
	ClassID           string          `gorm:"column:class_id;type:varchar(128)" json:"class_id"`
	Price             decimal.Decimal `gorm:"column:price;type:decimal(20,3)" json:"price"`
	Quantity          uint64          `gorm:"column:quantity;type:bigint" json:"quantity"`
	Text              string          `gorm:"column:text;type:text" json:"text"`
	SaleableInventory uint64          `gorm:"column:saleable_inventory;type:bigint" json:"saleable_inventory"`
	Identifier        string          `gorm:"column:identifier;type:varchar(128)" json:"identifier"`
	OccupiedInventory uint64          `gorm:"column:occupied_inventory;type:bigint" json:"occupied_inventory"`
	Frozen_inventory  uint64          `gorm:"column:frozen_inventory;type:bigint" json:"frozen_inventory"`
	State             string          `gorm:"column:state;type:varchar(128)" json:"state"`
	SaleTime          time.Time       `gorm:"column:sale_time;type:datetime" json:"sale_time"`             // 藏品发售时间
	SyncChainTime     time.Time       `gorm:"column:sync_chain_time;type:datetime" json:"sync_chain_time"` // 藏品上链时间
	BookStartTime     time.Time       `gorm:"column:book_start_time;type:datetime" json:"book_start_time"` // 预约开始时间
	BookEndTime       time.Time       `gorm:"column:book_end_time;type:datetime" json:"book_end_time"`     // 预约结束时间
	CanBook           bool            `gorm:"column:can_book;type:tinyint(1)" json:"can_book"`             // 是否可预约
	LockVersion       uint32          `gorm:"column:lock_version;type:int" json:"lock_version"`
	CreatorID         string          `gorm:"column:creator_id;type:varchar(128)" json:"creator_id"`
	Version           uint32          `gorm:"column:version;type:int" json:"version"` // 修改版本
	CreatedAt         time.Time       `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
