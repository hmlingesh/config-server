package config

import (
	"io/ioutil"
	"strings"
	"encoding/json"
	"errors"
)

type ServerConfig struct {
	Port                int
	CertificateFilePath string `json:"certificate_file_path"`
	PrivateKeyFilePath  string `json:"private_key_file_path"`
	Store               string
	Database            DBConfig
}

type DBConnectionConfig struct {
	MaxOpenConnections int `json:"max_open_connections"`
	MaxIdleConnections int `json:"max_idle_connections"`
}

type DBConfig struct {
	Adapter           string
	User              string
	Password          string
	Host              string
	Port              int
	Name              string `json:"db_name"`
	ConnectionOptions DBConnectionConfig `json:"connection_options"`
}

func ParseConfig(filename string) (ServerConfig, error) {

	config := ServerConfig{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		return config, err
	}

	if (config.CertificateFilePath == "" || config.PrivateKeyFilePath == "") {
		return config, errors.New("Certificate file path and key file path should be defined")
	}

	if (&config.Database != nil) && (&config.Database.Adapter != nil) {
		config.Database.Adapter = strings.ToLower(config.Database.Adapter)
	}

	return config, nil;
}
