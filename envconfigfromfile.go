package envconfigfromfile

import (
	"fmt"
	"io/ioutil"
)

// EnvConfigFromFile can be used as a custom type for kelseyhightower/envconfig variables
// It assumes that the value of the tagged environment variable is a file path, then when
// envconfig processes this variable, it reads the file found in that file path and sets
// the value of the struct to be the contents of that file.
type EnvConfigFromFile struct {
	FilePath string
	Value    string
}

// Set satisfies kelseyhightower/envconfig Setter interface
func (c *EnvConfigFromFile) Set(value string) error {
	c.FilePath = value
	var fileContents []byte
	var err error
	if c.FilePath != "" {
		fileContents, err = ioutil.ReadFile(c.FilePath)
		if err != nil {
			err = fmt.Errorf("envconfigfromfile: %s", err.Error())
		} else {
			c.Value = string(fileContents)
		}
	}
	return err
}

func (c *EnvConfigFromFile) String() string {
	return c.Value
}
