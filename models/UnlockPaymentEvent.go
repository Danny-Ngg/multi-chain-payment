package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
)

type EventUnlockPayment struct {
	ID                   int64  `json:"id"`
	TxHash               string `json:"tx_hash"`
	BlockNo              uint64 `json:"block_no"`
	Network              string `json:"network"`
	UnlockFromAddress    string `json:"unlock_from_address"`
	UnlockToUserAddress  string `json:"unlock_to_user_address"`
	UnlockToUserAmount   string `json:"unlock_to_user_amount"`
	UnlockToAdminAddress string `json:"unlock_to_admin_address"`
	UnlockToAdminAmount  string `json:"unlock_to_admin_amount"`
	UnlockTime           string `json:"unlock_time"`
	CreateAt             string `json:"create_at"`
}

// FindEventUnlockPayments (&UnlockPaymentEvent{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindEventUnlockPayments(whereCondition interface{}, orderCondition, limit, offset string) ([]*EventUnlockPayment, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*EventUnlockPayment
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
