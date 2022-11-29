package forecast

import "time"

type BankForecast struct {
	Points map[int][]BankPoint
}

type BankPoint struct {
	MidPrice float64
	Date     time.Time
}

type ForecastEl struct {
	Data  string
	Price float64
}
