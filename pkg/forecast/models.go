package forecast

import "time"

type BankForecast struct {
	Points map[int][]BankPoint
	Trend  string
}

type BankPoint struct {
	MidPrice float64
	Date     time.Time
}

type ForecastEl struct {
	Date  time.Time `db:"date"`
	Price float64   `db:"amount"`
}
