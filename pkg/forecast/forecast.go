package forecast

type forecast struct {
}

type Forecast interface {
	ForecastingBankData(bankData []ForecastEl) (response *BankForecast, err error)
}

func NewForecast() (obj Forecast, err error) {
	return &forecast{}, err
}

func (f *forecast) ForecastingBankData(bankData []ForecastEl) (response *BankForecast, err error) {
	var sum float64
	var changeYear int
	var yi float64
	var bankForecast BankForecast
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
	for i := 0; i < n; i++ {
		bankForecast.Points[yearStart] = make([]BankPoint, 0)
		if i == bankDataLen {
			yi = a + b*float64(i+1)
			bankForecast.Points[yearStart] = append(bankForecast.Points[yearStart], BankPoint{
				MidPrice: yi,
				Date:     bankData[bankDataLen-i-1].Data,
			})
			continue
		}
		yi = a + b*float64(i+1)
		bankForecast.Points[yearStart] = append(bankForecast.Points[yearStart], BankPoint{
			MidPrice: yi,
			Date:     bankData[bankDataLen-i-1].Data,
		})
		if changeYear == 4 {
			changeYear = 0
			yearStart++
		}
		changeYear++
	}
	response = &bankForecast
	return
}
