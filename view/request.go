package view

import "transfer-service/model"

type CreateAccountRequest struct {
	AccountID      int64          `json:"account_id"`
	InitialBalance model.Decimal5 `json:"initial_balance"`
}

type QueryAccountRequest struct {
	AccountID int64 `form:"account_id"`
}

type TransactionRequest struct {
	SourceAccountID      int64          `json:"source_account_id"`
	DestinationAccountID int64          `json:"destination_account_id"`
	Amount               model.Decimal5 `json:"amount"`
}
