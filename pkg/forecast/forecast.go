package forecast

import (
	"probabilisticTimeSeriesModeling/internal/credit"
	"strconv"
	"time"
)

type forecast struct {
}

type Forecast interface {
	ForecastingBankData(bankData []ForecastEl, params credit.ForecastingBankDataRequest) (response *BankForecast, err error)
}

func NewForecast() (obj Forecast, err error) {
	return &forecast{}, err
}

func (f *forecast) ForecastingBankData(bankData []ForecastEl, params credit.ForecastingBankDataRequest) (response *BankForecast, err error) {
	var sum float64
	var yi float64
	var bankForecast BankForecast
	years, err := strconv.Atoi(params.Years)

	if err != nil {
		return
	}
	period := f.FindTimePeriod(bankData[0].Date, bankData[1].Date)
	year := bankData[0].Date

	if len(bankData) < 1 {
		return
	}
	bankDataLen := len(bankData)
	for i := range bankData {
		if i+1 == bankDataLen {
			break
		}
		sum += bankData[i+1].Price - bankData[i].Price
	}
	b := sum / float64(bankDataLen)
	a := bankData[0].Price
	bankForecast.Points = make(map[int][]BankPoint)
	for i := 0; i < bankDataLen+years*4-1; i++ {
		yi = a + b*float64(i)
		if err != nil {
			return
		}
		bankForecast.Points[year.Year()] = append(bankForecast.Points[year.Year()], BankPoint{
			MidPrice: yi,
			Date:     year,
		})
		year = year.Add(period)
	}
	response = &bankForecast
	if b < 0 {
		response.Trend = "down"
	} else {
		response.Trend = "up"
	}
	return
}

func (f *forecast) FindTimePeriod(t1, t2 time.Time) (period time.Duration) {
	period = t2.Sub(t1)
	return
}
