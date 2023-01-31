package forecast

import "time"

type BankForecast struct {
	Points map[int][]BankPoint
	Trend  string
}

type BankPoint struct {
	ID       int `db:"id"`
	MidPrice float64
	Date     time.Time
}

type ForecastEl struct {
	ID    int       `db:"id"`
	Date  time.Time `db:"date"`
	Price float64   `db:"amount"`
}
