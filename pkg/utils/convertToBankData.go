package utils

import (
	"fmt"
	"probabilisticTimeSeriesModeling/pkg/forecast"
	"strconv"
)

func ConvertToBankData(before []interface{}) (result forecast.ForecastEl, err error) {
	result.Data = fmt.Sprintf("%v", before[0])
	price, err := strconv.ParseFloat(fmt.Sprintf("%v", before[1]), 64)
	if err != nil {
		return
	}
	result.Price = price
	return
}
