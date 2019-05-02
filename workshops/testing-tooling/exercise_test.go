package main

import (
	"testing"
	"testing/quick"
)

func Test_Add(t *testing.T) {
	commutativity := func(a, b int) bool {
		return Add(a, b) == Add(b, a)
	}

	associativity := func(a, b, c int) bool {
		return Add(a, Add(b, c)) == Add(Add(a, b), c)
	}

	identity := func(a int) bool {
		return Add(a, 0) == a
	}

	if err := quick.Check(commutativity, nil); err != nil {
		t.Error(err)
	}

	if err := quick.Check(associativity, nil); err != nil {
		t.Error(err)
	}

	if err := quick.Check(identity, nil); err != nil {
		t.Error(err)
	}
}
