package sqlite

import (
	"github.com/wizard7414/epos-domain/v2/model"
	"time"
)

func createOrGetAttributeCode(db *EposDbSqlite, attributeCode model.ViewAttributeCode) (model.StoreAttributeCode, error) {
	attributeCodes, findErr := db.AttributeCode.GetByTitle(attributeCode.Title)
	var updatedAttributeCode model.StoreAttributeCode

	if findErr != nil {
		return updatedAttributeCode, findErr
	} else if len(attributeCodes) == 0 {
		attributeCode.ID = time.Now().UnixNano()
		createErr := db.AttributeCode.Create(attributeCode.TransformToStore())
		if createErr != nil {
			return updatedAttributeCode, createErr
		} else {
			updatedAttributeCode = attributeCode.TransformToStore()
		}
	} else if len(attributeCodes) == 1 {
		updatedAttributeCode = attributeCodes[0]
	}
	return updatedAttributeCode, nil
}

func createOrGetAttributeType(db *EposDbSqlite, attributeType model.ViewAttributeType) (model.StoreAttributeType, error) {
	attributeTypes, findErr := db.AttributeType.GetByTitle(attributeType.Title)
	var updatedAttributeType model.StoreAttributeType

	if findErr != nil {
		return updatedAttributeType, findErr
	} else if len(attributeTypes) == 0 {
		attributeType.ID = time.Now().UnixNano()
		createErr := db.AttributeType.Create(attributeType.TransformToStore())
		if createErr != nil {
			return updatedAttributeType, createErr
		} else {
			updatedAttributeType = attributeType.TransformToStore()
		}
	} else if len(attributeTypes) == 1 {
		updatedAttributeType = attributeTypes[0]
	}
	return updatedAttributeType, nil
}

func createOrGetResource(db *EposDbSqlite, resource model.ViewResource) (model.StoreResource, error) {
	resources, findErr := db.Resource.GetByTitle(resource.Title)
	var updatedResource model.StoreResource

	if findErr != nil {
		return updatedResource, findErr
	} else if len(resources) == 0 {
		resource.ID = time.Now().UnixNano()
		createErr := db.Resource.Create(resource.TransformToStore())
		if createErr != nil {
			return updatedResource, createErr
		} else {
			updatedResource = resource.TransformToStore()
		}
	} else if len(resources) == 1 {
		updatedResource = resources[0]
	}
	return updatedResource, nil
}

func createOrGetSource(db *EposDbSqlite, source model.ViewSource) (model.StoreSource, error) {
	sources, findErr := db.Source.GetByTitle(source.Title)
	var updatedSource model.StoreSource

	if findErr != nil {
		return updatedSource, findErr
	} else if len(sources) == 0 {
		source.ID = time.Now().UnixNano()
		createErr := db.CreateSource(source)
		if createErr != nil {
			return updatedSource, createErr
		} else {
			updatedSource = source.TransformToStore()
		}
	} else if len(sources) == 1 {
		updatedSource = sources[0]
	}
	return updatedSource, nil
}

func createOrGetObjectType(db *EposDbSqlite, objectType model.ViewObjectType) (model.StoreObjectType, error) {
	objectTypes, findErr := db.ObjectType.GetByTitle(objectType.Title)
	var updateObjectType model.StoreObjectType

	if findErr != nil {
		return updateObjectType, findErr
	} else if len(objectTypes) == 0 {
		objectType.ID = time.Now().UnixNano()
		createErr := db.ObjectType.Create(objectType.TransformToStore())
		if createErr != nil {
			return updateObjectType, createErr
		} else {
			updateObjectType = objectType.TransformToStore()
		}
	} else if len(objectTypes) == 1 {
		updateObjectType = objectTypes[0]
	}
	return updateObjectType, nil
}
