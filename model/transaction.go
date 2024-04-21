package model

const (
	TransactionTypeInvalid int32 = iota
	TransactionTypeCreation
	TransactionTypeTransfer
)

type Transaction struct {
	Base
	Type                 int32
	Amount               Decimal5
	SourceAccountID      int64
	DestinationAccountID int64
}
