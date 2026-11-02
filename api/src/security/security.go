package security

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca um hash nela
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(passWithHash, passString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passWithHash), []byte(passString))
}
