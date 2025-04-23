package hash

import "golang.org/x/crypto/bcrypt"

// Hash возвращает зашифрованый пароль
func Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// CompareHashAndPass сравнивает зашифрованный пароль с другим паролем
// в случае несовпадения возвращает ошибку, иначе - nil
func CompareHashAndPass(hash, realPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(realPassword))
}
