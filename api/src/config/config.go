package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringDatabaseConnection é a string de conexão com o Postrgres
	StringDatabaseConnection = ""
	// Port onde a API vai estar rodando
	Port = 0
)

// CarrLoadConfigs vai inicializar as variáveis de ambiente
func LoadConfigs() {

	var errorLoad error

	if errorLoad = godotenv.Load(); errorLoad != nil {
		slog.Error("Erro ao carregar as variáveis de ambiente.", "error", errorLoad)
		log.Fatal(errorLoad)
	}

	Port, errorLoad = strconv.Atoi(os.Getenv("API_PORT"))
	if errorLoad != nil {
		slog.Warn("Erro ao converter a string Port para número, usando Port: 9000.", "warn", errorLoad)
		Port = 9000
	}

	StringDatabaseConnection = fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

}
