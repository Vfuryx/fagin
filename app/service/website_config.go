package service

import (
	"fagin/app/models/website_config"
	"fagin/pkg/errorw"
)

type websiteConfigService struct{}

var WebsiteConfigService websiteConfigService

// ShowInfo 获取网站信息
func (*websiteConfigService) ShowInfo(id uint, columns []string) (*website_config.WebsiteConfig, error) {
	wc := website_config.New()
	err := wc.Dao().FindById(id, columns)
	if err != nil {
		return nil, errorw.UP(err)
	}
	return wc, nil
}

func (*websiteConfigService) UpdateInfo(id uint, data map[string]interface{}) error {
	err := website_config.NewDao().Update(id, data)
	if err != nil {
		return errorw.UP(err)
	}
	// 删除缓存
	//go func() {
	//	_, err = cache.WebsiteConfig.DelInfo()
	//	if err != nil {
	//		go app.Log().Errorln(err, string(debug.Stack()))
	//	}
	//}()
	return nil
}
