package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"

	"github.com/sysnote8main/discordgo-bot-template/pkg/util"
)

var (
	configPath = "./config.toml"
	config     Config
)

type Config struct {
	ConfigVersion string `toml:"configVersion"`
	Token         string `toml:"token"`
	Prefix        string `toml:"prefix"`
}

func configFileExists() bool {
	return util.FileExists(configPath)
}

func GetConfig() Config {
	return config
}

func ReadConfig() {
	if configFileExists() {
		slog.Info("Config file is found! Loading...")
		err := read()
		if err == nil {
			slog.Info("Successfully to load config")
			if config.ConfigVersion == GetDefault().ConfigVersion {
				return
			} else {
				slog.Warn("Config version is miss matched. Replacing new one.")
			}
		}
		slog.Warn("Failed to read config", slog.Any("error", err))
		oldConfigPath := configPath + ".old"
		slog.Warn(fmt.Sprintf("Moving old config... (path:%s)", oldConfigPath))
		err = os.Rename(configPath, oldConfigPath)
		config = GetDefault()
		if err != nil {
			slog.Error("Failed to move old config!", slog.Any("error", err))
			slog.Warn("Continue with default config without saving to file")
		} else {
			slog.Info("Successfully to move old config")
			slog.Info("Writing new config file...")
			err = WriteConfig()
			if err != nil {
				slog.Error("Failed to write config", slog.Any("error", err))
			}
		}
	} else {
		slog.Warn("Config file is not found.")
		config = GetDefault()
		err := WriteConfig()
		if err != nil {
			slog.Error("Failed to write config", slog.Any("error", err))
		}
	}
}

func read() error {
	d, err := util.ReadFile(configPath)
	if err != nil {
		return err
	}
	_, err = toml.Decode(*d, &config)
	if err != nil {
		return err
	}
	return nil
}

func WriteConfig() error {
	slog.Info("Trying to write config...")
	b, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	err = util.WriteFile(configPath, string(b))
	return err
}
