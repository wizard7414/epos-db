package sqlite

import (
	"github.com/wizard7414/epos-domain/v2/model"
	"testing"
	"time"
)

func TestSourceDaoSqlite_Create(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	source := model.StoreSource{
		ID:         0,
		Title:      "testSource",
		Url:        "testSourceUrl",
		Resource:   0,
		ChangeDate: time.Now().Unix(),
	}

	createErr := base.Source.Create(source)
	if createErr != nil {
		t.Fail()
	}
}

func TestSourceDaoSqlite_GetById(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	source, findErr := base.Source.GetById(1600682767129601576)

	if source.Title != "testSource" || findErr != nil {
		t.Fail()
	}
}

func TestSourceDaoSqlite_GetByTitle(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	sourceList, findErr := base.Source.GetByTitle("testSource")

	if len(sourceList) != 1 || sourceList[0].Title != "testSource" || findErr != nil {
		t.Fail()
	}
}

func TestSourceDaoSqlite_UpdateDateById(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	updateErr := base.Source.UpdateDateById(1600682767129601576, time.Now().Unix())

	if updateErr != nil {
		t.Fail()
	}
}

func TestSourceDaoSqlite_UpdateDateByTitle(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	updateErr := base.Source.UpdateDateByTitle("testSource", time.Now().Unix())

	if updateErr != nil {
		t.Fail()
	}
}

func TestSourceDaoSqlite_Delete(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	result := base.Source.Delete(1)

	if result != nil {
		t.Fail()
	}
}
