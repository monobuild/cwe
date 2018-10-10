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
	variables, err := cwe.buildEnvironment()
	if err != nil {
		return err
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	p, err := syntax.NewParser().Parse(strings.NewReader(program), "")
	if err != nil {
		return err
	}

	r := interp.Runner{
		Dir: dir,
		Env: variables,

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
func (cwe *CallWithEnvironment) buildEnvironment() (interp.Environ, error) {
	env := os.Environ()
	for k, v := range cwe.Environment {
		if !cwe.Quiet {
			fmt.Println(fmt.Sprintf("%s: %s", k, v))
		}
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	if !cwe.Quiet {
		fmt.Println()
	}
	return interp.EnvFromList(env)
}
