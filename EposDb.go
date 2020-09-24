package epos_db

import (
	"github.com/wizard7414/epos-domain/v2/model"
)

type EposDb interface {
	InitDb(path string) error

	CloseDb() error

	CreateAttribute(attribute model.ViewAttribute) error

	CreateSource(source model.ViewSource) error

	CreateEntry(entry model.ViewEntry) error

	CreateObject(object model.ViewObject) error

	GetObjectAttributes(objectId int64) ([]model.ViewAttribute, error)
}

type AttributeCodeDao interface {
	Create(attributeCode model.StoreAttributeCode) error

	Delete(attributeCodeId int64) error

	GetById(attributeCodeId int64) (model.StoreAttributeCode, error)

	GetByTitle(attributeCodeTitle string) ([]model.StoreAttributeCode, error)
}

type AttributeTypeDao interface {
	Create(attributeType model.StoreAttributeType) error

	Delete(attributeTypeId int64) error

	GetById(attributeTypeId int64) (model.StoreAttributeType, error)

	GetByTitle(attributeTypeTitle string) ([]model.StoreAttributeType, error)
}

type AttributeDao interface {
	Create(attribute model.StoreAttribute) error

	Delete(attributeId int64) error

	GetById(attributeId int64) (model.StoreAttribute, error)

	GetByCode(attributeCode int64) ([]model.StoreAttribute, error)
}

type ResourceDao interface {
	Create(resource model.StoreResource) error

	Delete(resourceId int64) error

	GetById(resourceId int64) (model.StoreResource, error)

	GetByTitle(resourceTitle string) ([]model.StoreResource, error)
}

type SourceDao interface {
	Create(source model.StoreSource) error

	Delete(sourceId int64) error

	GetById(sourceId int64) (model.StoreSource, error)

	GetByTitle(sourceTitle string) ([]model.StoreSource, error)

	UpdateDateById(sourceId int64, newDate int64) error

	UpdateDateByTitle(sourceTitle string, newDate int64) error
}

type EntryDao interface {
	Create(entry model.StoreEntry) error

	Delete(entryId int64) error

	GetById(entryId int64) (model.StoreEntry, error)

	GetByTitle(entryTitle string) ([]model.StoreEntry, error)
}

type ObjectTypeDao interface {
	Create(objectType model.StoreObjectType) error

	Delete(objectTypeId int64) error

	GetById(objectTypeId int64) (model.StoreObjectType, error)

	GetByTitle(objectTypeTitle string) ([]model.StoreObjectType, error)
}

type ObjectDao interface {
	Create(object model.StoreObject) error

	Delete(objectId int64) error

	GetById(objectId int64) (model.StoreObject, error)

	GetByTitle(objectTitle string) ([]model.StoreObject, error)
}

type ObjectAttributeDao interface {
	Create(objectAttribute model.StoreObjectAttribute) error

	Delete(objectId int64, attributeId int64) error

	GetById(objectId int64, attributeId int64) (model.StoreObjectAttribute, error)

	GetByObject(objectId int64) ([]model.StoreObjectAttribute, error)

	GetByAttribute(attributeId int64) ([]model.StoreObjectAttribute, error)
}
