package build

import (
	"fmt"
	"go/build"
	"strings"
)

var buildContext build.Context

type goose struct {
	goos   string
	goarch []string
}

func (g *goose) String() string {
	return fmt.Sprintf("GOOS=${%s} GOARCH=${%v}", g.goos, g.goarch)
}

func (g *goose) Active() (s []string) {
	for _, v := range g.goarch {
		if v == strings.ToLower(buildContext.GOOS) {
			s = append(s, v)
		}
	}
	return
}

// BuildList returns a list of all possible combinations of GOOS and GOARCH.
// get this list from
//      go tool dist list
// Ref: https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63
func BuildList() []goose {
	return []goose{
		{"aix", []string{"ppc64"}},
		{"android", []string{"arm64"}},
		{"darwin", []string{"amd64"}},
		{"dragonfly", []string{"amd64"}},
		{"freebsd", []string{"386", "amd64", "arm", "arm64"}},
		{"illumos", []string{"amd64"}},
		{"js", []string{"wasm"}},
		{"linux", []string{"386", "amd64", "arm", "arm64", "ppc64", "ppc64le", "mips", "mipsle", "mips64", "mips64le", "riscv64", "s390x"}},
		{"netbsd", []string{"386", "amd64", "arm", "arm64"}},
		{"openbsd", []string{"386", "amd64", "arm", "arm64"}},
		{"plan9", []string{"386", "amd64", "arm"}},
		{"solaris", []string{"amd64"}},
		{"windows", []string{"386", "amd64", "arm"}},
	}
}
