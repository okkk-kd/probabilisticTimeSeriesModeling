package utils

import (
	"fmt"
	"probabilisticTimeSeriesModeling/pkg/forecast"
	"strconv"
	"time"
)

func ConvertToBankData(before []interface{}) (result forecast.ForecastEl, err error) {
	result.Date, err = time.Parse("2006-01-02", fmt.Sprintf("%v", before[0]))
	if err != nil {
		return
	}
	price, err := strconv.ParseFloat(fmt.Sprintf("%v", before[1]), 64)
	if err != nil {
		return
	}
	result.Price = price
	return
}
