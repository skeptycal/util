package getpage

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	defaultInitialPageCount = 8
)

var (
	pageSet *SafePageSet = nil
)

func init() {
	if pageSet == nil {
		pageSet = NewPageSet("pageset", 30)
	}
}

// NewPageSet creates a new SafePageSet. cacheTime is the time.Duration
// that pages are returned directly from the cache.
// Set to:
//  -1 to disable the cache.
//   0 to use the default cache age.
func NewPageSet(name string, cacheTime time.Duration) *SafePageSet {

	if cacheTime == 0 {
		cacheTime = defaultMaxPageCacheAge
	}

	return &SafePageSet{
		name:      fmt.Sprintf("pageset%v", time.Now()),
		cacheTime: defaultMaxPageCacheAge,
		pages:     make(map[string]*strings.Builder, defaultInitialPageCount),
	}
}

type PageSet interface {
	New(url string) *strings.Builder
	Add(url string) error
	Len() int
}

// SafePageSet maintains a thread safe pool of string builders.
//
// Reference: https://tour.golang.org/concurrency/9
type SafePageSet struct {
	name      string        `default:"pageset"`
	cacheTime time.Duration `default=defaultMaxPageCacheAge`
	mu        sync.Mutex
	pages     map[string]*strings.Builder
}

func (s *SafePageSet) New(url string) *strings.Builder {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, n := range s.pages {
		if n == nil {
			return n
		}
	}
	s.pages[url] = &strings.Builder{}
	return s.pages[url]
}

func (s *SafePageSet) String() {
	for k, v := range s.pages {
		fmt.Printf(" %s ...  %v\n", k, v)
	}
}

func (s *SafePageSet) Len() int {
	return len(s.pages)
}
