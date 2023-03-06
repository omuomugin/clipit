# clipit: CLI Simple Snippet Manager

## Usage

```shell
help
```

### clip
```shell
clipit clip YOUR-COMMAND
```

#### variables
```shell
clipit clip gh pr view <pr> 
```

with choices and default values
```shell
clipit clip git push <remote=github|gitlab> <branch=master>
```

### exec
```shell
clipit exec git push <remote=github|gitlab> <branch=master>
```

This will first set up prompt to ask your choice for remote, then set up another prompt for inputting branch name.


### list
```shell
clipit list
```

list all the command 