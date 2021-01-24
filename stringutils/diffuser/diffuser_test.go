// Reference
/*
Reference: https://teivah.medium.com/a-closer-look-at-go-sync-package-9f4e4a28c35a

The article discusses the broadcasting functionality of sync.Cond, but overlooks its primary usecase: synchronizing based on a condition (it's right in the name!).

You can think of sync.Cond as solving this problem: "I want to take action as soon as some condition is satisfied, but I can't just write a spin-loop, because those are inefficient. What I really want is to check the condition whenever it may have changed, and otherwise wait."

That's why sync.Cond is almost always called in a loop. Imagine we have some goroutines that are processing a slice of items, and in another goroutine, we want to take some action when len(items) == 0. We could write a spin-loop:

*/

package diffuser

import "testing"

func BenchmarkMakeItems(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _, _ = makeItems(8)
    }
}

func BenchmarkMakeStrings(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = makeStrings(8)
    }
}

func BenchmarkLoopBad(b *testing.B) {
    items := makeStrings(8)

    for i := 0; i < b.N; i++ {
        loopBad(items)
    }
}

func BenchmarkLoopGood(b *testing.B) {
    items := makeStrings(8)

    for i := 0; i < b.N; i++ {
        loopGood(items)
    }
}
