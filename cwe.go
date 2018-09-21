package cwe

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"os"
)

// CallWithEnvironment contains the additional environment variables to pass to the child process
type CallWithEnvironment struct {
	Environment map[string]string `yaml:"env"`   // Environment contains additional variables to add to the environment
	Quiet       bool              `yaml:"quiet"` // Quiet toggles output of added variables
}

// NewCallWithEnvironment creates a new instance
func NewCallWithEnvironment() (*CallWithEnvironment, error) {
	if _, err := os.Stat(callWithEnvironmentFileName); os.IsNotExist(err) {
		return nil, errors.New(fmt.Sprintf("no %s file in directory", callWithEnvironmentFileName))
	}

	var newCWE CallWithEnvironment

	in, err := ioutil.ReadFile(callWithEnvironmentFileName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(in, &newCWE)
	if err != nil {
		return nil, err
	}

	return &newCWE, nil
}
