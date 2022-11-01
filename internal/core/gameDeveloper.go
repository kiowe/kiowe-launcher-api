package core

type GameDeveloper struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
