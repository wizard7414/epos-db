package sqlite

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos-domain/v2/model"
)

type SourceDaoSqlite struct {
	DB *gorm.DB
}

func (dao *SourceDaoSqlite) Create(source model.StoreSource) error {
	return dao.DB.Save(&source).Error
}

func (dao *SourceDaoSqlite) Delete(sourceId int64) error {
	result := dao.DB.Delete(&model.StoreSource{}, sourceId)
	return result.Error
}

func (dao *SourceDaoSqlite) GetById(sourceId int64) (model.StoreSource, error) {
	source := model.StoreSource{}
	result := dao.DB.First(&source, sourceId)
	return source, result.Error
}

func (dao *SourceDaoSqlite) GetByTitle(sourceTitle string) ([]model.StoreSource, error) {
	var sources []model.StoreSource
	result := dao.DB.Where("title = ?", sourceTitle).Find(&sources)
	return sources, result.Error
}

func (dao *SourceDaoSqlite) UpdateDateById(sourceId int64, newDate int64) error {
	result := dao.DB.Model(model.StoreSource{}).Where("id = ?", sourceId).Update("change_date", newDate)
	return result.Error
}

func (dao *SourceDaoSqlite) UpdateDateByTitle(sourceTitle string, newDate int64) error {
	result := dao.DB.Model(model.StoreSource{}).Where("title = ?", sourceTitle).Update("change_date", newDate)
	return result.Error
}
