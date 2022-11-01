package core

type GamePublisher struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
