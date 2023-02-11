package user

import (
	"database/sql"
	"errors"
	"marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
	"time"

	"github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	client *sql.DB
}

type UserModel struct {
	id        string
	tenantId  string
	name      string
	phone     string
	email     string
	createdAt time.Time
	updatedAt time.Time
}

func NewUserRepository(client *sql.DB) IUserRepository {
	return &UserRepository{
		client: client,
	}
}

func (repository *UserRepository) Create(entity entity.IUser) (status.RepositoryInternalStatus, error) {
	var mysqlErr *mysql.MySQLError

	model := UserModel{
		id:        entity.GetId(),
		tenantId:  entity.GetTenantId(),
		name:      entity.GetName(),
		phone:     entity.GetPhone(),
		email:     entity.GetEmail(),
		createdAt: entity.GetCreatedAt(),
	}

	statement, err := repository.client.Prepare("insert into users (id, tenant_id, name, phone, email, created_at) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.id, model.tenantId, model.name, model.phone, model.email, model.createdAt)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		// Unique key violated
		return status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}

func (repository *UserRepository) FindById(id string) (entity.IUser, error) {

	row, err := repository.client.Query("select id, tenant_id, name, phone, email, created_at, updated_at from users where id = ?", id)
	if err != nil {
		return entity.User{}, err
	}
	defer row.Close()

	var userModel UserModel
	if row.Next() {
		if err := row.Scan(&userModel.id, &userModel.tenantId, &userModel.name, &userModel.phone, &userModel.email, &userModel.createdAt, &userModel.updatedAt); err != nil {
			return entity.User{}, err
		}

		user, err := entity.NewUser(userModel.id, userModel.tenantId, userModel.name, userModel.phone, userModel.email, userModel.createdAt, userModel.updatedAt)
		if err != nil {
			return entity.User{}, err
		}
		return user, nil
	}
	return nil, nil
}

func (repository *UserRepository) FindByEmail(email string) (entity.IUser, error) {

	row, err := repository.client.Query("select id, tenant_id, name, phone, email, created_at, updated_at from users where email = ?", email)
	if err != nil {
		return entity.User{}, err
	}
	defer row.Close()

	var userModel UserModel
	if row.Next() {
		if err := row.Scan(&userModel.id, &userModel.tenantId, &userModel.name, &userModel.phone, &userModel.email, &userModel.createdAt, &userModel.updatedAt); err != nil {
			return entity.User{}, err
		}

		user, err := entity.NewUser(userModel.id, userModel.tenantId, userModel.name, userModel.phone, userModel.email, userModel.createdAt, userModel.updatedAt)
		if err != nil {
			return entity.User{}, err
		}
		return user, nil
	}
	return nil, nil
}

func (repository *UserRepository) List(filterParameters []filter.FilterParameter) ([]entity.IUser, error) {

	nameFilter := ""
	emailFilter := ""
	for _, filterParameter := range filterParameters {
		switch filterParameter.Name {
		case "name":
			nameFilter = filterParameter.Value
		case "email":
			emailFilter = filterParameter.Value
		}
	}
	var rows *sql.Rows
	var err error
	if len(filterParameters) == 0 {
		rows, err = repository.client.Query("select id, tenant_id, name, phone, email, created_at, updated_at from users")
	} else {
		if len(nameFilter) > 0 && len(emailFilter) == 0 {
			rows, err = repository.client.Query("select id, tenant_id, name, phone, email, created_at, updated_at from users where name = ?", nameFilter)
		}
		if len(nameFilter) == 0 && len(emailFilter) > 0 {
			rows, err = repository.client.Query("select id, tenant_id, name, phone, email, created_at, updated_at from users where email = ?", emailFilter)
		}
		if len(nameFilter) > 0 && len(emailFilter) > 0 {
			rows, err = repository.client.Query("select id, tenant_id, name, phone, email, created_at, updated_at from users where name = ? and email = ?", nameFilter, emailFilter)
		}
	}

	if err != nil {
		return []entity.IUser{}, err
	}
	defer rows.Close()

	users := []entity.IUser{}
	for rows.Next() {
		var userModel UserModel

		if err := rows.Scan(&userModel.id, &userModel.tenantId, &userModel.name, &userModel.phone, &userModel.email, &userModel.createdAt, &userModel.updatedAt); err != nil {
			return []entity.IUser{}, err
		}

		user, err := entity.NewUser(userModel.id, userModel.tenantId, userModel.name, userModel.phone, userModel.email, userModel.createdAt, userModel.updatedAt)
		if err != nil {
			return []entity.IUser{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository *UserRepository) Update(entity entity.IUser) error {

	model := UserModel{
		id:        entity.GetId(),
		name:      entity.GetName(),
		phone:     entity.GetPhone(),
		email:     entity.GetEmail(),
		updatedAt: entity.GetUpdatedAt(),
	}

	statement, err := repository.client.Prepare("update users set name = ?, phone = ?, email = ?, updated_at = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(model.name, model.phone, model.email, model.updatedAt, model.id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepository) Delete(id string) error {

	statement, err := repository.client.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
