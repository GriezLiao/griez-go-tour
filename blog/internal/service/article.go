package service

type CountArticleRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title string `form:"title" binding:"required, min=3,max=100"`
	Desc string `form:"desc" binding:"required, min=3,max=100"`
	Content string `form:"content" binding:"required, min=3,max=100"`
	CreateBy string  `form:"create_by" binding:"required, min=3,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID uint32 `form:"id" binding:"required, gte=1"`
	Title string `form:"title" binding:"required, min=3,max=100"`
	Desc string `form:"desc" binding:"required, min=3,max=100"`
	Content string `form:"content" binding:"required, min=3,max=100"`
	ModifiedBy string  `form:"modified_by" binding:"required, min=3,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required, gte=1"`
}