package cwe

import (
	"fmt"
	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
	"os"
	"strings"
)

// Run executes the program with the environment required
func (cwe *CallWithEnvironment) Run(program string) error {
	variables := cwe.buildEnvironment()
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	p, err := syntax.NewParser().Parse(strings.NewReader(program), "")
	if err != nil {
		return err
	}

	env := variables
	r := interp.Runner{
		Dir: dir,
		Env: env,

		Exec: interp.DefaultExec,
		Open: interp.OpenDevImpls(interp.DefaultOpen),

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if err = r.Reset(); err != nil {
		return err
	}
	err = r.Run(p)
	return err
}

// buildEnvironment returns the amended env list
func (cwe *CallWithEnvironment) buildEnvironment() []string {
	current := os.Environ()
	for k, v := range cwe.Environment {
		current = append(current, fmt.Sprintf("%s=%s", k, v))
	}
	return current
}
