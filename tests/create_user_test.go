package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/merynayr/jwtauth/internal/model"
	"github.com/merynayr/jwtauth/internal/repository"
	repositoryMocks "github.com/merynayr/jwtauth/internal/repository/mocks"
	"github.com/merynayr/jwtauth/internal/service/user"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
	}

	var (
		ctx = context.Background()
	)

	tests := []struct {
		name               string
		args               args
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
			},
			err: nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)

				mock.CreateUserMock.Set(func(_ context.Context, u *model.User) (*model.User, error) {
					require.NotEmpty(t, u.ID, "expected non-empty UUID")
					return u, nil
				})

				return mock
			},
		},
		{
			name: "repository returns error",
			args: args{
				ctx: ctx,
			},
			err: errors.New("failed to create user"),
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)

				mock.CreateUserMock.Set(func(_ context.Context, _ *model.User) (*model.User, error) {
					return nil, errors.New("failed to create user")
				})

				return mock
			},
		},
		{
			name: "context canceled",
			args: func() args {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return args{ctx: ctx}
			}(),
			err: context.Canceled,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)

				mock.CreateUserMock.Set(func(ctx context.Context, u *model.User) (*model.User, error) {
					if ctx.Err() != nil {
						return nil, ctx.Err()
					}
					return u, nil
				})

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mc := minimock.NewController(t)
			defer mc.Finish()

			userRepoMock := tt.userRepositoryMock(mc)

			service := user.NewService(userRepoMock)
			createdUser, err := service.CreateUser(tt.args.ctx)

			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
				require.NotNil(t, createdUser)
				require.NotEmpty(t, createdUser.ID)
			}
		})
	}
}
