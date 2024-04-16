`puri` ("parse uri") is a commandline utility that reads and extracts uri components,

[![Circleci Builds](https://circleci.com/gh/simonmittag/puri.svg?style=shield)](https://circleci.com/gh/simonmittag/puri)
[![Github Workflows](https://github.com/simonmittag/mse6/workflows/Go/badge.svg)](https://github.com/simonmittag/puri/actions)
[![Github Issues](https://img.shields.io/github/issues/simonmittag/puri)](https://github.com/simonmittag/puri/issues)
[![Github Activity](https://img.shields.io/github/commit-activity/m/simonmittag/puri)](https://img.shields.io/github/commit-activity/m/simonmittag/puri)  
[![CodeClimate Maintainability](https://api.codeclimate.com/v1/badges/70cd59e4dfd2801f8661/maintainability)](https://codeclimate.com/github/simonmittag/puri/maintainability)
[![CodeClimate Test Coverage](https://api.codeclimate.com/v1/badges/70cd59e4dfd2801f8661/test_coverage)](https://codeclimate.com/github/simonmittag/puri/test_coverage)
[![Go Version](https://img.shields.io/github/go-mod/go-version/simonmittag/puri)](https://img.shields.io/github/go-mod/go-version/simonmittag/puri)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Version](https://img.shields.io/badge/version-0.1.4-orange)](https://github.com/simonmittag/puri/releases/tag/v0.1.4)

## What's New
### v0.1.4
* initial revision

## Up and running
### Homebrew
```
brew tap simonmittag/cli &&
  brew install puri &&
  puri 
```

### Golang
```bash
git clone https://github.com/simonmittag/puri && cd puri && 
go install github.com/simonmittag/puri/cmd/puri && 
puri 
```

## Usage
```
λ puri[v0.1.4]
Usage: puri [-h]|[-v]|[-p name] scheme://host:port?k=v
  -h    print usage instructions
  -p string
        extract uri param
  -v    print puri version
```

## Examples

Get URI parm
```
λ puri -p q https://www.google.com?q=blah
  blah
```

## Contributions
The puri team welcomes all [contributors](https://github.com/simonmittag/puri/blob/master/CONTRIBUTING.md). Everyone interacting with the project's codebase, issue trackers, chat rooms and mailing lists
is expected to follow the [code of conduct](https://github.com/simonmittag/puri/blob/master/CODE_OF_CONDUCT.md)