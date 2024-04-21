package service

import (
	"context"
	"errors"
	"transfer-service/model"
	"transfer-service/repo"

	"github.com/joomcode/errorx"
	"github.com/shopspring/decimal"
)

type AccountService struct {
	accountRepo repo.AccountRepo
}

func NewAccountService(repo repo.AccountRepo) AccountService {
	return AccountService{
		accountRepo: repo,
	}
}

func (s AccountService) CreateAccount(c context.Context, id int64, balance model.Decimal5) *errorx.Error {
	return s.accountRepo.CreateAccount(c, id, balance)
}

func (s AccountService) GetAccount(c context.Context, id int64) (model.Account, *errorx.Error) {
	return s.accountRepo.GetAccount(c, id)
}

func (s AccountService) MakeTransfer(c context.Context, sourceID int64, destID int64, amount model.Decimal5) *errorx.Error {
	if amount.LessThan(decimal.Zero) {
		return errorx.EnsureStackTrace(errors.New("transaction amount cannot be less than 0"))
	}

	return s.accountRepo.MakeTransfer(c, sourceID, destID, amount)
}
