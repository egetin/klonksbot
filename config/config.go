package config

import (
	"flag"
	"os"

	"github.com/egetin/klonksbot/database"
	"github.com/naoina/toml"
)

type Config struct {
	Basic    BotConfig
	Database database.DbConfig
}

type BotConfig struct {
	Nick     string
	Ident    string
	RealName string
	Server   string
	Port     string
	SSL      bool
	QuitMsg  string
	Channels []string
}

var (
	configPath = flag.String("config-file", "config.toml", "Path to the configuration file")
)

func init() {
	flag.Parse()
}

func GetBotConfig() (config *Config) {
	parseConfigFile(config)

	return
}

func parseConfigFile(config *Config) {
	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := toml.NewDecoder(f).Decode(&config); err != nil {
		panic(err)
	}
}
