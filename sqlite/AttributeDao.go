package sqlite

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos-domain/v2/model"
)

type AttributeDaoSqlite struct {
	DB *gorm.DB
}

func (dao *AttributeDaoSqlite) Create(attribute model.StoreAttribute) error {
	return dao.DB.Save(&attribute).Error
}

func (dao *AttributeDaoSqlite) Delete(attributeId int64) error {
	result := dao.DB.Delete(&model.StoreAttribute{}, attributeId)
	return result.Error
}

func (dao *AttributeDaoSqlite) GetById(attributeId int64) (model.StoreAttribute, error) {
	attribute := model.StoreAttribute{}
	result := dao.DB.First(&attribute, attributeId)
	return attribute, result.Error
}

func (dao *AttributeDaoSqlite) GetByCode(attributeCode int64) ([]model.StoreAttribute, error) {
	var attributes []model.StoreAttribute
	result := dao.DB.Where("code = ?", attributeCode).Find(&attributes)
	return attributes, result.Error
}
