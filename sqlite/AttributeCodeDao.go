package sqlite

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos-domain/v2/model"
)

type AttributeCodeDaoSqlite struct {
	DB *gorm.DB
}

func (dao *AttributeCodeDaoSqlite) Create(attributeCode model.StoreAttributeCode) error {
	return dao.DB.Save(&attributeCode).Error
}

func (dao *AttributeCodeDaoSqlite) Delete(attributeCodeId int64) error {
	result := dao.DB.Delete(&model.StoreAttributeCode{}, attributeCodeId)
	return result.Error
}

func (dao *AttributeCodeDaoSqlite) GetById(attributeCodeId int64) (model.StoreAttributeCode, error) {
	attributeCode := model.StoreAttributeCode{}
	result := dao.DB.First(&attributeCode, attributeCodeId)
	return attributeCode, result.Error
}

func (dao *AttributeCodeDaoSqlite) GetByTitle(attributeCodeTitle string) ([]model.StoreAttributeCode, error) {
	var attributeCodes []model.StoreAttributeCode
	result := dao.DB.Where("title = ?", attributeCodeTitle).Find(&attributeCodes)
	return attributeCodes, result.Error
}
