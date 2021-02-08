package gogen

import "testing"

// Notes from good Go Testing video
// Reference: https://www.youtube.com/watch?v=ndmB0bj7eyw&ab_channel=TheGoProgrammingLanguage
/*

- t.Parallel() // top of function to allow parallel execution

- slice of structs for testing input

- test coverage
    - go test -coverprofile=cover.out
    - go tool cover -func=cover.out
    - go tool cover -html=cover.out (nice looking)

*/

func TestParallel(t *testing.T) {
    t.Parallel()
}
