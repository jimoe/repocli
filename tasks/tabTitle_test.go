package tasks

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/jimoe/repocli/arguments"
	"github.com/jimoe/repocli/config"
)

var mockCfg = &config.Config{
	CliName: "testcli",
	Version: "v1.1.1",
	Yaml: &config.Yaml{
		Filename:        "someconf.yml",
		Path:            "/path/to/config/file",
		PathAndFilename: "/path/to/config/file/someconf.yml",
	},
	YamlConfig: &config.YamlConfig{
		Editors: []*config.Editor{
			{
				Name:   "code",
				Params: ".",
			},
			{
				Name: "golang",
			},
		},
		Repoes: []*config.Repo{
			{
				Name:    "some-repo",
				Path:    "/path/to/a/stairway",
				Editor:  "golang",
				Aliases: []string{"some", "repo"},
				Terminal: &config.Terminal{
					Title: "Tsome",
				},
			},
			{
				Name:    "some-repo-without-alias",
				Path:    "/path/to/another/stairway",
				Editor:  "golang",
				Aliases: []string{},
			},
			{
				Name:    "mono-repo",
				Path:    "/jupp",
				Editor:  "code",
				Aliases: []string{"mono"},
				Terminal: &config.Terminal{
					Title: "Tmono",
				},
				MonoRepo: []*config.MonoRepo{
					{
						SubPath: "packages/fun",
						Terminal: &config.Terminal{
							Title: "Tfun",
						},
					},
					{
						SubPath: "packages/droll",
						Terminal: &config.Terminal{
							Title: "Tdroll",
						},
					},
				},
			},
		},
	},
}

func TestGetTabTitle(t *testing.T) {
	soh := stdOutHandler{}

	soh.capture()
	GetTabTitle(mockCfg, arguments.NewPath("/path/to/a/stairway"))
	got := soh.restore()
	exp := "Tsome"
	if got != exp {
		t.Errorf("unexpected response: got '%s', exp '%s'", got, exp)
	}

	soh.capture()
	GetTabTitle(mockCfg, arguments.NewPath("/not/in/config"))
	got = soh.restore()
	exp = ""
	if got != exp {
		t.Errorf("unexpected response: got '%s', exp '%s'", got, exp)
	}
}

type stdOutHandler struct {
	origStdOut *os.File
	read       *os.File
	write      *os.File
}

func (s *stdOutHandler) capture() {
	s.origStdOut = os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		panic("failed to initialize capture stdout")
	}

	os.Stdout = w
	s.read = r
	s.write = w
}

func (s *stdOutHandler) restore() string {
	s.write.Close()

	out, err := ioutil.ReadAll(s.read)
	if err != nil {
		panic("failed to finalize capture stdout")
	}

	os.Stdout = s.origStdOut

	return strings.TrimSpace(string(out))
}
