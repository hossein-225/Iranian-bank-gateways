package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type BitPayConfig struct {
	RequestURL string `mapstructure:"requestURL"`
	VerifyURL  string `mapstructure:"verifyURL"`
	PayURL     string `mapstructure:"payURL"`
}

type MellatConfig struct {
	URL        string `mapstructure:"url"`
	GatewayURL string `mapstructure:"gatewayURL"`
}

type SamanConfig struct {
	RequestURL string `mapstructure:"requestURL"`
	VerifyURL  string `mapstructure:"verifyURL"`
	PayURL     string `mapstructure:"payURL"`
	ReverseURL string `mapstructure:"reverseURL"`
}

type ZarinpalConfig struct {
	RequestURL    string `mapstructure:"requestURL"`
	VerifyURL     string `mapstructure:"verifyURL"`
	InquiryURL    string `mapstructure:"inquiryURL"`
	UnverifiedURL string `mapstructure:"unverifiedURL"`
	PayURL        string `mapstructure:"payURL"`
}

type SaderatConfig struct {
	RequestURL  string `mapstructure:"requestURL"`
	PayURL      string `mapstructure:"payURL"`
	AdviseURL   string `mapstructure:"adviseURL"`
	RollBackURL string `mapstructure:"rollBackURL"`
}

type Config struct {
	BitPay   BitPayConfig   `mapstructure:"bitpay"`
	Mellat   MellatConfig   `mapstructure:"mellat"`
	Saman    SamanConfig    `mapstructure:"saman"`
	Zarinpal ZarinpalConfig `mapstructure:"zarinpal"`
	Saderat  SaderatConfig  `mapstructure:"saderat"`
}

var AppConfig Config

func LoadConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return fmt.Errorf("failed to read the configuration file: %w", err)
		}
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("failed to Unmarshal the configuration file: %w", err)
	}

	return nil
}
