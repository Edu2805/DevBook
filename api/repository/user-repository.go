package repository

import (
	"api/src/models"
	"database/sql"
	"log/slog"
)

// Users representa um repositório de usuários
type Users struct {
	db *sql.DB
}

// NewUsersRepository cria um repositório de usuários
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create cria o usuário e retorna o seu ID
func (u Users) Create(usersModel models.User) (uint64, error) {
	stm, stmError := u.db.Prepare(
		`INSERT INTO users (user_name, nick, email, user_password)
			VALUES($1, $2, $3, $4) RETURNING id`,
	)
	if stmError != nil {
		slog.Error("Erro no prepare statement para inserir o usuário", "error", stmError)
		return 0, stmError
	}
	defer stm.Close()

	var id uint64
	stmError = stm.QueryRow(usersModel.Name, usersModel.Nick, usersModel.Email, usersModel.Password).Scan(&id)
	if stmError != nil {
		return 0, stmError
	}

	return id, nil
}
