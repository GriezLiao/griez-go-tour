package v1

import (
	"github.com/GriezLiao/griez-go-tour/blog/global"
	"github.com/GriezLiao/griez-go-tour/blog/internal/service"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/app"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/convert"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/errorcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorFormat("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSiz(c)}
	totalTows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.ErrorFormat("svc.CountTag err： %v", err)
		response.ToErrorResponse(errorcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.ListTag(&param, &pager)
	if err != nil {
		global.Logger.ErrorFormat("svc.ListTag err： %v", err)
		response.ToErrorResponse(errorcode.ErrorListTagFail)
		return
	}

	response.ToResponseList(tags, totalTows)
	return
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param:= service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorFormat("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err :=svc.CreateTag(&param)
	if err != nil {
		global.Logger.ErrorFormat("svc.CreateTag err： %v", err)
		response.ToErrorResponse(errorcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param:= service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorFormat("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err :=svc.UpdateTag(&param)
	if err != nil {
		global.Logger.ErrorFormat("svc.UpdateTag err： %v", err)
		response.ToErrorResponse(errorcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param:= service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorFormat("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err :=svc.DeleteTag(&param)
	if err != nil {
		global.Logger.ErrorFormat("svc.DeleteTag err： %v", err)
		response.ToErrorResponse(errorcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
