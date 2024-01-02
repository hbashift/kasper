package mapping

import "github.com/google/uuid"

type Authorization struct {
	ClientType string    `json:"client_type,omitempty"`
	Token      uuid.UUID `json:"token,omitempty"`
}

type AuthorizeInfo struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type ChangePassword struct {
	Email       string `json:"email"`
	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}
