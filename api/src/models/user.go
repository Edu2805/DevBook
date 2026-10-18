package models

import (
	"errors"
	"log/slog"
	"strings"
	"time"
)

const INVALID_EMPTY_NAME_MESSAGE = "O nome é obrigatório e não pode estar em branco."
const INVALID_EMPTY_NICK_MESSAGE = "O nickname é obrigatório e não pode estar em branco."
const INVALID_EMPTY_EMAIL_MESSAGE = "O email é obrigatório e não pode estar em branco."
const INVALID_EMPTY_PASSWORD_MESSAGE = "A senha é obrigatória e não pode estar em branco."

// Usuário representa um usuário utilizando a rede social
type User struct {
	ID        uint64    `json:"id,omitempty"` // omitempty omite o id quando o mesmo nao possuir valores
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt"`
}

// Prepare vai chamar os métodos para validar e formatar o usuário recebido. isRegister informa se o Prepare é usado no contexto de cadastro de usuário ou não
func (user *User) Prepare(isRegister bool) error {
	if prepareError := user.validate(isRegister); prepareError != nil {
		return prepareError
	}

	user.formatStrings()
	return nil
}

func (user *User) validate(isRegister bool) error {
	if user.Name == "" {
		slog.Warn(INVALID_EMPTY_NAME_MESSAGE)
		return errors.New(INVALID_EMPTY_NAME_MESSAGE)
	}
	if user.Nick == "" {
		slog.Warn(INVALID_EMPTY_NICK_MESSAGE)
		return errors.New(INVALID_EMPTY_NICK_MESSAGE)
	}
	if user.Email == "" {
		slog.Warn(INVALID_EMPTY_EMAIL_MESSAGE)
		return errors.New(INVALID_EMPTY_EMAIL_MESSAGE)
	}
	if isRegister == true && user.Password == "" {
		slog.Warn(INVALID_EMPTY_PASSWORD_MESSAGE)
		return errors.New(INVALID_EMPTY_PASSWORD_MESSAGE)
	}

	return nil
}

func (user *User) formatStrings() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
