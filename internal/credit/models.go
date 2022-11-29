package credit

import "time"

type Dataset struct {
	Datasett struct {
		ID                  int             `json:"id"`
		DatasetCode         string          `json:"dataset_code"`
		DatabaseCode        string          `json:"database_code"`
		Name                string          `json:"name"`
		Description         string          `json:"description"`
		RefreshedAt         time.Time       `json:"refreshed_at"`
		NewestAvailableDate string          `json:"newest_available_date"`
		OldestAvailableDate string          `json:"oldest_available_date"`
		ColumnNames         []string        `json:"column_names"`
		Frequency           string          `json:"frequency"`
		Type                string          `json:"type"`
		Premium             bool            `json:"premium"`
		Limit               interface{}     `json:"limit"`
		Transform           interface{}     `json:"transform"`
		ColumnIndex         interface{}     `json:"column_index"`
		StartDate           string          `json:"start_date"`
		EndDate             string          `json:"end_date"`
		Data                [][]interface{} `json:"data"`
		Collapse            interface{}     `json:"collapse"`
		Order               interface{}     `json:"order"`
		DatabaseID          int             `json:"database_id"`
	} `json:"dataset"`
}

type ForecastingBankDataRequest struct {
	Years string
	Code  string
}

type GetCodesListResponse struct {
	Name string `db:"name"`
	Code string `db:"code"`
}

type GetCodeDataByID struct {
	ID   string
	Code string
}

type DeleteCodeDataByID struct {
	ID   string
	Code string
}

type UpdateCodeDataByID struct {
	ID     string
	Code   string
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
}

type AddCodeData struct {
	Code   string
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
}

type GetCodeDataByIDResponse struct {
	Amount float64   `db:"amount"`
	Date   time.Time `db:"date"`
}
