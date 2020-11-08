package config

type Config struct {
	CliConfig
	YamlConfig
}

var Cfg Config

func Load() (Config, error) {
	cli := getCliConfig()

	yaml, err := loadYaml()
	if err != nil {
		return Config{}, err
	}

	Cfg = Config{cli, yaml}
	// fmt.Printf("++++++++++++++++++++++ %#v\n\n", Cfg)
	// fmt.Println("++++", Cfg.SourceHome)

	return Cfg, nil
}
