// section: test
package keys_test

import (
	"testing"

	"github.com/gopherguides/learn/_training/testing/async/src/keys"
)

func TestStoreExercise(t *testing.T) {
	store := keys.NewStore()

	t.Log("store.Set...")
	store.Set("foo", "bar")

	t.Log("store.Get...")
	v, err := store.Get("foo")

	if err != nil {
		t.Fatal(err)
	}

	s, ok := v.(string)
	if !ok {
		t.Fatalf("unexpected value type.  got %T, exp: string", v)
	}

	if got, exp := s, "bar"; got != exp {
		t.Fatalf("unexpected value.  got: %s, exp: %s", got, exp)
	}
}

// section: test

/*
// section: output
$ go test . -v -run TestStoreExercise -count 1
=== RUN   TestStoreExercise
--- FAIL: TestStoreExercise (0.00s)
    exercise_test.go:13: store.Set...
    exercise_test.go:16: store.Get...
    exercise_test.go:20: foo not found
FAIL
FAIL    github.com/gopherguides/learn/_training/testing/async/src/keys  0.006s

// section: output
*/
