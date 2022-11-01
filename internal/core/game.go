package core

import "github.com/jackc/pgtype"

type Game struct {
	Id           int         `json:"id"`
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	IdDevelopers int         `json:"id_developers"`
	IdPublishers int         `json:"id_publishers"`
	IdCategories int         `json:"id_categories"`
	SystemReq    string      `json:"system_req"`
	AgeLimit     string      `json:"age_limit"`
	Description  *string     `json:"description,omitempty"`
	ReleaseDate  pgtype.Date `json:"release_date"`
	Version      string      `json:"version"`
	Rating       float64     `json:"rating"`
}
