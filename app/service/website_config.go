package service

import (
	"fagin/app/models/website_config"
)

type websiteConfigService struct{}

var WebsiteConfigService websiteConfigService

// 获取网站信息
func (websiteConfigService) ShowInfo(id uint, columns []string) (*website_config.WebsiteConfig, error) {
	wc := website_config.New()
	err := wc.Dao().FindById(id, columns)
	if err != nil {
		return nil, err
	}
	return wc, err
}

func (websiteConfigService) UpdateInfo(id uint, data map[string]interface{}) error {
	err := website_config.Dao().Update(id, data)
	if err != nil {
		return err
	}
	// 删除缓存
	//go func() {
	//	_, err = cache.WebsiteConfig.DelInfo()
	//	if err != nil {
	//		log.Log.Errorln(err)
	//	}
	//}()
	return nil
}
