package sqlite

import (
	"github.com/wizard7414/epos-domain/v2/model"
	"testing"
	"time"
)

func TestEntryDaoSqlite_Create(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	entry := model.StoreEntry{
		ID:      0,
		Title:   "testEntry",
		Url:     "testEntryUrl",
		Source:  0,
		AddDate: time.Now().Unix(),
	}

	createErr := base.Entry.Create(entry)
	if createErr != nil {
		t.Fail()
	}
}

func TestEntryDaoSqlite_GetById(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	entry, findErr := base.Entry.GetById(1)

	if entry.Title != "testEntry" || findErr != nil {
		t.Fail()
	}
}

func TestEntryDaoSqlite_GetByTitle(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	entries, findErr := base.Entry.GetByTitle("testEntry")

	if len(entries) != 1 || entries[0].Title != "testEntry" || findErr != nil {
		t.Fail()
	}
}

func TestEntryDaoSqlite_Delete(t *testing.T) {
	base, err := Init()
	defer base.CloseDb()
	if base == nil || err != nil {
		t.Fail()
	}

	result := base.Entry.Delete(1)

	if result != nil {
		t.Fail()
	}
}
