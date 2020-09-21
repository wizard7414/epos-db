package sqlite

import (
	"github.com/wizard7414/epos-domain/v2/model"
	"testing"
)

func TestAttributeDaoSqlite_Create(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	attribute := model.StoreAttribute{
		ID:            0,
		Code:          0,
		AttributeType: 0,
		Value:         "testValue",
	}

	createErr := base.Attribute.Create(attribute)
	if createErr != nil {
		t.Fail()
	}
}

func TestAttributeDaoSqlite_GetById(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	attribute, findErr := base.Attribute.GetById(1)

	if attribute.Value != "testValue" || findErr != nil {
		t.Fail()
	}
}

func TestAttributeDaoSqlite_GetByCode(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	attributeList, findErr := base.Attribute.GetByCode(0)

	if len(attributeList) != 1 || attributeList[0].Value != "testValue" || findErr != nil {
		t.Fail()
	}
}

func TestAttributeDaoSqlite_Delete(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	result := base.Attribute.Delete(1)

	if result != nil {
		t.Fail()
	}
}
