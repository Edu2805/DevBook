package routes

import (
	"api/src/controllers"
	"net/http"
)

const (
	UsersURI   = "/users"
	UsersIDURI = UsersURI + "/{userId}"
)

var routersUsers = []Route{
	{
		URI:          UsersURI,
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		RequiredAuth: false,
	},
	{
		URI:          UsersURI,
		Method:       http.MethodGet,
		Function:     controllers.FindUsers,
		RequiredAuth: true,
	},
	{
		URI:          UsersIDURI,
		Method:       http.MethodGet,
		Function:     controllers.FindUserById,
		RequiredAuth: true,
	},
	{
		URI:          UsersIDURI,
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		RequiredAuth: true,
	},
	{
		URI:          UsersIDURI,
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		RequiredAuth: true,
	},
}
