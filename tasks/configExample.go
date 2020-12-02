package tasks

import (
	"fmt"

	"github.com/jimoe/repocli/config"
)

var example = `Yaml config
Place config in same dir and with the same name as the executable. Just add '.yml' to the filename.

For the 'editors' object:
	The editor will be executed from the repo path (set in the 'repoes' object).
	If 'params' includes the string '<path>' then it will be replaced with the repo path.

EXAMPLE:

editors:
	-	name: goland
		params: nosplash <path>
	- name: code
		params: .
repoes:
  - name:    some-repo-name
    path:    /home/username/code/some-repo-name
    editor:  goland
    aliases:
      - some
			- some-repo
    terminal:
      title: SOME
  - name:    another-repo-name
    path:    /home/username/code/another-repo-name
    editor:  code
    monorepo:
			- subpath: packages/name
				terminal:
          title: A name
			- subpath: packages/whatever
				terminal:
          title: A whatever

`

func ConfigExample(cfg *config.Config) {
	fmt.Print(example)
}
