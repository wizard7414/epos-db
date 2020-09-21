package sqlite

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos-domain/v2/model"
)

type EntryDaoSqlite struct {
	DB *gorm.DB
}

func (dao *EntryDaoSqlite) Create(entry model.StoreEntry) error {
	return dao.DB.Save(&entry).Error
}

func (dao *EntryDaoSqlite) Delete(entryId int64) error {
	result := dao.DB.Delete(&model.StoreEntry{}, entryId)
	return result.Error
}

func (dao *EntryDaoSqlite) GetById(entryId int64) (model.StoreEntry, error) {
	entry := model.StoreEntry{}
	result := dao.DB.First(&entry, entryId)
	return entry, result.Error
}

func (dao *EntryDaoSqlite) GetByTitle(entryTitle string) ([]model.StoreEntry, error) {
	var entries []model.StoreEntry
	result := dao.DB.Where("title = ?", entryTitle).Find(&entries)
	return entries, result.Error
}
