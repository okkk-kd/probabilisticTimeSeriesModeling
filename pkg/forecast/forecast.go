package forecast

import (
	"fmt"
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
	var changeYear int
	var yi float64
	var bankForecast BankForecast
	var timeDate time.Time
	years, err := strconv.Atoi(params.Years)
	if err != nil {
		return
	}
	yearStart := 2010

	if len(bankData) < 1 {
		return
	}
	bankDataLen := len(bankData)
	for i := range bankData {
		if i+1 == bankDataLen {
			break
		}
		sum += bankData[i].Price - bankData[i+1].Price
	}
	b := sum / float64(bankDataLen)
	a := bankData[bankDataLen-1].Price
	bankForecast.Points = make(map[int][]BankPoint)
	changeYear = 1
	for i := 0; i < bankDataLen+years*4-1; i++ {
		if changeYear == 4 {
			changeYear = 0
			yearStart++
			bankForecast.Points[yearStart] = make([]BankPoint, 0)
		}
		if i >= bankDataLen {
			yi = a + b*float64(i)
			timeDate, err = f.MakeDate(yearStart, changeYear)
			if err != nil {
				return
			}
			bankForecast.Points[yearStart] = append(bankForecast.Points[yearStart], BankPoint{
				MidPrice: yi,
				Date:     timeDate,
			})
		} else {
			yi = a + b*float64(i)
			timeDate, err = f.MakeDate(yearStart, changeYear)
			if err != nil {
				return
			}
			bankForecast.Points[yearStart] = append(bankForecast.Points[yearStart], BankPoint{
				MidPrice: yi,
				Date:     timeDate,
			})
		}
		changeYear++
	}
	response = &bankForecast
	return
}

func (f *forecast) MakeDate(year, level int) (str time.Time, err error) {
	switch level {
	case 0:
		str, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-31", year))
		if err != nil {
			return
		}
	case 1:
		str, err = time.Parse("2006-01-02", fmt.Sprintf("%d-06-30", year))
		if err != nil {
			return
		}
	case 2:
		str, err = time.Parse("2006-01-02", fmt.Sprintf("%d-09-30", year))
		if err != nil {
			return
		}
	case 3:
		str, err = time.Parse("2006-01-02", fmt.Sprintf("%d-12-31", year))
		if err != nil {
			return
		}
	}

	return
}
