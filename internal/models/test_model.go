package models

import (
	"time"
)

// TestModel 测试模型
type TestModel struct {
	ID        int64     `json:"id" xorm:"'id' pk autoincr"`
	Name      string    `json:"name" xorm:"VARCHAR(255) not null comment('测试名称')"`
	Value     string    `json:"value" xorm:"VARCHAR(255) not null comment('测试值')"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

// TableName 返回表名
func (t TestModel) TableName() string {
	return "test_model"
}
