package config

import "encoding/json"

type RedisConfig struct {
	ServerPort string `toml:"server_port"`
	Password   string `toml:"password"`
	Database   int    `toml:"database"`
	KeyPrefix  string `toml:"key_prefix"`
	KeyTTL     int    `toml:"key_ttl"`
}

type API struct {
	HTTPPrefix string `toml:"http_prefix"`
}
type APP struct {
	PortBegin int `toml:"port_begin"`
}
type Config struct {
	RedisConfig RedisConfig `toml:"redis"`
	API         API         `toml:"api"`
	APP         APP         `toml:"app"`
}

func (conf *Config) String() string {
	b, _ := json.Marshal(conf)
	return string(b)
}

var Conf Config
