package repository

import (
	"api/src/models"
	"database/sql"
	"log/slog"
	"strings"
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

// FindByNameOrNick monta a query para busca de usuários pelo nome ou NickName
func (userRepository Users) FindByNameOrNick(nameOrNick string) ([]models.User, error) {
	nameOrNick = strings.TrimSpace(nameOrNick)
	if nameOrNick == "" {
		return []models.User{}, nil
	}

	pattern := "%" + nameOrNick + "%"

	rows, err := userRepository.db.Query(
		`SELECT id, user_name, nick, email, created_at
		 FROM users
		 WHERE user_name ILIKE $1 OR nick LIKE $2`,
		pattern,
		pattern,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// FindUserById monta a query para busca de usuários pelo pelo seu ID no banco
func (userRepository Users) FindUserById(userID uint64) (models.User, error) {
	rows, err := userRepository.db.Query(
		`SELECT id, user_name, nick, email, created_at
		 FROM users
		 WHERE id = $1`,
		userID,
	)

	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil

}

// UpdateUserById altera as informações de um usuário no banco de dados
func (userRepository Users) UpdateUserById(userID uint64, user models.User) error {
	result, err := userRepository.db.Exec(
		`UPDATE users SET user_name = $1, nick = $2, email = $3 WHERE id = $4`,
		user.Name,
		user.Nick,
		user.Email,
		userID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
