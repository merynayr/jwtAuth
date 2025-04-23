package auth

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/merynayr/jwtauth/internal/client/db"
	"github.com/merynayr/jwtauth/internal/model"
	"github.com/merynayr/jwtauth/internal/repository"
)

// Название таблицы и столбцов
const (
	refreshTokensTable = "jwt"

	RefreshIDColumn = "id"
	UserIDColumn    = "user_id"
	TokenHashColumn = "token_hash"
	AccessJTIColumn = "access_jti"
	IPColumn        = "ip"
	UserAgentColumn = "user_agent"
	CreatedAtColumn = "created_at"
	ExpiresAtColumn = "expires_at"
)

// Структура репо с клиентом базы данных (интерфейсом)
type repo struct {
	db db.Client
}

// NewRepository возвращает новый объект репо слоя
func NewRepository(db db.Client) repository.AuthRepository {
	return &repo{db: db}
}

// CreateRefreshToken сохраняет refresh токен в базу
func (r *repo) CreateRefreshToken(ctx context.Context, token *model.RefreshToken) (*model.RefreshToken, error) {
	query, args, err := sq.Insert(refreshTokensTable).
		PlaceholderFormat(sq.Dollar).
		Columns(
			RefreshIDColumn,
			UserIDColumn,
			TokenHashColumn,
			AccessJTIColumn,
			IPColumn,
			UserAgentColumn,
			CreatedAtColumn,
			ExpiresAtColumn,
		).
		Values(
			token.ID,
			token.UserID,
			token.TokenHash,
			token.AccessJTI,
			token.IP,
			token.UserAgent,
			token.CreatedAt,
			token.ExpiresAt,
		).
		Suffix("RETURNING *").
		ToSql()

	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "refresh_token_repository.CreateRefreshToken",
		QueryRaw: query,
	}
	var tokenObject model.RefreshToken
	err = r.db.DB().ScanOneContext(ctx, &tokenObject, q, args...)
	if err != nil {
		return nil, err
	}

	return &tokenObject, nil
}

func (r *repo) FindRefreshToken(ctx context.Context, jti string) (*model.RefreshToken, error) {
	query, args, err := sq.
		Select("*").
		From(refreshTokensTable).
		Where(sq.Eq{AccessJTIColumn: jti}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "refresh_token_repository.FindByHash",
		QueryRaw: query,
	}

	var token model.RefreshToken
	err = r.db.DB().ScanOneContext(ctx, &token, q, args...)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *repo) DeleteByID(ctx context.Context, jti string) error {
	query, args, err := sq.
		Delete(refreshTokensTable).
		Where(sq.Eq{AccessJTIColumn: jti}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "refresh_token_repository.DeleteByID",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}
