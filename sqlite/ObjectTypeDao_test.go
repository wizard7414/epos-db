package sqlite

import (
	"github.com/wizard7414/epos-domain/v2/model"
	"testing"
)

func TestObjectTypeDaoSqlite_Create(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	objectType := model.StoreObjectType{
		ID:    0,
		Title: "testObjectType",
	}

	createErr := base.ObjectType.Create(objectType)
	if createErr != nil {
		t.Fail()
	}
}

func TestObjectTypeDaoSqlite_GetById(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	objectType, findErr := base.ObjectType.GetById(1)

	if objectType.Title != "testObjectType" || findErr != nil {
		t.Fail()
	}
}

func TestObjectTypeDaoSqlite_GetByTitle(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	objectTypes, findErr := base.ObjectType.GetByTitle("testObjectType")

	if len(objectTypes) != 1 || objectTypes[0].Title != "testObjectType" || findErr != nil {
		t.Fail()
	}
}

func TestObjectTypeDaoSqlite_Delete(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	result := base.ObjectType.Delete(1)

	if result != nil {
		t.Fail()
	}
}
