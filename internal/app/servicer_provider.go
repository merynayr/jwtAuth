package app

import (
	"context"
	"log"

	"github.com/merynayr/jwtauth/internal/client/db"
	"github.com/merynayr/jwtauth/internal/client/db/pg"
	"github.com/merynayr/jwtauth/internal/client/db/transaction"
	"github.com/merynayr/jwtauth/internal/closer"
	"github.com/merynayr/jwtauth/internal/config"
	"github.com/merynayr/jwtauth/internal/config/env"
	"github.com/merynayr/jwtauth/internal/repository"
	"github.com/merynayr/jwtauth/internal/service"

	"github.com/merynayr/jwtauth/internal/api/auth"
	"github.com/merynayr/jwtauth/internal/api/user"

	authRepository "github.com/merynayr/jwtauth/internal/repository/auth"
	authService "github.com/merynayr/jwtauth/internal/service/auth"

	userRepository "github.com/merynayr/jwtauth/internal/repository/user"
	userService "github.com/merynayr/jwtauth/internal/service/user"

	"github.com/merynayr/jwtauth/internal/middleware"
)

// Структура приложения со всеми зависимости
type serviceProvider struct {
	pgConfig      config.PGConfig
	httpConfig    config.HTTPConfig
	loggerConfig  config.LoggerConfig
	swaggerConfig config.SwaggerConfig
	authConfig    config.AuthConfig

	dbClient  db.Client
	txManager db.TxManager

	userAPI        *user.API
	userService    service.UserService
	userRepository repository.UserRepository

	authAPI        *auth.API
	authService    service.AuthService
	authRepository repository.AuthRepository

	middleware middleware.Middleware
}

// NewServiceProvider возвращает новый объект API слоя
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig инициализирует конфиг PostgreSQL
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

// HTTPConfig инициализирует конфиг http сервера
func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

// LoggerConfig инициализирует конфиг логгера
func (s *serviceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := env.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config:%v", err)
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

// AuthSwaggerConfig инициализирует конфиг swagger
func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

// AuthConfig инициализирует конфиг auth сервиса
func (s *serviceProvider) AuthConfig() config.AuthConfig {
	if s.authConfig == nil {
		cfg, err := env.NewAuthConfig()
		if err != nil {
			log.Fatalf("failed to get auth config")
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

// DBClient инициализирует подключение к БД
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("db ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// TxManager инициализирует менеджер транзакций
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// UserAPI инициализирует api слой user
func (s *serviceProvider) UserAPI(ctx context.Context) *user.API {
	if s.userAPI == nil {
		s.userAPI = user.NewAPI(s.UserService(ctx))
	}

	return s.userAPI
}

// UserService иницилизирует сервисный слой auth
func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
		)
	}

	return s.userService
}

// UserRepository иницилизирует репо слой user
func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

// AuthAPI инициализирует api слой auth
func (s *serviceProvider) AuthAPI(ctx context.Context) *auth.API {
	if s.authAPI == nil {
		s.authAPI = auth.NewAPI(s.AuthService(ctx), s.AuthConfig())
	}

	return s.authAPI
}

// AuthService иницилизирует сервисный слой auth
func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.AuthRepository(ctx),
			s.AuthConfig(),
		)
	}

	return s.authService
}

// UserRepository иницилизирует репо слой user
func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authRepository.NewRepository(s.DBClient(ctx))
	}

	return s.authRepository
}

// Middleware инициализирует middleware доступа
func (s *serviceProvider) Middleware(_ context.Context) middleware.Middleware {
	if s.middleware == nil {
		s.middleware = middleware.NewMiddlewareProvider()
	}
	return s.middleware
}
