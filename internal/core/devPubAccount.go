package core

type DevPubAccount struct {
	Id          int     `json:"id"`
	Login       string  `json:"login"`
	Password    string  `json:"password"`
	Email       string  `json:"email"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
