package credentials

import (
	"database/sql"
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/credentials/entity"
)

type UserCredentialsRepository struct {
	client *sql.DB
}

type UserCredentialsModel struct {
	id        string
	userId    string
	tenantId  string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func NewUserCredentialsRepository(client *sql.DB) IUserCredentialsRepository {
	return &UserCredentialsRepository{
		client: client,
	}
}

func (repository *UserCredentialsRepository) Create(entity entity.IUserCredentials) error {
	model := UserCredentialsModel{
		id:        entity.GetId(),
		userId:    entity.GetUserId(),
		password:  entity.GetPassword(),
		createdAt: entity.GetCreatedAt(),
	}

	statement, err := repository.client.Prepare("insert into user_credentials (id, user_id, password, created_at) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(model.id, model.userId, model.password, model.createdAt)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserCredentialsRepository) Update(entity entity.IUserCredentials) error {
	model := UserCredentialsModel{
		id:        entity.GetId(),
		userId:    entity.GetUserId(),
		password:  entity.GetPassword(),
		updatedAt: entity.GetUpdatedAt(),
	}

	statement, err := repository.client.Prepare("update user_credentials set password = ?, updated_at = ? where user_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(model.password, model.updatedAt, model.userId)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserCredentialsRepository) FindByUserId(userId string) (entity.IUserCredentials, error) {

	row, err := repository.client.Query("select user_credentials.id, user_credentials.user_id, users.tenant_id, user_credentials.password, user_credentials.created_at, user_credentials.updated_at from user_credentials inner join users on user_credentials.user_id = users.id where user_credentials.user_id = ?", userId)
	if err != nil {
		return entity.UserCredentials{}, err
	}
	defer row.Close()

	var userCredentialsModel UserCredentialsModel
	if row.Next() {
		if err := row.Scan(&userCredentialsModel.id, &userCredentialsModel.userId, &userCredentialsModel.tenantId, &userCredentialsModel.password, &userCredentialsModel.createdAt, &userCredentialsModel.updatedAt); err != nil {
			return entity.UserCredentials{}, err
		}

		user, err := entity.NewUserCredentials(userCredentialsModel.id, userCredentialsModel.userId, userCredentialsModel.tenantId, userCredentialsModel.password, userCredentialsModel.createdAt, userCredentialsModel.updatedAt)
		if err != nil {
			return entity.UserCredentials{}, err
		}
		return user, nil
	}
	return nil, nil
}

func (repository *UserCredentialsRepository) FindByUserEmail(userEmail string) (entity.IUserCredentials, error) {
	row, err := repository.client.Query("select user_credentials.id, user_credentials.user_id, users.tenant_id, user_credentials.password, user_credentials.created_at, user_credentials.updated_at from user_credentials inner join users on user_credentials.user_id = users.id where users.email = ?", userEmail)
	if err != nil {
		return entity.UserCredentials{}, err
	}
	defer row.Close()

	var userCredentialsModel UserCredentialsModel
	if row.Next() {
		if err := row.Scan(&userCredentialsModel.id, &userCredentialsModel.userId, &userCredentialsModel.tenantId, &userCredentialsModel.password, &userCredentialsModel.createdAt, &userCredentialsModel.updatedAt); err != nil {
			return entity.UserCredentials{}, err
		}

		userCredentials, err := entity.NewUserCredentials(userCredentialsModel.id, userCredentialsModel.tenantId, userCredentialsModel.userId, userCredentialsModel.password, userCredentialsModel.createdAt, userCredentialsModel.updatedAt)
		if err != nil {
			return entity.UserCredentials{}, err
		}
		return userCredentials, nil
	}
	return nil, nil
}
