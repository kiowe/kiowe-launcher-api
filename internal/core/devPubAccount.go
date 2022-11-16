package core

type DevPubAccount struct {
	Id          int     `json:"id"`
	Login       string  `json:"login"`
	Password    string  `json:"password"`
	Email       string  `json:"email"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type DevPubAccPw struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type DevPubAccountDTO struct {
	Login       string  `json:"login"`
	Password    string  `json:"password"`
	Email       string  `json:"email"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type LoginDevPubAccountDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
