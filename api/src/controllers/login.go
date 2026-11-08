package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		slog.Error("Erro ao enviar os dados do usuário", "error", err)
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		slog.Error("Erro ao converter o json para a struct user", "error", err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		slog.Error("Erro ao iniciar a conexão com o banco de dados no login", "error", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	existentUser, err := repository.FindUserByEmail(user.Email)
	if err != nil {
		slog.Error("Erro ao chamar o repositório no login", "error", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	err = security.VerifyPassword(existentUser.Password, user.Password)
	if err != nil {
		slog.Warn("Usuário não autorizado.", "warn", err)
		response.ResponseError(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.GenerateToken(existentUser.ID)
	if err != nil {
		slog.Error("Erro ao gerar o token do usuário", "error", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	slog.Info("Login realizado com sucesso.")
	fmt.Println(token)
	w.Write([]byte(token))
}
