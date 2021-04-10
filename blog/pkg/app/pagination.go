package app

import (
	"github.com/GriezLiao/griez-go-tour/blog/global"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/convert"
	"github.com/gin-gonic/gin"
)

func GetPage(ctx *gin.Context) int {
	page := convert.StrTo(ctx.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSiz(ctx *gin.Context) int {
	pageSize := convert.StrTo(ctx.Query("pageSize")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}

	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	res := 0
	if page > 0 {
		res = (page - 1) * pageSize
	}
	return res
}
