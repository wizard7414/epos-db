package sqlite

import (
	"github.com/wizard7414/epos-domain/v2/model"
	"testing"
)

func TestResourceDaoSqlite_Create(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	resource := model.StoreResource{
		ID:    0,
		Title: "testResource",
		Url:   "testResourceUrl",
	}

	createErr := base.Resource.Create(resource)
	if createErr != nil {
		t.Fail()
	}
}

func TestResourceDaoSqlite_GetById(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	resource, findErr := base.Resource.GetById(1)

	if resource.Title != "testResource" || findErr != nil {
		t.Fail()
	}
}

func TestResourceDaoSqlite_GetByTitle(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	resourceList, findErr := base.Resource.GetByTitle("testResource")

	if len(resourceList) != 1 || resourceList[0].Title != "testResource" || findErr != nil {
		t.Fail()
	}
}

func TestResourceDaoSqlite_Delete(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	result := base.Resource.Delete(1)

	if result != nil {
		t.Fail()
	}
}
