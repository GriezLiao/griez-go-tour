package api

import (
	"github.com/GriezLiao/griez-go-tour/blog/global"
	"github.com/GriezLiao/griez-go-tour/blog/internal/service"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/app"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/convert"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/errorcode"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(ctx.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errorcode.InvalidParams)
		return
	}

	svc := service.New(ctx.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.ErrorFormat("svc.UploadFile error:v%", err)
		response.ToErrorResponse(errorcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
