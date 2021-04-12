package dao

import "github.com/GriezLiao/griez-go-tour/blog/internal/model"

func (dao Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(dao.engine)
}
