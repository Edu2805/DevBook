package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
	"strings"
)

// CreateUser: Cria um usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, requestError := ioutil.ReadAll(r.Body)
	if requestError != nil {
		slog.Error("Erro ao enviar os dados do usuário", "error", requestError)
		response.ResponseError(w, http.StatusUnprocessableEntity, requestError)
		return
	}

	var user models.User
	if requestError = json.Unmarshal(bodyRequest, &user); requestError != nil {
		slog.Error("Erro ao converter o json para a struct user", "error", requestError)
		response.ResponseError(w, http.StatusBadRequest, requestError)
		return
	}

	if requestError = user.Prepare(); requestError != nil {
		response.ResponseError(w, http.StatusBadRequest, requestError)
		return
	}

	db, requestError := db.Connect()
	if requestError != nil {
		slog.Error("Erro ao iniciar a conexão com o banco de dados na inserção do usuário", "error", requestError)
		response.ResponseError(w, http.StatusInternalServerError, requestError)
		return
	}

	repository := repository.NewUsersRepository(db)
	user.ID, requestError = repository.Create(user)
	if requestError != nil {
		slog.Error("Erro ao chamar o repositório para inserir o usuário", "error", requestError)
		response.ResponseError(w, http.StatusInternalServerError, requestError)
		return
	}

	slog.Info(fmt.Sprintf("Usuário inserido com sucesso, id: %d", user.ID))
	response.JSON(w, http.StatusCreated, user)

}

// FindUsers: Busca todos os usuários
func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.TrimSpace(r.URL.Query().Get("user_name"))

	db, dbError := db.Connect()
	if dbError != nil {
		slog.Error("Erro ao iniciar a conexão com o banco de dados na busca de usuários", "error", dbError)
		response.ResponseError(w, http.StatusInternalServerError, dbError)
		return
	}
	defer db.Close()

	userRepository := repository.NewUsersRepository(db)
	users, dbError := userRepository.FindByNameOrNick(nameOrNick)

	if dbError != nil {
		slog.Error("Erro ao iniciar a conexão com o banco de dados na busca de usuários", "error", dbError)
		response.ResponseError(w, http.StatusInternalServerError, dbError)
		return
	}

	slog.Info("Lista de usuários retornada com sucesso")
	response.JSON(w, http.StatusOK, users)
}

// FindUserById: Busca um usuário pelo seu id
func FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bucando usuário por id"))
}

// UpdateUser: Atualiza um usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

// DeleteUser: Exclui um usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
