package create

type InputCreatePeriodDto struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Year      int    `json:"year"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type OutputCreatePeriodDto struct {
	Id        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Year      int    `json:"year"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	CreatedAt string `json:"created_at"`
}
