package controllers

import "net/http"

// CreateUser: Cria um usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
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
