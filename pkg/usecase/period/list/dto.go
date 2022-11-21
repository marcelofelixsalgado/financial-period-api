package list

type InputListPeriodDto struct {
}

type OutputListPeriodDto struct {
	Id        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Year      int    `json:"year"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
