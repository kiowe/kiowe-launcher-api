package core

import (
	"github.com/jackc/pgtype"
)

type Game struct {
	Id           int         `json:"id"`
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	IdDevelopers int         `json:"id_developers"`
	IdPublishers int         `json:"id_publishers"`
	IdCategories *int        `json:"id_categories,omitempty"`
	SystemReq    string      `json:"system_req"`
	AgeLimit     string      `json:"age_limit"`
	Description  *string     `json:"description,omitempty"`
	ReleaseDate  pgtype.Date `json:"release_date"`
	Version      string      `json:"version"`
	Rating       float64     `json:"rating"`
}

type GamePage struct {
	Games    []*Game `json:"games"`
	Page     int     `json:"page"`
	Total    int     `json:"total"`
	LastPage int     `json:"last_page"`
}

type CreateGame struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	IdDevelopers int     `json:"id_developers"`
	IdPublishers int     `json:"id_publishers"`
	IdCategories *int    `json:"id_categories,omitempty"`
	SystemReq    string  `json:"system_req"`
	AgeLimit     string  `json:"age_limit"`
	Description  *string `json:"description,omitempty"`
	ReleaseDate  string  `json:"release_date"`
	Version      string  `json:"version"`
	Rating       float64 `json:"rating"`
}

type CreateGameDTO struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	IdCategories *int    `json:"id_categories,omitempty"`
	SystemReq    string  `json:"system_req"`
	AgeLimit     string  `json:"age_limit"`
	Description  *string `json:"description,omitempty"`
	ReleaseDate  string  `json:"release_date"`
	Version      string  `json:"version"`
	Rating       float64 `json:"rating"`
}

type UpdateGameDTO struct {
	Name         *string  `json:"name,omitempty"`
	Price        *float64 `json:"price,omitempty"`
	IdDevelopers *int     `json:"id_developers,omitempty"`
	IdPublishers *int     `json:"id_publishers,omitempty"`
	IdCategories *int     `json:"id_categories,omitempty"`
	SystemReq    *string  `json:"system_req,omitempty"`
	AgeLimit     *string  `json:"age_limit,omitempty"`
	Description  *string  `json:"description,omitempty"`
	ReleaseDate  *string  `json:"release_date,omitempty"`
	Version      *string  `json:"version,omitempty"`
	Rating       *float64 `json:"rating,omitempty"`
}
