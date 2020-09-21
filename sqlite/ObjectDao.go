package sqlite

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos-domain/v2/model"
)

type ObjectDaoSqlite struct {
	DB *gorm.DB
}

func (dao *ObjectDaoSqlite) Create(object model.StoreObject) error {
	return dao.DB.Save(&object).Error
}

func (dao *ObjectDaoSqlite) Delete(objectId int64) error {
	result := dao.DB.Delete(&model.StoreObject{}, objectId)
	return result.Error
}

func (dao *ObjectDaoSqlite) GetById(objectId int64) (model.StoreObject, error) {
	object := model.StoreObject{}
	result := dao.DB.First(&object, objectId)
	return object, result.Error
}

func (dao *ObjectDaoSqlite) GetByTitle(objectTitle string) ([]model.StoreObject, error) {
	var objects []model.StoreObject
	result := dao.DB.Where("title = ?", objectTitle).Find(&objects)
	return objects, result.Error
}
