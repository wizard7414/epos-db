package sqlite

import (
	"github.com/wizard7414/epos-domain/v2/model"
	"testing"
)

func TestAttributeTypeDaoSqlite_Create(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	attributeType := model.StoreAttributeType{
		ID:    0,
		Title: "testType",
	}

	createErr := base.AttributeType.Create(attributeType)
	if createErr != nil {
		t.Fail()
	}
}

func TestAttributeTypeDaoSqlite_GetById(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	attributeType, findErr := base.AttributeType.GetById(1)

	if attributeType.Title != "testType" || findErr != nil {
		t.Fail()
	}
}

func TestAttributeTypeDaoSqlite_GetByTitle(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	attributeTypeList, findErr := base.AttributeType.GetByTitle("testType")

	if len(attributeTypeList) != 1 || attributeTypeList[0].Title != "testType" || findErr != nil {
		t.Fail()
	}
}

func TestAttributeTypeDaoSqlite_Delete(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	result := base.AttributeType.Delete(1)

	if result != nil {
		t.Fail()
	}
}
