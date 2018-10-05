# cwe

... is a tool to build software within a mono repository. Though you might use cwe to run the build it self it is best combined with another build tool like [task](https://github.com/go-task/task) or make.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

cwe is build with Go. So you need to have a Go environment up and running. Support for Go with modules is planned but not in place. See Go's [Getting Started](https://golang.org/doc/install) to setup your Go environment.

### Installing

To get the code of cwe you can run go get:

    go get -u github.com/monobuild/cwe

Within `$GOPATH/src/github.com/monobuild/cwe` you should be able to run a test:

    go run cmd/cwe/main.go

## Running the tests

TODO

## Deployment

You can download the binary from the releases page or use the deb package to install it on a Debian system.

## Usage

### How does cwe work

Let's start with a sample:

    cwe --extra-env a=b c=d -- echo ${TEST} ${a} ${c}

with a .cwe.env containing the following data:

    env:
      TEST: "Hello world!"

the result is:

    Hello World! b d

### Command line parameters

`--extra-env` allows to add an additional environment variable using the commandline

`--quiet` makes cwe no printing out own information

Pass `--` before the real command to have cwe not parsing program's argument

### What is in a .cwe.env file

The .cwe.env file contains a dictionary named env and is serialized in YAML ( Yet Another Markup Language ).

#### Windows

On Windows the file is called _cwe.env

### Customization

command line options, config file, etc

## History

|Version|Description|
|---|---|
|1.0.1|Move namespace|
|1.0.0|Initial version|

## Contributing

Please read [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/sascha-andres/cwe/tags).

## Authors

* **Sascha Andres** - *Initial work* - [sascha-andres](https://github.com/sascha-andres)

See also the list of [contributors](https://github.com/sascha-andres/cwe/contributors) who participated in this project.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* [Contributor Covenant](https://www.contributor-covenant.org/) as the source for the code of conduct
* [PurpleBooth](https://github.com/PurpleBooth) for the [README blueprint](https://gist.githubusercontent.com/PurpleBooth/109311bb0361f32d87a2/raw/8254b53ab8dcb18afc64287aaddd9e5b6059f880/README-Template.md)
