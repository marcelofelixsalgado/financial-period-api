package delete

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IDeleteUseCase interface {
	Execute(InputDeleteCategoryDto) (OutputDeleteCategoryDto, status.InternalStatus, error)
}

type DeleteUseCase struct {
	repository category.ICategoryRepository
}

func NewDeleteUseCase(repository category.ICategoryRepository) IDeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (deleteUseCase *DeleteUseCase) Execute(input InputDeleteCategoryDto) (OutputDeleteCategoryDto, status.InternalStatus, error) {

	// Find the entity before update
	category, err := deleteUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputDeleteCategoryDto{}, status.InternalServerError, err
	}
	if category == nil {
		return OutputDeleteCategoryDto{}, status.InvalidResourceId, err
	}

	// Apply in dabatase
	err = deleteUseCase.repository.Delete(input.Id)
	if err != nil {
		return OutputDeleteCategoryDto{}, status.InternalServerError, err
	}

	outputDeleteCategoryDto := OutputDeleteCategoryDto{}

	return outputDeleteCategoryDto, status.Success, nil
}
