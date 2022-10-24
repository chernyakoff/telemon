package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitnub.com/chernyakoff/telegram-auth-cli/pkg/api"
	"gitnub.com/chernyakoff/telegram-auth-cli/pkg/auth"
	"gitnub.com/chernyakoff/telegram-auth-cli/pkg/cfg"
	"gitnub.com/chernyakoff/telegram-auth-cli/pkg/prompt"
)

func init() {
	cobra.OnInitialize()
}

var rootCmd = &cobra.Command{
	Use:   "telegram-auth-cli",
	Short: "A simple CLI for create auth sessions in db",
	Long:  "A simple CLI for create auth sessions in db",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := cfg.Init()
		if err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		SetEndpoint()
		account := prompt.Wizard()

		result, err := auth.Do(account)
		if err != nil {
			fmt.Printf("Error: %#v", err)
		} else {
			response, _ := api.Send(result)
			fmt.Println(response)
		}
		return nil
	},
}

func SetEndpoint() {
	endpoint, err := cfg.GetEndpoint()
	if err != nil {
		panic(err)
	}
	online := api.Ping(endpoint)
	if !online {
		invalidEndpoint(endpoint)
		for {
			endpoint, _ = prompt.Endpoint("Укажите правильный URL сервиса")
			online = api.Ping(endpoint)
			if online {
				fmt.Println("Новый URL сервиса сохранён")
				cfg.SaveEndpoint(endpoint)
				break
			} else {
				invalidEndpoint(endpoint)
			}
		}
	} else {
		fmt.Printf("Сервер работает по адресу %s\n", endpoint)
	}

}

func invalidEndpoint(endpoint string) {
	fmt.Printf("Cервер по адресу %s находится не в доступе\n", endpoint)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
