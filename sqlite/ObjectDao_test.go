package sqlite

import (
	"github.com/wizard7414/epos-domain/v2/model"
	"testing"
	"time"
)

func TestObjectDaoSqlite_Create(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	object := model.StoreObject{
		ID:         1,
		Title:      "testObject",
		Entry:      1,
		Url:        "testObjectUrl",
		AddDate:    time.Now().Unix(),
		ObjectType: 1,
	}

	createErr := base.Object.Create(object)
	if createErr != nil {
		t.Fail()
	}
}

func TestObjectDaoSqlite_GetById(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	object, findErr := base.Object.GetById(1)

	if object.Title != "testObject" || findErr != nil {
		t.Fail()
	}
}

func TestObjectDaoSqlite_GetByTitle(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	objects, findErr := base.Object.GetByTitle("testObject")

	if len(objects) != 1 || objects[0].Title != "testObject" || findErr != nil {
		t.Fail()
	}
}

func TestObjectDaoSqlite_GetByUrl(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	objects, findErr := base.Object.GetByUrl("testObjectUrl")

	if len(objects) != 1 || objects[0].Title != "testObject" || findErr != nil {
		t.Fail()
	}
}

func TestObjectDaoSqlite_Delete(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	result := base.Object.Delete(1)

	if result != nil {
		t.Fail()
	}
}
