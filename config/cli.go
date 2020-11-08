package config

type CliConfig struct {
	CliName    string
	Version    string
	SourceHome string
	BinHome    string
}

func getCliConfig() CliConfig {
	return CliConfig{
		CliName:    "editorAndChangeDirTest",
		Version:    "v1.0.0",
		SourceHome: "/home/jim/code/privat/editor-and-change-dir/",
		BinHome:    "/home/jim/bin/",
	}
}
