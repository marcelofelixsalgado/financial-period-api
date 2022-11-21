package list

type InputListPeriodDto struct {
}

type Period struct {
	Id        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Year      int    `json:"year"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type OutputListPeriodDto struct {
	Periods []Period `json:"-"`
}
