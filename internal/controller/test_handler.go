package controller

import (
	"github.com/labstack/echo/v4"
	"go-starter/internal/service"
	"go-starter/utils"
	"net/http"
	"strconv"
)

// TestController 测试控制器
type TestController struct {
	TestService service.TestService
}

// InitTestController 初始化测试控制器
func InitTestController(e *echo.Echo, testService service.TestService) {
	controller := &TestController{
		TestService: testService,
	}
	
	// 注册路由
	testGroup := e.Group("/test")
	testGroup.GET("/db/:id", controller.GetByID)
	testGroup.GET("/list", controller.GetList)
}

// GetByID godoc
// @Summary 测试数据库查询单条记录
// @Description 测试从数据库获取单条 TestModel 数据
// @Tags Test
// @Accept json
// @Produce json
// @Param id path int true "TestModel ID"
// @Success 200 {object} SimpleResponse{data=models.TestModel} "TestModel 数据"
// @Failure 400,404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /test/db/{id} [get]
func (t *TestController) GetByID(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "无效的 ID 参数"})
	}

	id := int64(idParam)
	testModel, err := t.TestService.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, SimpleResponse{
		Message: "success",
		Data:    testModel,
	})
}

// GetList godoc
// @Summary 测试数据库查询列表
// @Description 测试从数据库获取 TestModel 列表数据
// @Tags Test
// @Accept json
// @Produce json
// @Param limit query int false "限制返回记录数量"
// @Success 200 {object} SimpleResponse{data=[]models.TestModel} "TestModel 列表数据"
// @Failure 400,404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /test/list [get]
func (t *TestController) GetList(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	limit := 10 // 默认10条
	
	if limitStr != "" {
		limitVal, err := strconv.Atoi(limitStr)
		if err == nil && limitVal > 0 {
			limit = limitVal
		}
	}
	
	testModels, err := t.TestService.GetList(c.Request().Context(), limit)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	
	return c.JSON(http.StatusOK, SimpleResponse{
		Message: "success",
		Data:    testModels,
	})
}
