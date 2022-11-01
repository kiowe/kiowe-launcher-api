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

type CreateGameDTO struct {
	Name         string  `form:"name"`
	Price        float64 `form:"price"`
	IdDevelopers int     `form:"id_developers"`
	IdPublishers int     `form:"id_publishers"`
	IdCategories *int    `form:"id_categories,omitempty"`
	SystemReq    string  `form:"system_req"`
	AgeLimit     string  `form:"age_limit"`
	Description  *string `form:"description,omitempty"`
	ReleaseDate  string  `form:"release_date"`
	Version      string  `form:"version"`
	Rating       float64 `form:"rating"`
}

type UpdateGameDTO struct {
	Name         string  `form:"name,omitempty"`
	Price        float64 `form:"price,omitempty"`
	IdDevelopers int     `form:"id_developers,omitempty"`
	IdPublishers int     `form:"id_publishers,omitempty"`
	IdCategories int     `form:"id_categories,omitempty"`
	SystemReq    string  `form:"system_req,omitempty"`
	AgeLimit     string  `form:"age_limit,omitempty"`
	Description  string  `form:"description,omitempty"`
	ReleaseDate  string  `form:"release_date,omitempty"`
	Version      string  `form:"version,omitempty"`
	Rating       float64 `form:"rating,omitempty"`
}
