package cfg

import (
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	configFile      = "telegram-auth-cli.toml"
	defaultEndpoint = "http://localhost:8888"
)

var HomeDirectory string

func Init() error {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	HomeDirectory = home
	if !isConfigFileExist() {
		err := writeDefaultConfig()
		if err != nil {
			return err
		}
	}
	return nil
}

type Config struct {
	Endpoint string `json:"endpoint"`
}

//Check existence of the configuration file
func isConfigFileExist() bool {
	if fi, err := os.Stat(path.Join(HomeDirectory, configFile)); err != nil || fi.IsDir() {
		return false
	}
	return true
}

//Read the configuration
func readConfigFile() (*Config, error) {
	viper.SetConfigFile(path.Join(HomeDirectory, configFile))
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Endpoint: viper.GetString("config.endpoint"),
	}, nil
}

//Write data of config to the configuration file
func writeConfigFile(cfg *Config) error {
	_, err := os.Create(path.Join(HomeDirectory, configFile))
	if err != nil {
		return err
	}
	viper.SetConfigName("telegram-auth-cli")
	viper.AddConfigPath(HomeDirectory)
	viper.Set("config.endpoint", cfg.Endpoint)
	return viper.WriteConfig()
}

//Write default config into the configuration file
func writeDefaultConfig() error {
	return writeConfigFile(&Config{
		Endpoint: defaultEndpoint,
	})
}

func GetEndpoint() (string, error) {
	cfg, err := readConfigFile()
	if err != nil {
		return "", err
	}
	return cfg.Endpoint, nil
}

func SaveEndpoint(endpoint string) error {
	err := writeConfigFile(&Config{endpoint})
	if err != nil {
		return err
	}
	return nil
}
