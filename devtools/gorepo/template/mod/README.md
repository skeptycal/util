# {{reponame}}

>Tricky and fun utilities for Go programs.

---
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/skeptycal/{{reponame}}/Go) ![Codecov](https://img.shields.io/codecov/c/github/skeptycal/{{reponame}})
 [![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](code-of-conduct.md)

![Twitter Follow](https://img.shields.io/twitter/follow/skeptycal.svg?label=%40skeptycal&style=social) ![GitHub followers](https://img.shields.io/github/followers/skeptycal.svg?style=social)

---

## A Bit of Lore


---

## Getting Started

### Prerequisites

Go ... [latest stable release][go]

If you use macOS, you can use [HomeBrew][brew]:

    brew install go

That's it ... Go is cool like that. It just works.

This installs everything most users will ever need, including a linter, formatter, test suite, performance tuning, debugging, and build tools for all environments.

There is also an [extension suite for VSCode][vscode] that is amazing. It includes integration to the above tools plus intellisense, hover information, snippets, autocompletes, autoformatting, automatic tracking of imports, automatic tool runs on save, automatic testing on save, semantic colors, and more.

Developed and tested on the latest stable release of Go for macOS. Check out the [Go Blog][goblog] for interesting information.

"Should" work on Windows and Linux because Go is cool that way.

---

### Installation

The easiest way is [go get ][goget]. This downloads the package, any required dependencies, and installs everything into your $GOPATH automatically.

```sh
# add repo to $GOPATH
go get -u github.com/skeptycal/{{reponame}}

cd ${GOPATH}/src/github.com/skeptycal/{{reponame}}

# run tests if you want
./go.test.sh

```

---

### Basic Usage

>This is a copy of the sample `main.go` available in the `sample` folder:

```go
package main

import "github.com/skeptycal/{{reponame}}"

func main() {
    // ... do stuff here ...

}

```

To try it out:

```bash
# change to the sample folder
cd _example

# run the main.go program
go run ./main.go

```

---

## Code of Conduct and Contributing

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us. Please read the [Code of Conduct](CODE_OF_CONDUCT.md) for details before submitting anything.

---

## Versioning

We use SemVer for versioning. For the versions available, see the tags on this repository.

---

## Contributors
- Michael Treanor ([GitHub][github] / [Twitter][twitter]) - Initial work, updates, maintainer
- [Francesc Campoy][Campoy] - He is an inspiration and great YouTube videos on [justforfunc][justforfunc]!

See also the list of contributors who participated in this project.

Great Go packages that I use much of the time:

- Color - A well documented [color package][color] - by [fatih][fatih]
- [Fiber][fiber] - ⚡️ Express inspired web framework written in Go
- [Fast HTTP][fast] - Fast HTTP package for Go
- [EasyJSON][easy] - Fast JSON serializer for golang

---

## License

Licensed under the MIT License <https://opensource.org/licenses/MIT>. See the [LICENSE](LICENSE) file for details.


[twitter]: (https://www.twitter.com/skeptycal)
[github]: (https://github.com/skeptycal)
[Campoy]: (https://github.com/campoy)
[color]: (https://github.com/fatih/color)
[fatih]: (https://github.com/fatih)
[fiber]: (https://github.com/gofiber/fiber)
[fast]: (https://github.com/valyala/fasthttp)
[easy]: (https://github.com/mailru/easyjson)
[justforfunc]: (https://www.youtube.com/c/JustForFunc)
[goget]: (https://golang.org/pkg/cmd/go/internal/get/)
[vscode]: (https://code.visualstudio.com/docs/languages/go)
[goblog]: (https://blog.golang.org/)
[brew]: (https://brew.sh/)
[go]: (https://golang.org/)
