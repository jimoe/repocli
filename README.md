# REPOCLI

Command Line Interface (**CLI**) meant to make it easier to work with multiple **repo**sitories

It is written in [go](https://golang.org/) (golang)

Edit a YAML config file to describe your repositories.

## Commands

* **editor** - Open repo in the editor defined for that repo
* **getdir** - Get the root directory of a repo
* **tabtitle** - Get terminal tab title for repo

After you build this cli run *repocli help* for all details

## Build

As this cli is written in [go](https://golang.org/) it should be able to run on all systems, but it is only tested on Ubuntu 20.04.

If you do not have go installed you can follow go's [install instructions](https://golang.org/doc/install) to install it.  It is quite simple.

Download this repo with git: `git clone git@github.com:jimoe/repocli.git`

run `make build` to build it the default directory (~/bin)
or `make build /some/direcotory` if you want to install it somewhere else

Make sure the install directory is in you PATH, then verify by running `repocli --version`

## Practical usage

`repocli edtitor <repo-name/alias>` works out of the box, but if you want to change directory or change terminal tab title this can only be done in the shell itself.  I use **bash**, and here are som example usage in bash,

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
```

### Set terminal tab title in bash
```shell
setTabTitleList() {
  readarray -t tabTitleList < <(repocli tabtitle)
}
setTabTitleList
getTabTitleFromList() {
  local pwd=$(pwd -P)
  local IFS=";" # splitChar
  local str
  local parts
  for str in "${tabTitleList[@]}"; do
    read -ra parts <<< "$str" # split string into array using $IFS as split-char
    if [[ "$pwd" == "${parts[0]}" ]]; then
      echo "${parts[1]}"
      break
    fi
  done
}
setTabTitle() {
  local title=$(getTabTitleFromList)
  echo -en "\e]0;$title\a"
}
PROMPT_COMMAND=setTabTitle
```

# Example config
