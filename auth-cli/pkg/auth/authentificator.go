package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"github.com/manifoldco/promptui"
)

type noSignUp struct{}

func (c noSignUp) SignUp(ctx context.Context) (auth.UserInfo, error) {
	return auth.UserInfo{}, errors.New("not implemented")
}

func (c noSignUp) AcceptTermsOfService(ctx context.Context, tos tg.HelpTermsOfService) error {
	return &auth.SignUpRequired{TermsOfService: tos}
}

type Authenticator struct {
	noSignUp
	phone string
}

type promptContent struct {
	label string
}

func promptGetInput(pc promptContent) (string, error) {

	prompt := promptui.Prompt{
		Label: pc.label,
	}
	result, err := prompt.Run()
	return result, err
}

func (a Authenticator) Phone(_ context.Context) (string, error) {
	return a.phone, nil
}

func (a Authenticator) Password(_ context.Context) (string, error) {
	password, err := promptGetInput(promptContent{
		"Пароль",
	})
	return strings.TrimSpace(string(password)), err

}

func (a Authenticator) Code(_ context.Context, _ *tg.AuthSentCode) (string, error) {
	code, err := promptGetInput(promptContent{
		"Код из телеграм",
	})
	return strings.TrimSpace(code), err
}
