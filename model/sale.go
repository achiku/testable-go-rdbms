package model

import (
	"time"

	"github.com/AdrianLungu/decimal"
)

// DailySummaryStats daily summary
type DailySummaryStats struct {
	Date       time.Time
	ItemID     int64
	ItemName   string
	SaleAmount decimal.Decimal
}

// GetDailySummary get daily summary
func GetDailySummary(tx Query, d time.Time) ([]DailySummaryStats, error) {
	from := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local)
	to := from.AddDate(0, 0, 1)
	rows, err := tx.Query(`
	SELECT
		date_trunc('day', s.sold_at)
		, i.id
		, i.name
		, sum(s.paid_amount)
	FROM sale s
	JOIN item i
	ON s.item_id = i.id
	WHERE s.sold_at >= $1
	AND s.sold_at < $2
	GROUP BY
		date_trunc('day', s.sold_at)
		, i.id
		, i.name
	`, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sts []DailySummaryStats
	for rows.Next() {
		var st DailySummaryStats
		rows.Scan(
			&st.Date,
			&st.ItemID,
			&st.ItemName,
			&st.SaleAmount,
		)
		sts = append(sts, st)
	}
	return sts, nil
}
