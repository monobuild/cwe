// Copyright © 2018 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"livingit.de/code/cwe"
	"livingit.de/code/cwe/cmd/cwe/methods"
	"os"
	"strings"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "no program to run provided")
		os.Exit(1)
	}
	program := strings.Join(args, " ")

	c, err := cwe.NewCallWithEnvironment()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	var quiet bool

	if !strings.HasPrefix(program, "\"") {
		if strings.Contains(program, " -- ") {
			var myFlags arrayFlags

			splittedArguments := strings.SplitN(program, " -- ", 2)
			program = strings.TrimSpace(splittedArguments[1])

			fs := flag.NewFlagSet("extraenv", flag.ExitOnError)
			fs.Var(&myFlags, "extra-env", "Allow adding environment variable through cli")
			fs.BoolVar(&quiet, "quiet", false, "Set to true to hide information")

			err := fs.Parse(strings.Split(splittedArguments[0], " "))
			if err != nil {
				return
			}
			for _, v := range myFlags {
				if strings.Contains(v, "=") {
					splittedValue := strings.SplitN(v, "=", 2)
					c.Add(splittedValue[0], splittedValue[1])
				}
			}
		}
	}

	if quiet {
		c.Quiet = true
	} else {
		methods.PrintHeader()
	}

	c.Run(program)
}
