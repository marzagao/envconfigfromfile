# envconfigfromfile

[![Build Status](https://travis-ci.org/marzagao/envconfigfromfile.svg?branch=master)](https://travis-ci.org/marzagao/envconfigfromfile)

My colleague @adammck found that:

> long environment variables can be stored in `/etc/environment`, but `pam_env` (which is responsible for loading those into the shell when a user logs) has a bug which causes it to truncate lines longer than 1024 bytes.

So to simplify the loading of environment variables longer than 1024 bytes using the convenient [envconfig](https://github.com/kelseyhightower/envconfig) tagging, I created this library to set a configuration variable using the contents of a file. The value of the environment variable named in the `envconfig` tag should be the path of the file to be read from.

## Example:

Given that:
* The `TEST_FIELD_FILE_PATH` environment variable has the value "contents.txt"
* The file `contents.txt` has the value `some-content`

The following program:
```go
package main

import (
  "fmt"

  "github.com/kelseyhightower/envconfig"
  "github.com/marzagao/envconfigfromfile"
)

type TestStruct struct {
  TestField *envconfigfromfile.EnvConfigFromFile `envconfig:"TEST_FIELD_FILE_PATH"`
}

func main() {
  testStruct := TestStruct{}
  envconfig.Process("", &testStruct)
  fmt.Println(testStruct.TestField.String())
}
```
Will return `some-content`.
