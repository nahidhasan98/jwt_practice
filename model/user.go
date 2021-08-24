package model

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"Password,omitempty"`
}
