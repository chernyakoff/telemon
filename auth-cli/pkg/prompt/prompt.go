package prompt

import (
	"fmt"

	"github.com/gernest/ita"
	"github.com/manifoldco/promptui"
	"github.com/mitchellh/mapstructure"
)

var v Validate

var t = ita.New(&Transform{})

var questions = []Question{
	{
		Label:    "Телефон",
		Name:     "Phone",
		Validate: v.PhoneNumber,
	},
	{
		Name:     "AppId",
		Validate: v.AppId,
	},
	{
		Name:     "AppHash",
		Validate: v.AppHash,
	},
}

func answer(dict map[string]interface{}, q Question) {
	for {
		prompt := promptui.Prompt{
			Label:    q.GetLabel(),
			Validate: q.Validate,
		}
		input, err := prompt.Run()
		if err == nil {
			dict[q.Name] = transform(q, input)
			break
		} else {
			fmt.Println(err)
		}
	}
}

func transform(q Question, input string) interface{} {
	var result interface{}
	called := t.Call(q.Name, input)
	result, exists := called.GetResults().First()
	if !exists {
		return input
	}
	return result
}

func convert(dict map[string]interface{}) Account {
	var account Account

	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &account,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(dict)
	if err != nil {
		panic(err)
	}
	return account
}

func Wizard() Account {
	dict := make(map[string]interface{})
	for _, q := range questions {
		answer(dict, q)
	}
	account := convert(dict)
	return account

}

func Endpoint(name string) (string, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: v.EmptyInput,
	}
	return prompt.Run()
}
