package list

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IListUseCase interface {
	Execute(InputListUserDto, []filter.FilterParameter) (OutputListUserDto, status.InternalStatus, error)
}

type ListUseCase struct {
	repository user.IUserRepository
}

func NewListUseCase(repository user.IUserRepository) IListUseCase {
	return &ListUseCase{
		repository: repository,
	}
}

func (listUseCase *ListUseCase) Execute(input InputListUserDto, filterParameters []filter.FilterParameter) (OutputListUserDto, status.InternalStatus, error) {

	periods, err := listUseCase.repository.List(filterParameters)
	if err != nil {
		return OutputListUserDto{}, status.InternalServerError, err
	}

	outputListUserDto := OutputListUserDto{}

	if len(periods) == 0 {
		return outputListUserDto, status.NoRecordsFound, nil
	}

	for _, item := range periods {
		user := User{
			Id:    item.GetId(),
			Name:  item.GetName(),
			Phone: item.GetPhone(),
			Email: item.GetEmail(),
		}
		outputListUserDto.Users = append(outputListUserDto.Users, user)
	}
	return outputListUserDto, status.Success, nil
}
