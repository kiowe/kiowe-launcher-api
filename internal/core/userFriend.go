package core

type UserFriend struct {
	Id       int  `json:"id"`
	IdUser   *int `json:"id_user,omitempty"`
	IdFriend *int `json:"id_friend,omitempty"`
}
