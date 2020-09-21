package sqlite

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos-domain/v2/model"
)

type ResourceDaoSqlite struct {
	DB *gorm.DB
}

func (dao *ResourceDaoSqlite) Create(resource model.StoreResource) error {
	return dao.DB.Save(&resource).Error
}

func (dao *ResourceDaoSqlite) Delete(resourceId int64) error {
	result := dao.DB.Delete(&model.StoreResource{}, resourceId)
	return result.Error
}

func (dao *ResourceDaoSqlite) GetById(resourceId int64) (model.StoreResource, error) {
	resource := model.StoreResource{}
	result := dao.DB.First(&resource, resourceId)
	return resource, result.Error
}

func (dao *ResourceDaoSqlite) GetByTitle(resourceTitle string) ([]model.StoreResource, error) {
	var resources []model.StoreResource
	result := dao.DB.Where("title = ?", resourceTitle).Find(&resources)
	return resources, result.Error
}
