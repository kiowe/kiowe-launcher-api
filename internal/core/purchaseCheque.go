package core

import "github.com/jackc/pgtype"

type PurchaseCheque struct {
	Id             int         `json:"id"`
	DateOfPurchase pgtype.Date `json:"date_of_purchase"`
}
