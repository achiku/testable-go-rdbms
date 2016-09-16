package model

import (
	"database/sql"
	"log"
	"testing"

	"github.com/AdrianLungu/decimal"
	"github.com/achiku/mergo"
)

// TestCreateItemData create sale test data
func TestCreateItemData(t *testing.T, tx Query, item *Item) *Item {
	itemDefault := &Item{
		Price:       decimal.NewFromFloat(1000),
		Name:        "item1",
		Description: sql.NullString{String: "test desc"},
	}
	if err := mergo.MergeWithOverwrite(itemDefault, item, TestStructMergeFunc); err != nil {
		log.Fatal(err)
	}
	if err := itemDefault.Create(tx); err != nil {
		t.Fatal(err)
	}
	return itemDefault
}
