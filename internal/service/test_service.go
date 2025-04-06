package service

import (
	"context"
	"go-starter/internal/models"
	"go-starter/internal/repository/mysql"
	"time"
)

// TestService 测试服务接口
type TestService interface {
	// GetByID 根据ID获取测试数据
	GetByID(ctx context.Context, id int64) (testModel models.TestModel, err error)
	// GetList 获取测试数据列表
	GetList(ctx context.Context, limit int) (testModels []models.TestModel, err error)
}

// TestServiceImpl 测试服务实现
type TestServiceImpl struct {
	repo           mysql.TestRepository
	contextTimeout time.Duration
}

// NewTestService 创建测试服务实例
func NewTestService(testRepo mysql.TestRepository, timeout time.Duration) TestService {
	if testRepo == nil {
		panic("Test Repository is nil")
	}
	if timeout == 0 {
		panic("Timeout is empty")
	}
	return &TestServiceImpl{
		repo:           testRepo,
		contextTimeout: timeout,
	}
}

// GetByID 根据ID获取测试数据
func (s *TestServiceImpl) GetByID(ctx context.Context, id int64) (testModel models.TestModel, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	testModel, err = s.repo.GetByID(ctx, id)
	return
}

// GetList 获取测试数据列表
func (s *TestServiceImpl) GetList(ctx context.Context, limit int) (testModels []models.TestModel, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	testModels, err = s.repo.GetList(ctx, limit)
	return
}
