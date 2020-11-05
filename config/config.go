package config

type Config struct {
	EnvConfig
	YamlConfig
}

var Cfg Config

func Load() (Config, error) {
	env, err := LoadEnv()
	if err != nil {
		return Config{}, err
	}

	yaml, err := loadYaml()
	if err != nil {
		return Config{}, err
	}

	Cfg = Config{env, yaml}
	// fmt.Printf("++++++++++++++++++++++ %#v\n\n", Cfg)
	// fmt.Println("++++", Cfg.SourceHome)

	return Cfg, nil
}
