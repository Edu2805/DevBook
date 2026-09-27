package models

import "time"

// Usuário representa um usuário utilizando a rede social
type User struct {
	ID        uint      `json:"id,omitempty"` // omitempty omite o id quando o mesmo nao possuir valores
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt"`
}
