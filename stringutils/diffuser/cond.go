/*

Reference: https://teivah.medium.com/a-closer-look-at-go-sync-package-9f4e4a28c35a

The article discusses the broadcasting functionality of sync.Cond, but overlooks its primary usecase: synchronizing based on a condition (it's right in the name!).

You can think of sync.Cond as solving this problem: "I want to take action as soon as some condition is satisfied, but I can't just write a spin-loop, because those are inefficient. What I really want is to check the condition whenever it may have changed, and otherwise wait."

That's why sync.Cond is almost always called in a loop. Imagine we have some goroutines that are processing a slice of items, and in another goroutine, we want to take some action when len(items) == 0. We could write a spin-loop:

*/

package diffuser

import "sync"

var (
    mu sync.Mutex                // for loopBad
    cond sync.Cond               // for loopGood

    mutex = &sync.Mutex{} // from medium article

)


func mediumUpdate(item *string) {
    mutex.Lock()
    // Update shared variable (e.g. slice, pointer on a structure, etc.)
    if *item == "" {
        *item = "item"
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
func loopBad(items []*string) {
    for done := false; !done; {
        mu.Lock()
        // process 'items' ...
        done = len(items) == 0
        mu.Unlock()
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
func loopGood(items []*string) {
    cond.L.Lock()
    for len(items) != 0 {
        cond.Wait()
    }
    cond.L.Unlock()
}
