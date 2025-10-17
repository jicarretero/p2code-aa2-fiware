package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	MQTT struct {
		MqttUrl      string `toml:"mqtt_url"`
		MqttTopic    string `toml:"mqtt_topic"`
		MqttUser     string `toml:"mqtt_user"`
		MqttPassword string `toml:"mqtt_password"`
	} `toml:"mqtt"`

	Brokerld struct {
		URL                string `toml:"url"`
		Context            string `toml:"context"`
		WalletType         string `toml:"wallet_type"`
		WalletAddress      string `toml:"wallet_address"`
		WalletToken        string `toml:"wallet_token"`
		Tenant             string `toml:"tenant"`
		UseIDM             bool   `toml:"use_idm"`
		IdmWalletDirectory string `toml:"idm_wallet_directory"`
		IdmWalletHelper    string `toml:"idm_wallet_helper"`
		IdmServiceURL      string `toml:"idm_service_url"`
		IdmScope           string `toml:"idm_scope"`
		IdmClientId        string `toml:"idm_client_id"`
	} `toml:"brokerld"`
}

func ReadConfig(filePath string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(filePath, &config)
	if err != nil {
		return nil, fmt.Errorf("error decoding TOML file: %w", err)
	}
	return &config, nil
}
