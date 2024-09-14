package config

func GetDefault() Config {
	return Config{
		ConfigVersion: "0.0.1-b2",
		Token:         "Token is here!",
		Prefix:        "!",
		GuildId:       "",
	}
}
