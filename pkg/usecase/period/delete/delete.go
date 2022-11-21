package delete

// func Execute(input InputDeletePeriodDto, repository repository.IRepository) (OutputDeletePeriodDto, error) {

// 	var outputDeletePeriodDto OutputDeletePeriodDto

// 	entity, err := entity.NewPeriod(input.Code, input.Name, input.Year, startDate, endDate)
// 	if err != nil {
// 		return outputCreatePeriodDto, err
// 	}

// 	// Persists in dabatase
// 	err = repository.Create(entity)
// 	if err != nil {
// 		return outputCreatePeriodDto, err
// 	}

// 	outputCreatePeriodDto = OutputCreatePeriodDto{
// 		Id:        entity.GetId(),
// 		Code:      entity.GetCode(),
// 		Name:      entity.GetName(),
// 		Year:      entity.GetYear(),
// 		StartDate: entity.GetStartDate().String(),
// 		EndDate:   entity.GetEndDate().String(),
// 		CreatedAt: entity.GetCreatedAt(),
// 	}

// 	return outputCreatePeriodDto, nil
// }
