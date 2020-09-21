package sqlite

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos-domain/v2/model"
)

type AttributeTypeDaoSqlite struct {
	DB *gorm.DB
}

func (dao *AttributeTypeDaoSqlite) Create(attributeType model.StoreAttributeType) error {
	return dao.DB.Save(&attributeType).Error
}

func (dao *AttributeTypeDaoSqlite) Delete(attributeTypeId int64) error {
	result := dao.DB.Delete(&model.StoreAttributeType{}, attributeTypeId)
	return result.Error
}

func (dao *AttributeTypeDaoSqlite) GetById(attributeTypeId int64) (model.StoreAttributeType, error) {
	attributeType := model.StoreAttributeType{}
	result := dao.DB.First(&attributeType, attributeTypeId)
	return attributeType, result.Error
}

func (dao *AttributeTypeDaoSqlite) GetByTitle(attributeTypeTitle string) ([]model.StoreAttributeType, error) {
	var attributeTypes []model.StoreAttributeType
	result := dao.DB.Where("title = ?", attributeTypeTitle).Find(&attributeTypes)
	return attributeTypes, result.Error
}
