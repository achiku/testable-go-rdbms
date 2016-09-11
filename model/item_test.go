package model

import (
	"database/sql"
	"testing"

	"github.com/AdrianLungu/decimal"
)

func TestItem_ItemCreate(t *testing.T) {
	tx, cleanup := TestSetupTx(t)
	defer cleanup()

	item := Item{
		Name:        "test item1",
		Price:       decimal.NewFromFloat(1000),
		Description: sql.NullString{String: "this is test"},
	}
	if err := item.Create(tx); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", item)

	targetItem, err := GetItemByPk(tx, item.ID)
	if err != nil {
		t.Fatal(err)
	}
	if targetItem.Name != item.Name {
		t.Errorf("want %s got %s", item.Name, targetItem.Name)
	}
}
