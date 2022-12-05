package filter

type FilterParameter struct {
	Name     string
	Value    string
	Criteria Criteria
}

type Criteria string

const (
	Exact          Criteria = "exact"
	Like           Criteria = "like"
	List           Criteria = "list"
	Greater        Criteria = "greater"
	GreaterOrEqual Criteria = "greaterOrEqual"
	Lesser         Criteria = "lesser"
	LesserOrEqual  Criteria = "lesserOrEqual"
)
