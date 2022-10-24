package prompt

import (
	"github.com/manifoldco/promptui"
)

type Account struct {
	Phone   int64  `mapstructure:"Phone"`
	AppHash string `mapstructure:"AppHash"`
	Session string `mapstructure:"Session"`
	AppId   int64  `mapstructure:"AppId"`
}

type Question struct {
	Name     string
	Label    string
	Type     string
	Validate promptui.ValidateFunc
}

func (q Question) GetLabel() string {
	if q.Label != "" {
		return q.Label
	}
	return q.Name
}
