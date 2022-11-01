package core

import "github.com/jackc/pgtype"

type User struct {
	Id               int         `json:"id"`
	Login            string      `json:"login"`
	Password         string      `json:"password"`
	Email            string      `json:"email"`
	PhoneNumber      *string     `json:"phone_number,omitempty"`
	Nickname         string      `json:"nickname"`
	Description      *string     `json:"description,omitempty"`
	RegistrationDate pgtype.Date `json:"registration_date"`
	IdLibrary        int         `json:"id_library"`
	IdGroups         *int        `json:"id_groups,omitempty"`
	IdStatus         int         `json:"id_status"`
	IdWishlist       *int        `json:"id_wishlist,omitempty"`
}
