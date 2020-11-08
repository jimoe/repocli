package config

type Config struct {
	CliConfig
	YamlConfig
}

func Load() (Config, error) {
	cli := getCliConfig()

	yaml, err := loadYaml()
	if err != nil {
		return Config{}, err
	}

	cfg := Config{cli, yaml}
	return cfg, nil
}
