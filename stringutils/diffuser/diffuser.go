// Package diffuser implements a stream and interprets text.
//
// Diffuse (verb) spread or cause to spread over a wide area or among a large number of people.
// Diffuse (adj) spread out over a large area; not concentrated.
package diffuser

import (
	"fmt"
	"strings"
	"sync"
)

var (
	// mu sync.Mutex                // for loopBad
	cond sync.Cond // for loopGood

	mutex = &sync.Mutex{} // from medium article
)

type stringMutex struct {
	list []string
	mu   sync.Mutex
}

type builders struct {
	pool []*builder
	mu   sync.Mutex
}

func (b *builders) add(sb *builder) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.pool = append(b.pool, sb)
}

type builder struct {
	sb *strings.Builder
}

func mediumUpdate(item string) {
	mutex.Lock()
	// Update shared variable (e.g. slice, pointer on a structure, etc.)
	if item == "" {
		item = "item" + fmt.Sprintf("%v", &item)
	}
	mutex.Unlock()
}

// loopBad is an example of a loop that uses a mutex with
// constant locks and unlocks.
//
// The article discusses the broadcasting functionality of sync.Cond,
//but overlooks its primary usecase: synchronizing based on a condition
// (it's right in the name!).
//
// The mediumUpdate function above is from the article.
//
// You can think of sync.Cond as solving this problem: "I want to take action as soon as some condition is satisfied, but I can't just write a spin-loop, because those are inefficient. What I really want is to check the condition whenever it may have changed, and otherwise wait."
//
// That's why sync.Cond is almost always called in a loop. Imagine we have some goroutines that are processing a slice of items, and in another goroutine, we want to take some action when len(items) == 0. We could write a spin-loop:
func loopBad(smu *stringMutex) {
	for done := false; !done; {
		smu.mu.Lock()
		// process 'items' ...
		done = len(smu.list) == 0
		smu.mu.Unlock()
	}
}

// loopGood is an example of a loop that uses cond.
//
// But this is obviously terrible; we'll be constantly locking
// and unlocking. Instead, we should use a sync.Cond
//
// Then, any time a goroutine modifies items, it should call
// cond.Signal(). This way, we can ensure that this loop only
// runs when the condition might be satisfied.
func loopGood(smu *stringMutex) {
	cond.L.Lock()
	for len(smu.list) != 0 {
		cond.Wait()
		cond.Signal()
	}
	cond.L.Unlock()
}

func makeItems(n int) (string, []string) {
	var itemValue = "foo"
	var item = itemValue

	var items = make([]string, n)

	// Create slice of pointers to strings
	for i := 0; i < n; i++ {
		tmp := fmt.Sprintf("%v-%d", item, i)
		items = append(items, tmp)
	}

	return item, items
}

func makeStrings(n int) *stringMutex {

	smu := stringMutex{}
	smu.mu.Lock()
	smu.list = make([]string, n)
	smu.mu.Unlock()

	return &smu
}

func Tryout() {

	bs := builders{}
	bs.add(&builder{&strings.Builder{}})

	item, _ := makeItems(8)

	mediumUpdate(item)

	items := makeStrings(8)

	loopBad(items)

	loopGood(items)

}

/*

Reference: https://teivah.medium.com/a-closer-look-at-go-sync-package-9f4e4a28c35a

The article discusses the broadcasting functionality of sync.Cond, but overlooks its primary usecase: synchronizing based on a condition (it's right in the name!).

You can think of sync.Cond as solving this problem: "I want to take action as soon as some condition is satisfied, but I can't just write a spin-loop, because those are inefficient. What I really want is to check the condition whenever it may have changed, and otherwise wait."

That's why sync.Cond is almost always called in a loop. Imagine we have some goroutines that are processing a slice of items, and in another goroutine, we want to take some action when len(items) == 0. We could write a spin-loop:

*/
