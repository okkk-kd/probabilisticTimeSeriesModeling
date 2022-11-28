package forecast

type BankForecast struct {
	Points []BankPoint
}

type BankPoint struct {
	MidPrice float64
	Date     string
}

type ForecastEl struct {
	Data  string
	Price float64
}
