package controllers

import (
	"api/constants"
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if requestError = user.Prepare(constants.USER_REGISTER); requestError != nil {
		response.ResponseError(w, http.StatusBadRequest, requestError)
		return
	}

	db, requestError := db.Connect()
	if requestError != nil {
		slog.Error("Erro ao iniciar a conexão com o banco de dados na inserção do usuário", "error", requestError)
		response.ResponseError(w, http.StatusInternalServerError, requestError)
		return
	}
	defer db.Close()

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
	parameter := mux.Vars(r)

	userID, parameterError := strconv.ParseUint(parameter["userId"], 10, 64)
	if parameterError != nil {
		slog.Error("Erro ao converter o parâmetro userId", "error", parameterError)
		response.ResponseError(w, http.StatusBadRequest, parameterError)
		return
	}

	db, dbError := db.Connect()
	if dbError != nil {
		slog.Error("Erro ao iniciar a conexão com o banco de dados na busca de usuários", "error", dbError)
		response.ResponseError(w, http.StatusInternalServerError, dbError)
		return
	}
	defer db.Close()

	userRepository := repository.NewUsersRepository(db)
	user, parameterError := userRepository.FindUserById(userID)

	if parameterError != nil {
		slog.Error("Erro ao buscar o usuário no banco de dados", "error", parameterError)
		response.ResponseError(w, http.StatusInternalServerError, parameterError)
		return
	}

	if user.ID == 0 {
		slog.Warn("Usuário não encontrado.")
		response.JSON(w, http.StatusNotFound, map[string]string{
			"warn": "Usuário não encontrado",
		})
		return
	}

	slog.Info("Usuário retornado com sucesso")
	response.JSON(w, http.StatusOK, user)
}

// UpdateUser: Atualiza um usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)

	userID, err := strconv.ParseUint(parameter["userId"], 10, 64)
	if err != nil {
		slog.Error("Erro ao converter o parâmetro userId", "error", err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := auth.ExtractUserId(r)
	if err != nil {
		slog.Warn("Não autorizado.", "warn", err)
		response.ResponseError(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		err = errors.New("Operação não permitida")
		slog.Warn("Operação não permitida.", "warn", err)
		response.ResponseError(w, http.StatusForbidden, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		slog.Error("Erro ao enviar os dados do usuário", "error", err)
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		slog.Error("Erro ao converter o json para a struct user", "error", err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(constants.NOT_USER_REGISTER); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		slog.Error("Erro ao iniciar a conexão com o banco de dados na busca de usuários", "error", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewUsersRepository(db)
	if err = userRepository.UpdateUserById(userID, user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("Usuário não encontrado", "warn", err)
			response.JSON(w, http.StatusNotFound, map[string]string{
				"erro": "Usuário não encontrado",
			})
			return
		}
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Usuário atualizado com sucesso")
	response.JSON(w, http.StatusNoContent, nil)

}

// DeleteUser: Exclui um usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)

	userID, err := strconv.ParseUint(parameter["userId"], 10, 64)
	if err != nil {
		slog.Error("Erro ao converter o parâmetro userId", "error", err)
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := auth.ExtractUserId(r)
	if err != nil {
		slog.Warn("Não autorizado.", "warn", err)
		response.ResponseError(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		err = errors.New("Operação não permitida")
		slog.Warn("Operação não permitida.", "warn", err)
		response.ResponseError(w, http.StatusForbidden, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		slog.Error("Erro ao iniciar a conexão com o banco de dados na busca de usuários", "error", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewUsersRepository(db)
	if err = userRepository.DeleteUserById(userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("Usuário não encontrado", "warn", err)
			response.JSON(w, http.StatusNotFound, map[string]string{
				"erro": "Usuário não encontrado",
			})
			return
		}
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Usuário excluído com sucesso")
	response.JSON(w, http.StatusNoContent, nil)
}
