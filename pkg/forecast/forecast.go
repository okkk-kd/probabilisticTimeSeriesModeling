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
	if len(bankData) < 1 {
		return
	}
	var sum float64
	bankDataLen := len(bankData)
	for i := range bankData {
		if i+1 == bankDataLen {
			break
		}
		sum += bankData[i].Price - bankData[i+1].Price
	}
	b := sum / float64(bankDataLen)
	a := bankData[bankDataLen-1].Price
	var yi float64
	var bankForecast BankForecast
	bankForecast.Points = make([]BankPoint, 1)
	for i := bankDataLen - 1; i > 0; i-- {
		yi = a + b*float64(i) + 1
		bankForecast.Points = append(bankForecast.Points, BankPoint{
			MidPrice: yi,
			Date:     bankData[bankDataLen-i-1].Data,
		})
	}
	//for i, el := range bankData {
	//
	//}
	response = &bankForecast
	return
}
