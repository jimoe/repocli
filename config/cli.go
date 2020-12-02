package config

type CliConfig struct {
	CliName string
	Version string
}

func getCliConfig() *CliConfig {
	return &CliConfig{
		CliName: "repocli",
		Version: "v1.0.0",
	}
}
