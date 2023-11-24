# REPOCLI

Command Line Interface (**CLI**) meant to make it easier to work with multiple **repo**sitories

It is written in [go](https://golang.org/) (golang), so it should be able to run on all systems, but it is only tested on 
Ubuntu 20.04 and 22.04

Edit a YAML config file to describe your repositories.

## Commands

* **editor** (alias **e**) - Open repo in the editor defined for that repo.
* **getdir** (alias **d**) - Get the root directory of a repo.
* **tabtitle** (alias **t**) - Get terminal tab title for a repo.
* **config** - Command to initialize or describe a configuration-file. Can also show you the path to the config.

After you build this cli run `repocli help` for all commands or `repocli help <command>` for help on a command

## Build

If you do not have go installed you can follow go's [install instructions](https://golang.org/doc/install) to install it.  It is quite simple.

Download this repo with git: `git clone git@github.com:jimoe/repocli.git`

run `make build` to build it the default directory (~/bin)
or `sudo make build /opt/repocli` if you, for example, want to install it in the */opt* directory

Make sure the 'install' directory is in you PATH, then verify by running `repocli --version`

## Example config

```yaml
editors:
  - name: goland
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
```

It can also handle mono-repoes:
```yaml
  - name:    a-monorepo-name
    path:    /home/username/code/a-monorepo-name
    editor:  code
    aliases:
      - amono
    terminal:
      title: AMONO
    monorepo:
      - subpath: packages/packagename
        terminal:
          title: A packagename
      - subpath: packages/whatever
        terminal:
          title: A whatever
```

After the cli is built you can see a full example by running `repocli config example`

## Practical usage

`repocli editor <repo-name/alias>` works out of the box, but if you want to change directory or change terminal tab title this can only be done in the shell itself.  I use **bash**, and here are som example usage you may put in your `.bashrc`

### Bash aliases

```shell
handleRepocliOutput() {
  local output=$1
  if [[ $output == "/"* ]]; then
    cd $output
  elif [[ ! -z $output ]]; then # not empty output
    echo "$output"
  fi
}
# open repo in configered editor:
e() {
  local alias=$1
  repocli editor $alias
}
# change directory:
c() {
  local alias=$1
  local output=$(repocli getdir $alias)
  handleRepocliOutput "$output"
}
# open editor and change directory:
ec() {
  local alias=$1
  local output=$(repocli editor --returndir $alias)
  handleRepocliOutput "$output"
}
alias ce=ec
```

You may then run for example `e somealias` or `c somealias` given that *somealias* is the name or an alias of a 
repo in your config-file (or even if *some-alias* is the name of your repo)

### Set terminal tab title in bash
```shell
setRepocliTabTitleList() {
  readarray -t tabTitleList < <(/home/jim/bin/repocli tabtitle)
}
setRepocliTabTitleList
getRepocliTabTitleFromList() {
  local pwd=$(pwd -P)
  local IFS=";" # splitChar
  local line
  for line in "${tabTitleList[@]}"; do
    local parts
    read -ra parts <<< "$line" # split string into array using $IFS as split-char
    local path=${parts[0]}
    local title=${parts[1]}
    if [[ "$pwd" == "$path" ]]; then
      echo "$title"
      break
    fi
  done
}
setTabTitle() {
  local title=$(getRepocliTabTitleFromList)
  if [ -z "$title" ]; then
    title=$(pwd|rev|cut -d "/" -f 1-2|rev)
  fi
  echo -en "\e]0;$title\a"
}
PROMPT_COMMAND=setTabTitle
```

This should not affect your PROMPT (PS1) and it will set the tab title every time the prompt is writen
