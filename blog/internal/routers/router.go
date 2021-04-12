package routers

import (
	_ "github.com/GriezLiao/griez-go-tour/blog/docs"
	"github.com/GriezLiao/griez-go-tour/blog/global"
	"github.com/GriezLiao/griez-go-tour/blog/internal/middleware"
	"github.com/GriezLiao/griez-go-tour/blog/internal/routers/api"
	v1 "github.com/GriezLiao/griez-go-tour/blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	engine.Use(middleware.Translations())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 文件上传路由
	upload := api.NewUpload()
	engine.POST("/upload/file", upload.UploadFile)
	engine.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	engine.GET("/auth", api.GetAuth)


	tag := v1.NewTag()
	article := v1.NewArticle()

	apiv1 := engine.Group("api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	apiv1.Use(middleware.JWT())

	return engine
}
