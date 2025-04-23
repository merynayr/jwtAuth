package sys

import "github.com/merynayr/jwtauth/internal/sys/codes"

// Константы с текстами ошибок
const (
	ErrNotFound                = "not found"
	ErrInvalidRequest          = "invalid request"
	ErrAccessDenied            = "access denied"
	ErrInvalidAuthHeaderFormat = "invalid authorization header format"
	ErrAuthHeaderNotProvided   = "authorization header is not provided"
	ErrInvalidAccessToken      = "invalid access token"
	ErrInvalidRefreshToken     = "invalid refresh token"
	ErrInvalidCredentials      = "invalid login or password"
	ErrUserAlreadyExists       = "user already exists"
	ErrUserNotFound            = "user not found"
	ErrRefreshTokenExpired     = "refresh token expired"
	ErrAccessRefreshMismatch   = "access and refresh tokens are mismatch"
	ErrAccessTokenExpired      = "access token expired"
)

// Готовые объекты ошибок с поясняющими комментариями
var (
	// NotFoundError — ошибка 404: ресурс не найден (например, пользователь или склад)
	NotFoundError = NewCommonError(ErrNotFound, codes.NotFound)

	// InvalidRequestError — ошибка 400: некорректный формат запроса или тело запроса
	InvalidRequestError = NewCommonError(ErrInvalidRequest, codes.BadRequest)

	// AccessDeniedError — ошибка 403: у пользователя нет доступа к ресурсу или действию
	AccessDeniedError = NewCommonError(ErrAccessDenied, codes.Forbidden)

	// InvalidAuthHeaderFormatError — ошибка 401: заголовок авторизации передан в неверном формате
	InvalidAuthHeaderFormatError = NewCommonError(ErrInvalidAuthHeaderFormat, codes.Unauthorized)

	// AuthHeaderNotProvidedError — ошибка 401: заголовок авторизации отсутствует
	AuthHeaderNotProvidedError = NewCommonError(ErrAuthHeaderNotProvided, codes.Unauthorized)

	// InvalidAccessTokenError — ошибка 401: access-токен недействителен или просрочен
	InvalidAccessTokenError = NewCommonError(ErrInvalidAccessToken, codes.Unauthorized)

	// InvalidRefreshTokenError — ошибка 401: refresh-токен недействителен или не найден
	InvalidRefreshTokenError = NewCommonError(ErrInvalidRefreshToken, codes.Unauthorized)

	// InvalidCredentialsError — ошибка 401: неверный логин или пароль
	InvalidCredentialsError = NewCommonError(ErrInvalidCredentials, codes.Unauthorized)

	// UserAlreadyExistsError — ошибка 409: пользователь с таким логином уже существует
	UserAlreadyExistsError = NewCommonError(ErrUserAlreadyExists, codes.Conflict)

	// UserNotFoundError — ошибка 404: пользователь не найден
	UserNotFoundError = NewCommonError(ErrUserNotFound, codes.NotFound)

	// RefreshTokenExpiredError — ошибка 403: refresh-токен истёк
	RefreshTokenExpiredError = NewCommonError(ErrRefreshTokenExpired, codes.Forbidden)

	// AccessRefreshMismatchError — ошибка 403: mismatch между access и refresh токенами
	AccessRefreshMismatchError = NewCommonError(ErrAccessRefreshMismatch, codes.Forbidden)

	// TokenExpiredError — ошибка 401: access токен истёк
	AccessTokenExpiredError = NewCommonError(ErrAccessTokenExpired, codes.Unauthorized)
)
