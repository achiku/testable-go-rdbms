package model

import (
	"testing"
	"time"

	"github.com/AdrianLungu/decimal"
)

func TestSale_GetDailySummary(t *testing.T) {
	tx, cleanup := TestSetupTx(t)
	defer cleanup()

	dt := time.Now()
	u := TestCreateUserAccountData(t, tx, &UserAccount{})
	i1 := TestCreateItemData(t, tx, &Item{
		Name:  "beer",
		Price: decimal.NewFromFloat(500),
	})
	TestCreateSaleData(t, tx, i1, u, &Sale{
		SoldAt: dt,
	})
	i2 := TestCreateItemData(t, tx, &Item{
		Name:  "pizza",
		Price: decimal.NewFromFloat(1200),
	})
	TestCreateSaleData(t, tx, i2, u, &Sale{
		SoldAt: dt,
	})
	TestCreateSaleData(t, tx, i2, u, &Sale{
		SoldAt: dt.AddDate(0, 0, -2),
	})

	sts, err := GetDailySummary(tx, dt)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", sts)
}
