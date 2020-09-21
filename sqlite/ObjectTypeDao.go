package sqlite

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos-domain/v2/model"
)

type ObjectTypeDaoSqlite struct {
	DB *gorm.DB
}

func (dao *ObjectTypeDaoSqlite) Create(objectType model.StoreObjectType) error {
	return dao.DB.Save(&objectType).Error
}

func (dao *ObjectTypeDaoSqlite) Delete(objectTypeId int64) error {
	result := dao.DB.Delete(&model.StoreObjectType{}, objectTypeId)
	return result.Error
}

func (dao *ObjectTypeDaoSqlite) GetById(objectTypeId int64) (model.StoreObjectType, error) {
	objectType := model.StoreObjectType{}
	result := dao.DB.First(&objectType, objectTypeId)
	return objectType, result.Error
}

func (dao *ObjectTypeDaoSqlite) GetByTitle(objectTypeTitle string) ([]model.StoreObjectType, error) {
	var objectTypes []model.StoreObjectType
	result := dao.DB.Where("title = ?", objectTypeTitle).Find(&objectTypes)
	return objectTypes, result.Error
}
