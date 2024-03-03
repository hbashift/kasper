package models

type AuthorizeResponse struct {
	UserType   string `json:"user_type,omitempty"`
	Token      string `json:"token,omitempty"`
	Registered bool   `json:"registered,omitempty"`
}

type AuthorizeRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
