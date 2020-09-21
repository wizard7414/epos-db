package sqlite

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos-domain/v2/model"
)

type ObjectAttributeDaoSqlite struct {
	DB *gorm.DB
}

func (dao *ObjectAttributeDaoSqlite) Create(objectAttribute model.StoreObjectAttribute) error {
	return dao.DB.Save(&objectAttribute).Error
}

func (dao *ObjectAttributeDaoSqlite) Delete(objectId int64, attributeId int64) error {
	objectAttribute := model.StoreObjectAttribute{}
	result := dao.DB.Where("object = ? AND attribute = ?", objectId, attributeId).Delete(&objectAttribute)
	return result.Error
}

func (dao *ObjectAttributeDaoSqlite) GetById(objectId int64, attributeId int64) (model.StoreObjectAttribute, error) {
	objectAttribute := model.StoreObjectAttribute{}
	result := dao.DB.Where("object = ? AND attribute = ?", objectId, attributeId).First(&objectAttribute)
	return objectAttribute, result.Error
}

func (dao *ObjectAttributeDaoSqlite) GetByObject(objectId int64) ([]model.StoreObjectAttribute, error) {
	var objectAttributes []model.StoreObjectAttribute
	result := dao.DB.Where("object = ?", objectId).Find(&objectAttributes)
	return objectAttributes, result.Error
}

func (dao *ObjectAttributeDaoSqlite) GetByAttribute(attributeId int64) ([]model.StoreObjectAttribute, error) {
	var objectAttributes []model.StoreObjectAttribute
	result := dao.DB.Where("attribute = ?", attributeId).Find(&objectAttributes)
	return objectAttributes, result.Error
}
