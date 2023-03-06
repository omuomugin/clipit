# clipit: CLI Simple Snippet Manager
It's heavily inspired by [knqyf263/pet](https://github.com/knqyf263/pet) but have some features for my own usage.

## TL;DR
```shell
clipit clip "git push <remote=github|gitlab> <branch=master>"

clipit list

clipit exec "git push <remote=github|gitlab> <branch=master>"
# choose `github` or `gitlab` for `remote` in selectable prompt.
# Variable for <remote=github|gitlab>: 
#    github
#  ▸ gitlab  
#
# override or use `master` for `branch` in prompt.
# Variable for <branch=master>: master|
```

## Installation

## Usage
```shell
Usage:
  clipit [command]

Available Commands:
  clip        Save snippet.
  completion  Generate the autocompletion script for the specified shell
  config      Commands for setting configuration.
  edit        Edit commands with selected editor.
  exec        Execute passed command.
  help        Help about any command
  list        List all saved commands.

Flags:
  -h, --help   help for clipit
```

### exec
```shell
clipit exec "echo something"
```

#### Variables
you can use variables too.

```shell
clipit clip "gh pr view <pr-num>" 
```
This will launch prompt for setting `pr-num` when executing.

```shell
clipit clip "git push <remote=github|gitlab> <branch=master>"
```
You can set default values and choices too.  
`<remote=github|gitlab>` will launch selectable ui for choosing `github` or `gitlab`.  
`<branch=master>` will set default value but you can override it in prompt.

### clip
saves snippets in local file.
```shell
clipit clip "echo something"
```

### list
```shell
clipit list
```

List all the commands.  
Its main usage it to pass it to [peco/peco](https://github.com/peco/peco) or some kind of different tools.