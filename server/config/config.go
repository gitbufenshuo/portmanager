package config

// [redis]
// server_port = "127.0.0.1:6379"
// user = "user"
// password = "password"
// [api]
// http_prefix = "/portmanager"
// key_ttl = 20
// [app_list]
//     [my_web_server]
//     port = 30000
//     [my_job_server]
//     port = 30100
//     # [...]
type RedisConfig struct {
	ServerPort string `toml:"server_port"`
	Password   string `toml:"password"`
}
type API struct {
	HTTPPrefix string `toml:"http_prefix"`
	KeyTTL     string `toml:"key_ttl"`
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
