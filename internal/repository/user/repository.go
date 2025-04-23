package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/merynayr/jwtauth/internal/client/db"
	"github.com/merynayr/jwtauth/internal/model"
	"github.com/merynayr/jwtauth/internal/repository"
)

// Константы названий столбцов БД
const (
	usersTable = "users"

	UserIDColumn = "user_id"
)

// Структура репо с клиентом базы данных (интерфейсом)
type repo struct {
	db db.Client
}

// NewRepository возвращает новый объект репо слоя
func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

// CreateUser создаёт нового пользователя
func (r *repo) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	query, args, err := sq.Insert(usersTable).
		PlaceholderFormat(sq.Dollar).
		Columns(
			UserIDColumn,
		).
		Values(
			user.ID,
		).
		Suffix("RETURNING *").
		ToSql()

	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.CreateUser",
		QueryRaw: query,
	}
	var userObject model.User
	err = r.db.DB().ScanOneContext(ctx, &userObject, q, args...)
	if err != nil {
		return nil, err
	}

	return &userObject, nil
}
