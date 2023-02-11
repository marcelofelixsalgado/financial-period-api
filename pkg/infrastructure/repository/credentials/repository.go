package credentials

import "marcelofelixsalgado/financial-period-api/pkg/domain/credentials/entity"

type IUserCredentialsRepository interface {
	Create(entity.IUserCredentials) error
	Update(entity.IUserCredentials) error
	// FindById(id string) (entity.IUserCredentials, error)
	FindByUserId(userId string) (entity.IUserCredentials, error)
	FindByUserEmail(email string) (entity.IUserCredentials, error)
}
