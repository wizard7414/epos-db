package sqlite

import (
	"github.com/wizard7414/epos-domain/v2/model"
	"testing"
)

func TestObjectAttributeDaoSqlite_Create(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	objectAttribute := model.StoreObjectAttribute{
		Object:    0,
		Attribute: 0,
	}

	createErr := base.ObjectAttribute.Create(objectAttribute)
	if createErr != nil {
		t.Fail()
	}
}

func TestObjectAttributeDaoSqlite_GetById(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	objectAttribute, findErr := base.ObjectAttribute.GetById(0, 0)

	if objectAttribute.Attribute != 0 || objectAttribute.Object != 0 || findErr != nil {
		t.Fail()
	}
}

func TestObjectAttributeDaoSqlite_GetByObject(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	objectAttributes, findErr := base.ObjectAttribute.GetByObject(0)

	if len(objectAttributes) != 1 || objectAttributes[0].Object != 0 || findErr != nil {
		t.Fail()
	}
}

func TestObjectAttributeDaoSqlite_GetByAttribute(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	objectAttributes, findErr := base.ObjectAttribute.GetByAttribute(0)

	if len(objectAttributes) != 1 || objectAttributes[0].Attribute != 0 || findErr != nil {
		t.Fail()
	}
}

func TestObjectAttributeDaoSqlite_Delete(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	result := base.ObjectAttribute.Delete(0, 0)

	if result != nil {
		t.Fail()
	}
}
