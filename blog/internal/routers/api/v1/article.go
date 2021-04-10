package v1

import (
	"github.com/GriezLiao/griez-go-tour/blog/pkg/app"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/errorcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary 根据id获取文章
// @Produce  json
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param id query int false "文章id"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get]
func (t Article) Get(c *gin.Context)    {
	app.NewResponse(c).ToErrorResponse(errorcode.ServerError)
}

// @Summary 获取多个文章
// @Produce  json
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (t Article) List(c *gin.Context)   {}

// @Summary 新增文章
// @Produce  json
// @Param title body string true "文章标题" minlength(3) maxlength(100)
// @Param desc body string true "文章描述" minlength(3) maxlength(100)
// @Param content body string true "文章内容" minlength(10)
// @Param cover_image_url body string true "文章封面路径" minlength(3) maxlength(512)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (t Article) Create(c *gin.Context) {}

// @Summary 更新文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Param title body string true "文章标题" minlength(3) maxlength(100)
// @Param desc body string true "文章描述" minlength(3) maxlength(100)
// @Param content body string true "文章内容" minlength(10)
// @Param cover_image_url body string true "文章封面路径" minlength(3) maxlength(512)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (t Article) Update(c *gin.Context) {}

// @Summary 删除文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (t Article) Delete(c *gin.Context) {}
