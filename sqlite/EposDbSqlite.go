package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wizard7414/epos-domain/v2/model"
	"time"
)

type EposDbSqlite struct {
	AttributeCode   AttributeCodeDaoSqlite
	AttributeType   AttributeTypeDaoSqlite
	Attribute       AttributeDaoSqlite
	Resource        ResourceDaoSqlite
	Source          SourceDaoSqlite
	Entry           EntryDaoSqlite
	ObjectType      ObjectTypeDaoSqlite
	Object          ObjectDaoSqlite
	ObjectAttribute ObjectAttributeDaoSqlite
}

func (db *EposDbSqlite) InitDb(path string) error {
	dataBase, openErr := gorm.Open("sqlite3", path)
	if openErr == nil {
		db.AttributeCode.DB = dataBase
		db.AttributeType.DB = dataBase
		db.Attribute.DB = dataBase
		db.Resource.DB = dataBase
		db.Source.DB = dataBase
		db.Entry.DB = dataBase
		db.ObjectType.DB = dataBase
		db.Object.DB = dataBase
		db.ObjectAttribute.DB = dataBase
		return nil
	} else {
		return openErr
	}
}

func (db *EposDbSqlite) CloseDb() error {
	return db.AttributeCode.DB.Close()
}

func (db *EposDbSqlite) CreateAttribute(attribute model.ViewAttribute) error {
	attributeCode, codeErr := createOrGetAttributeCode(db, attribute.Code)
	if codeErr != nil {
		return codeErr
	}
	attributeType, typeErr := createOrGetAttributeType(db, attribute.AttrType)
	if typeErr != nil {
		return typeErr
	}
	if attribute.ID == 0 {
		attribute.ID = time.Now().UnixNano()
	}
	attribute.Code.ID = attributeCode.ID
	attribute.AttrType.ID = attributeType.ID
	return db.Attribute.Create(attribute.TransformToStore())
}

func (db *EposDbSqlite) CreateSource(source model.ViewSource) error {
	resource, resourceErr := createOrGetResource(db, source.Resource)
	if resourceErr != nil {
		return resourceErr
	}
	if source.ID == 0 {
		source.ID = time.Now().UnixNano()
	}
	source.Resource.ID = resource.ID
	return db.Source.Create(source.TransformToStore())
}

func (db *EposDbSqlite) CreateEntry(entry model.ViewEntry) error {
	source, sourceErr := createOrGetSource(db, entry.Source)
	if sourceErr != nil {
		return sourceErr
	}
	if entry.ID == 0 {
		entry.ID = time.Now().UnixNano()
	}
	entry.Source.ID = source.ID
	updateErr := db.Source.UpdateDateById(source.ID, time.Now().Unix())
	if updateErr != nil {
		return updateErr
	}
	return db.Entry.Create(entry.TransformToStore())
}

func (db *EposDbSqlite) CreateObject(object model.ViewObject) error {
	objectType, objectTypeErr := createOrGetObjectType(db, object.ObjectType)
	if objectTypeErr != nil {
		return objectTypeErr
	}
	if object.Entry.ID == 0 {
		object.Entry.ID = time.Now().UnixNano()
	}
	entryErr := db.CreateEntry(object.Entry)
	if entryErr != nil {
		return entryErr
	}
	if object.ID == 0 {
		object.ID = time.Now().UnixNano()
	}
	if len(object.Attributes) > 0 {
		for _, attribute := range object.Attributes {
			if attribute.ID == 0 {
				attribute.ID = time.Now().UnixNano()
			}
			attributeErr := db.CreateAttribute(attribute)
			if attributeErr != nil {
				return attributeErr
			}
			updatedObjectAttribute := model.StoreObjectAttribute{
				Object:    object.ID,
				Attribute: attribute.ID,
			}
			objectAttributeErr := db.ObjectAttribute.Create(updatedObjectAttribute)
			if objectAttributeErr != nil {
				return objectAttributeErr
			}
		}
	}
	object.ObjectType.ID = objectType.ID
	return db.Object.Create(object.TransformToStore())
}

func (db *EposDbSqlite) GetObjectAttributes(objectId int64) ([]model.ViewAttribute, error) {
	attributes, findErr := db.ObjectAttribute.GetByObject(objectId)
	var result []model.ViewAttribute
	if findErr != nil {
		return result, findErr
	}

	for index := range attributes {
		attribute, attributeErr := db.Attribute.GetById(attributes[index].Attribute)
		if attributeErr != nil {
			return result, attributeErr
		}

		code, codeErr := db.AttributeCode.GetById(attribute.Code)
		if codeErr != nil {
			return result, codeErr
		}

		attributeType, typeErr := db.AttributeType.GetById(attribute.AttributeType)
		if typeErr != nil {
			return result, typeErr
		}

		resultAttribute := attribute.TransformToView()
		resultAttribute.Code = code.TransformToView()
		resultAttribute.AttrType = attributeType.TransformToView()

		result = append(result, resultAttribute)
	}

	return result, nil
}

func (db *EposDbSqlite) GetObjectByUrl(objectUrl string) ([]model.ViewObject, error) {
	objects, findErr := db.Object.GetByUrl(objectUrl)
	var result []model.ViewObject
	if findErr != nil {
		return result, findErr
	}

	for index := range objects {
		resultObject := objects[index].TransformToView()

		objectType, objectTypeErr := db.ObjectType.GetById(objects[index].ObjectType)
		if objectTypeErr != nil {
			return result, objectTypeErr
		}
		resultObject.ObjectType = objectType.TransformToView()

		entry, entryErr := db.Entry.GetById(resultObject.Entry.ID)
		if entryErr != nil {
			return result, entryErr
		}
		resultObject.Entry = entry.TransformToView()

		entrySource, entrySourceErr := db.Source.GetById(resultObject.Entry.Source.ID)
		if entrySourceErr != nil {
			return result, entrySourceErr
		}
		resultObject.Entry.Source = entrySource.TransformToView()

		sourceResource, sourceResourceErr := db.Resource.GetById(resultObject.Entry.Source.Resource.ID)
		if sourceResourceErr != nil {
			return result, sourceResourceErr
		}
		resultObject.Entry.Source.Resource = sourceResource.TransformToView()

		attributes, attributesErr := db.GetObjectAttributes(resultObject.ID)
		if attributesErr != nil {
			return result, attributesErr
		}
		resultObject.Attributes = attributes

		result = append(result, resultObject)
	}

	return result, nil
}
