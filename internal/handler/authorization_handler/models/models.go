package models

type Authorization struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
