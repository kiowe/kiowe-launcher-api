package core

type GamePurchase struct {
	Id       int `json:"id"`
	IdUser   int `json:"id_user"`
	IdGame   int `json:"id_game"`
	IdCheque int `json:"id_cheque"`
}
