package mysql

import (
	"context"
	"go-starter/internal/models"
	"go-starter/utils"
	"xorm.io/xorm"
)

// TestRepository 测试仓库接口
type TestRepository interface {
	// GetByID 根据ID获取测试数据
	GetByID(ctx context.Context, id int64) (res models.TestModel, err error)
	// GetList 获取测试数据列表
	GetList(ctx context.Context, limit int) (res []models.TestModel, err error)
}

// mysqlTestRepository 测试仓库实现
type mysqlTestRepository struct {
	engine *xorm.Engine
}

// NewTestRepository 创建测试仓库实例
func NewTestRepository(engine *xorm.Engine) TestRepository {
	if engine == nil {
		panic("Database engine is null")
	}
	return &mysqlTestRepository{engine: engine}
}

// GetByID 根据ID获取测试数据
func (m *mysqlTestRepository) GetByID(ctx context.Context, id int64) (res models.TestModel, err error) {
	has, err := m.engine.ID(id).Get(&res)
	if err != nil {
		return models.TestModel{}, err
	}
	if !has {
		return res, utils.ErrNotFound
	}
	return
}

// GetList 获取测试数据列表
func (m *mysqlTestRepository) GetList(ctx context.Context, limit int) (res []models.TestModel, err error) {
	if limit <= 0 {
		limit = 10 // 默认限制10条
	}
	
	err = m.engine.Limit(limit).Find(&res)
	if err != nil {
		return nil, err
	}
	
	return
}
