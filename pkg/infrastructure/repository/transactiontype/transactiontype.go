package transactiontype

import (
	"database/sql"
	"marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
)

type TransactionTypeRepository struct {
	client *sql.DB
}

type TransactionTypeModel struct {
	Code string
	Name string
}

func NewTransactionTypeRepository(client *sql.DB) ITransactionTypeRepository {
	return &TransactionTypeRepository{
		client: client,
	}
}

func (repository TransactionTypeRepository) FindByCode(code string) (entity.ITransactionType, error) {

	row, err := repository.client.Query("select code, name from transaction_types where code = ?", code)
	if err != nil {
		return entity.TransactionType{}, err
	}
	defer row.Close()

	var transactionTypeModel TransactionTypeModel
	if row.Next() {
		if err := row.Scan(&transactionTypeModel.Code, &transactionTypeModel.Name); err != nil {
			return entity.TransactionType{}, err
		}

		transactionType, err := entity.NewTransactionType(transactionTypeModel.Code, transactionTypeModel.Name)
		if err != nil {
			return entity.TransactionType{}, err
		}
		return transactionType, nil
	}
	return nil, nil
}

func (repository TransactionTypeRepository) List(filterParameters []filter.FilterParameter) ([]entity.ITransactionType, error) {

	codeFilter := ""
	nameFilter := ""
	for _, filterParameter := range filterParameters {
		switch filterParameter.Name {
		case "code":
			codeFilter = filterParameter.Value
		case "name":
			nameFilter = filterParameter.Value
		}
	}

	var rows *sql.Rows
	var err error
	if len(filterParameters) == 0 {
		rows, err = repository.client.Query("select code, name from transaction_types")
	} else {
		if len(codeFilter) > 0 && len(nameFilter) == 0 {
			rows, err = repository.client.Query("select code, name from transaction_types where code = ?", codeFilter)
		}
		if len(codeFilter) == 0 && len(nameFilter) > 0 {
			rows, err = repository.client.Query("select code, name from transaction_types where name = ?", nameFilter)
		}
		if len(codeFilter) > 0 && len(nameFilter) > 0 {
			rows, err = repository.client.Query("select code, name from transaction_types where code = ? and name = ?", codeFilter, nameFilter)
		}
	}
	defer rows.Close()
	if err != nil {
		return []entity.ITransactionType{}, err
	}

	transactionTypes := []entity.ITransactionType{}
	for rows.Next() {
		var transactionTypeModel TransactionTypeModel

		if err := rows.Scan(&transactionTypeModel.Code, &transactionTypeModel.Name); err != nil {
			return []entity.ITransactionType{}, err
		}

		transactionType, err := entity.NewTransactionType(transactionTypeModel.Code, transactionTypeModel.Name)
		if err != nil {
			return []entity.ITransactionType{}, err
		}

		transactionTypes = append(transactionTypes, transactionType)
	}

	return transactionTypes, nil
}
