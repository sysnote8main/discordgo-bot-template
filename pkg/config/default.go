package config

func GetDefault() Config {
	return Config{
		ConfigVersion: "0.0.1-b1",
		Token:         "Token is here!",
		Prefix:        "!",
	}
}
