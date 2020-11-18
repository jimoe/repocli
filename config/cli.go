package config

type CliConfig struct {
	CliName string
	Version string
}

func getCliConfig() *CliConfig {
	return &CliConfig{
		CliName: "editorAndChangeDir",
		Version: "v1.0.0",
	}
}
