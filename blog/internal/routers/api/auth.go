package api

import (
	"github.com/GriezLiao/griez-go-tour/blog/global"
	"github.com/GriezLiao/griez-go-tour/blog/internal/service"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/app"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/errorcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(ctx *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.ErrorFormat("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(ctx.Request.Context())
	err := svc.CheckToken(&param)
	if err != nil {
		global.Logger.ErrorFormat("svc.CheckToken err： %v", err)
		response.ToErrorResponse(errorcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.ErrorFormat("svc.GenerateToken err： %v", err)
		response.ToErrorResponse(errorcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})

}
