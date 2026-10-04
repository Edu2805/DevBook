package controllers

import (
	"api/repository"
	"api/src/db"
	"api/src/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"log/slog"
	"net/http"
)

// CreateUser: Cria um usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, requestError := ioutil.ReadAll(r.Body)
	if requestError != nil {
		slog.Error("Erro ao enviar os dados do usuário", "error", requestError)
		log.Fatal(requestError)
	}

	var user models.User
	if requestError = json.Unmarshal(bodyRequest, &user); requestError != nil {
		slog.Error("Erro ao converter o json para a struct user", "error", requestError)
		log.Fatal(requestError)
	}

	db, requestError := db.Connect()
	if requestError != nil {
		slog.Error("Erro ao iniciar a conexão com o banco de dados na inserção do usuário", "error", requestError)
		log.Fatal(requestError)
	}

	repository := repository.NewUsersRepository(db)
	userID, requestError := repository.Create(user)
	if requestError != nil {
		slog.Error("Erro ao chamar o repositório para inserir o usuário", "error", requestError)
		log.Fatal(requestError)
	}

	slog.Info(fmt.Sprintf("Usuário inserido com sucesso, id: %d", userID))

}

// FindUsers: Busca todos os usuários
func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
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
