package sqlite

import (
	"github.com/wizard7414/epos-domain/v2/model"
	"testing"
)

func TestAttributeCodeDaoSqlite_Create(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	attributeCode := model.StoreAttributeCode{
		ID:    0,
		Title: "testCode",
	}

	createErr := base.AttributeCode.Create(attributeCode)
	if createErr != nil {
		t.Fail()
	}
}

func TestAttributeCodeDaoSqlite_GetById(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	attributeCode, findErr := base.AttributeCode.GetById(1)

	if attributeCode.Title != "testCode" || findErr != nil {
		t.Fail()
	}
}

func TestAttributeCodeDaoSqlite_GetByTitle(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	attributeCodeList, findErr := base.AttributeCode.GetByTitle("testCode")

	if len(attributeCodeList) != 1 || attributeCodeList[0].Title != "testCode" || findErr != nil {
		t.Fail()
	}
}

func TestAttributeCodeDaoSqlite_Delete(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	result := base.AttributeCode.Delete(1)

	if result != nil {
		t.Fail()
	}
}
