package auth

import (
	"context"
	"fmt"

	_ "github.com/gotd/td"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"gitnub.com/chernyakoff/telegram-auth-cli/pkg/prompt"
	"go.uber.org/zap"
)

func getAuthFlow(account prompt.Account) auth.Flow {

	return auth.NewFlow(
		Authenticator{phone: fmt.Sprint(account.Phone)},
		auth.SendCodeOptions{},
	)
}

func Do(account prompt.Account) (prompt.Account, error) {

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func() { _ = log.Sync() }()

	sessionStorage := &memorySession{}
	ctx := context.Background()
	flow := getAuthFlow(account)

	client := telegram.NewClient(int(account.AppId), account.AppHash, telegram.Options{
		SessionStorage: sessionStorage,
		//	Logger:         log,
	})

	err = client.Run(ctx, func(ctx context.Context) error {

		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return account, err
	}
	session, err := sessionStorage.LoadSession(ctx)
	if err != nil {
		return account, err
	}
	account.Session = string(session)
	return account, err
}
