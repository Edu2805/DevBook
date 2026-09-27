package db

import (
	"api/src/config"
	"database/sql"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func connect() (*sql.DB, error) {
	db, dbError := sql.Open("pgx", config.StringDatabaseConnection)
	if dbError != nil {
		slog.Error("Erro ao abrir conexão com o banco de dados", "error", dbError)
		return nil, dbError
	}

	if dbError = db.Ping(); dbError != nil {
		db.Close()
		slog.Error("Erro de conexão com o banco de dados", "error", dbError)
		return nil, dbError
	}

	slog.Info("Banco de dados conectado")
	return db, nil
}
