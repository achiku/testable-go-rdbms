package model

import (
	"log"
	"testing"

	"github.com/achiku/mergo"
)

// TestCreateSaleData create sale test data
func TestCreateSaleData(t *testing.T, tx Query, item *Item, ua *UserAccount, sale *Sale) *Sale {
	saleDefault := &Sale{
		AccountID:  ua.ID,
		ItemID:     item.ID,
		PaidAmount: item.Price,
		SoldAt:     sale.SoldAt,
	}
	if err := mergo.MergeWithOverwrite(saleDefault, sale, TestStructMergeFunc); err != nil {
		log.Fatal(err)
	}
	if err := saleDefault.Create(tx); err != nil {
		t.Fatal(err)
	}
	return saleDefault
}
