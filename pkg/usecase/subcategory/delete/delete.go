package delete

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
)

type IDeleteUseCase interface {
	Execute(InputDeleteUseCaseDto) (OutputDeleteUseCaseDto, status.InternalStatus, error)
}

type DeleteUseCase struct {
	repository subcategory.ISubCategoryRepository
}

func NewDeleteUseCase(repository subcategory.ISubCategoryRepository) IDeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (deleteUseCase *DeleteUseCase) Execute(input InputDeleteUseCaseDto) (OutputDeleteUseCaseDto, status.InternalStatus, error) {

	// Find the entity before update
	subCategory, err := deleteUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputDeleteUseCaseDto{}, status.InternalServerError, err
	}
	if subCategory == nil {
		return OutputDeleteUseCaseDto{}, status.InvalidResourceId, err
	}

	// Apply in dabatase
	err = deleteUseCase.repository.Delete(input.Id)
	if err != nil {
		return OutputDeleteUseCaseDto{}, status.InternalServerError, err
	}

	outputDeleteUseCaseDto := OutputDeleteUseCaseDto{}

	return outputDeleteUseCaseDto, status.Success, nil
}
