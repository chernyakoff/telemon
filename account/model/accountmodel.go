package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AccountModel = (*customAccountModel)(nil)

type (
	// AccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountModel.
	AccountModel interface {
		accountModel
		Upsert(ctx context.Context, account *Account) (string, error)
		FindAny(ctx context.Context) (*Account, error)
	}

	customAccountModel struct {
		*defaultAccountModel
	}
)

func (m *customAccountModel) Upsert(ctx context.Context, account *Account) (string, error) {
	message := ""
	var err error
	exists, _ := m.FindOne(ctx, account.Phone)
	if exists != nil {
		err := m.Update(ctx, account)
		if err != nil {
			message = "Экаунт упешно обновлён"
		}
	} else {
		_, err := m.Insert(ctx, account)
		if err != nil {
			message = "Экаунт упешно добавлен"
		}
	}
	return message, err

}
func (m *customAccountModel) FindAny(ctx context.Context) (*Account, error) {

	query := fmt.Sprintf("select %s from %s order by rand limit 0,1", accountRowsWithPlaceHolder, m.table)
	var resp Account
	err := m.conn.QueryRowCtx(ctx, &resp, query)
	return &resp, err

}

// NewAccountModel returns a model for the database table.
func NewAccountModel(conn sqlx.SqlConn) AccountModel {
	return &customAccountModel{
		defaultAccountModel: newAccountModel(conn),
	}
}
