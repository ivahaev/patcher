
# patcher
This is util to patch version in go programs. By default it will increment third segment of string `version` variable in `main.go` file.
For example:
**main.go**
```go
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	version  = "0.0.7"
	filename = "main.go"
	varName  = "version"
)
```

After running util, variable `version` will look exact this:
```
	version  = "0.0.8"
```
After, it will commit file and tag this commit with related version. Also it can push changes immediatly to repo.

## Install
```
    go get github.com/ivahaev/patcher
```

## Usage
```
    patcher [filename [varName]] [-p]
```

Where optional arguments: **filename** (**main.go** by default) is a file where variable **varName** (**version** by default) located.
When **-p** flag provided,  **patcher** will push changes immediatly to the repository.