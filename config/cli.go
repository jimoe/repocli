package config

type CliConfig struct {
	CliName    string `envconfig:"CLI_NAME" required:"true"`
	SourceHome string `envconfig:"SOURCE_HOME" required:"true"`
	BinHome    string `envconfig:"BIN_HOME" required:"true"`
}

func getCliConfig() CliConfig {
	return CliConfig{
		CliName:    "editorAndChangeDirTest",
		SourceHome: "/home/jim/code/privat/editor-and-change-dir/",
		BinHome:    "/home/jim/bin/",
	}
}
