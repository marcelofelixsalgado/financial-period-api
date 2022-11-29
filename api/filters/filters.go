package requests

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"
	"net/http"
)

func SetupFilters(r *http.Request) ([]repository.FilterParameter, error) {

	filterParameters := []repository.FilterParameter{}

	queryParams := r.URL.Query()
	for name, value := range queryParams {
		filterParameter := repository.FilterParameter{
			Name:  name,
			Value: value[0],
		}
		filterParameters = append(filterParameters, filterParameter)
	}

	return filterParameters, nil
}

// func getSuffix(fieldName string) string {
// 	suffixes := map[repository.Criteria]string {
// 		"_like", "_in", "_gt", "_gte", "_lt", "_lte"
// 	}
// }
