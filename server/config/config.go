package config

import "encoding/json"

type RedisConfig struct {
	ServerPort string `toml:"server_port"`
	Password   string `toml:"password"`
	Database   int    `toml:"database"`
	KeyPrefix  string `toml:"key_prefix"`
	KeyTTL     int    `toml:"key_ttl"`
}

func (conf *RedisConfig) String() string {
	b, _ := json.Marshal(conf)
	return string(b)
}

type API struct {
	HTTPPrefix string `toml:"http_prefix"`
}
type PortConfig struct {
	Port int `toml:"port"`
}
type Config struct {
	RedisConfig RedisConfig           `toml:"redis"`
	API         API                   `toml:"api"`
	APPList     map[string]PortConfig `toml:"app_list"`
}

var Conf Config
