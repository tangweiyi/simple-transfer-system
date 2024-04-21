package repo

import (
	"context"
	"errors"
	"transfer-service/model"

	"github.com/joomcode/errorx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewGormRepo(db *gorm.DB) AccountRepo {
	return AccountRepo{
		db: db,
	}
}

func (r AccountRepo) CreateAccount(c context.Context, accountID int64, balance model.Decimal5) *errorx.Error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&model.Account{
			Base: model.Base{
				ID: accountID,
			},
			Balance: balance,
		}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&model.Transaction{
			Type:                 model.TransactionTypeCreation,
			DestinationAccountID: accountID,
			Amount:               balance,
		}).Error
		return err
	})
	if err != nil {
		return errorx.EnsureStackTrace(err)
	}
	return nil
}

func (r AccountRepo) GetAccount(c context.Context, accountID int64) (model.Account, *errorx.Error) {
	ret := model.Account{}
	err := r.db.First(&ret, accountID).Error
	if err != nil {
		return ret, errorx.EnsureStackTrace(err)
	}
	return ret, nil
}

func (r AccountRepo) MakeTransfer(c context.Context, sourceID, destID int64, amount model.Decimal5) *errorx.Error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// only source is locked, since destination always get incremented
		var source model.Account
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&source, sourceID).Error
		if err != nil {
			return err
		}
		if source.Balance.LessThan(amount.Decimal) {
			return errors.New("insufficient balance")
		}
		if res := tx.Exec("UPDATE account SET balance = balance - ? WHERE id = ?", amount, sourceID); res.RowsAffected != 1 {
			return errors.New("update did not happen")
		} else {
			if res.Error != nil {
				return err
			}
		}
		if res := tx.Exec("UPDATE account SET balance = balance + ? WHERE id = ?", amount, destID); res.RowsAffected != 1 {
			return errors.New("update did not happen")
		} else {
			if res.Error != nil {
				return err
			}
		}

		err = tx.Create(&model.Transaction{
			Type:                 model.TransactionTypeTransfer,
			Amount:               amount,
			SourceAccountID:      sourceID,
			DestinationAccountID: destID,
		}).Error
		return err
	})
	if err != nil {
		return errorx.EnsureStackTrace(err)
	}
	return nil
}
