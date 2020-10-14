package sqlite

import (
	"github.com/wizard7414/epos-domain/v2/model"
	"testing"
	"time"
)

func Init() (*EposDbSqlite, error) {
	base := EposDbSqlite{
		AttributeType: AttributeTypeDaoSqlite{
			DB: nil,
		},
		AttributeCode: AttributeCodeDaoSqlite{
			DB: nil,
		},
		Attribute: AttributeDaoSqlite{
			DB: nil,
		},
		Resource: ResourceDaoSqlite{
			DB: nil,
		},
		Source: SourceDaoSqlite{
			DB: nil,
		},
		Entry: EntryDaoSqlite{
			DB: nil,
		},
		ObjectType: ObjectTypeDaoSqlite{
			DB: nil,
		},
		Object: ObjectDaoSqlite{
			DB: nil,
		},
		ObjectAttribute: ObjectAttributeDaoSqlite{
			DB: nil,
		},
	}
	err := base.InitDb("epos-miner.db")
	return &base, err
}

func TestInit(t *testing.T) {
	var base *EposDbSqlite
	var err error
	base, err = Init()
	defer base.CloseDb()
	if err != nil {
		t.Fail()
	}
}

func TestEposDbSqlite_CreateAttribute(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if err != nil {
		t.Fail()
	}

	testAttribute := model.ViewAttribute{
		ID: 0,
		Code: model.ViewAttributeCode{
			ID:    0,
			Title: "testCode",
		},
		AttrType: model.ViewAttributeType{
			ID:    0,
			Title: "testType",
		},
		Value: "testValue",
	}

	createErr := base.CreateAttribute(testAttribute)
	if createErr != nil {
		t.Fail()
	}
}

func TestEposDbSqlite_CreateSource(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if err != nil {
		t.Fail()
	}

	testSource := model.ViewSource{
		ID:    0,
		Title: "Test Source",
		Url:   "Test source URL",
		Resource: model.ViewResource{
			ID:    0,
			Title: "Test Resource",
			Url:   "Test resource URL",
		},
		ChangeDate: time.Now(),
	}

	createErr := base.CreateSource(testSource)
	if createErr != nil {
		t.Fail()
	}
}

func TestEposDbSqlite_CreateEntry(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if err != nil {
		t.Fail()
	}

	testEntry := model.ViewEntry{
		ID:    0,
		Title: "Test Entry",
		Url:   "Test entry Url",
		Source: model.ViewSource{
			ID:    0,
			Title: "Test Source",
			Url:   "Test source URL",
			Resource: model.ViewResource{
				ID:    0,
				Title: "Test Resource",
				Url:   "Test resource URL",
			},
			ChangeDate: time.Now(),
		},
		AddDate: time.Time{},
	}

	createErr := base.CreateEntry(testEntry)
	if createErr != nil {
		t.Fail()
	}
}

func TestEposDbSqlite_CreateObject(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if err != nil {
		t.Fail()
	}

	var attributes []model.ViewAttribute

	attributes = append(attributes, model.ViewAttribute{
		ID: 0,
		Code: model.ViewAttributeCode{
			ID:    0,
			Title: "testCode",
		},
		AttrType: model.ViewAttributeType{
			ID:    0,
			Title: "testType",
		},
		Value: "testValue",
	})

	testObject := model.ViewObject{
		ID:    0,
		Title: "Test Object",
		Entry: model.ViewEntry{
			ID:    0,
			Title: "Test Entry",
			Url:   "Test entry Url",
			Source: model.ViewSource{
				ID:    0,
				Title: "Test Source",
				Url:   "Test source URL",
				Resource: model.ViewResource{
					ID:    0,
					Title: "Test Resource",
					Url:   "Test resource URL",
				},
				ChangeDate: time.Now(),
			},
			AddDate: time.Time{},
		},
		Url:     "Test object Url",
		AddDate: time.Now(),
		ObjectType: model.ViewObjectType{
			ID:    0,
			Title: "Test Object Type",
		},
		Attributes: attributes,
	}

	createErr := base.CreateObject(testObject)
	if createErr != nil {
		t.Fail()
	}
}

func TestEposDbSqlite_GetObjectAttributes(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if err != nil {
		t.Fail()
	}

	attributes, findError := base.GetObjectAttributes(1600694962558158355)
	if findError != nil || len(attributes) != 1 {
		t.Fail()
	}
}

func TestEposDbSqlite_GetObjectByUrl(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if err != nil {
		t.Fail()
	}

	objects, findErr := base.GetObjectByUrl("Test object Url")
	if findErr != nil || len(objects) != 1 {
		t.Fail()
	}
}
