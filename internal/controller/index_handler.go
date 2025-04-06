package controller

import (
	"github.com/labstack/echo/v4"
	"go-starter/internal/service"
	"go-starter/utils"
	"net/http"
	"strconv"
	"time"
)

type IndexController struct {
	BusinessGroupService service.BusinessGroupService
}

func InitIndexController(e *echo.Echo, businessGroupService service.BusinessGroupService) {
	controller := &IndexController{
		BusinessGroupService: businessGroupService,
	}
	e.GET("/health", controller.Health)
	e.GET("/", controller.Health)
	e.GET("/test/db/:id", controller.TestDBQuery)
}

func (index *IndexController) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, SimpleResponse{
		Message: "success",
		Data:    time.Now().Format("2006-01-02 15:04:05"),
	})
}

// TestDBQuery godoc
// @Summary 测试数据库查询
// @Description 测试从数据库获取 BusinessGroup 数据
// @Tags Test
// @Accept json
// @Produce json
// @Param id path int true "BusinessGroup ID"
// @Success 200 {object} SimpleResponse{data=models.BusinessGroup} "BusinessGroup 数据"
// @Failure 400,404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /test/db/{id} [get]
func (index *IndexController) TestDBQuery(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "无效的 ID 参数"})
	}

	id := int64(idParam)
	businessGroup, err := index.BusinessGroupService.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, SimpleResponse{
		Message: "success",
		Data:    businessGroup,
	})
}
