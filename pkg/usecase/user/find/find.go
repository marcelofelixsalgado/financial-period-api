package find

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"
)

type IFindUseCase interface {
	Execute(InputFindUserDto) (OutputFindUserDto, status.InternalStatus, error)
}

type FindUseCase struct {
	repository user.IRepository
}

func NewFindUseCase(repository user.IRepository) IFindUseCase {
	return &FindUseCase{
		repository: repository,
	}
}

func (findUseCase *FindUseCase) Execute(input InputFindUserDto) (OutputFindUserDto, status.InternalStatus, error) {

	user, err := findUseCase.repository.Find(input.Id)
	if err != nil {
		return OutputFindUserDto{}, status.InternalServerError, err
	}
	if user == nil {
		return OutputFindUserDto{}, status.InvalidResourceId, err
	}

	if user.GetId() == "" {
		return OutputFindUserDto{}, status.NoRecordsFound, err
	}

	outputFindUserDto := OutputFindUserDto{
		Id:        user.GetId(),
		Name:      user.GetName(),
		Password:  user.GetPassword(),
		Phone:     user.GetPhone(),
		Email:     user.GetEmail(),
		CreatedAt: user.GetCreatedAt().Format(time.RFC3339),
	}

	if !user.GetUpdatedAt().IsZero() {
		outputFindUserDto.UpdatedAt = user.GetUpdatedAt().Format(time.RFC3339)
	}

	return outputFindUserDto, status.Success, nil
}
