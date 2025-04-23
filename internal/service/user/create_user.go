package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/merynayr/jwtauth/internal/model"
)

// Register создаёт пользователя и возвращает его ID
func (s *srv) CreateUser(ctx context.Context) (*model.User, error) {
	user := &model.User{
		ID: uuid.NewString(),
	}

	GUID, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return GUID, nil
}
