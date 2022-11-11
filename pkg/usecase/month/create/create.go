package create

import (
	"time"

	"github.com/google/uuid"
)

func Execute(input InputCreateMonthDto) OutputCreateMonthDto {
	return OutputCreateMonthDto{
		Id:        uuid.NewString(),
		Code:      input.Code,
		Name:      input.Name,
		Year:      input.Year,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		CreatedAt: time.Now(),
	}
}
