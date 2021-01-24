// Package diffuser implements a service that streams and interprets text.
//
// Diffuse (verb) spread or cause to spread over a wide area or among a large number of people.
// Diffuse (adj) spread out over a large area; not concentrated.
package diffuser

import (
	"strings"
	"sync"
)


type builders struct {
    pool []strings.Builder
    sync.WaitGroup
}
type diffuser struct {
    sb *strings.Builder
}
