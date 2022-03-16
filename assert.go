package interval

import (
	"fmt"
	"testing"

	. "golang.org/x/exp/slices"
)

/*
	Private Print Format function.

	Parameters:
		expected T	What the value should be
		got T 		What the value is

	Return:
		string
*/
func preFormattedErrorString[T any](got, expected T) string {
	return fmt.Sprintf("expected: %+v, got: %+v", expected, got)
}

func AssertTrue(value bool, t *testing.T) {
	t.Run(fmt.Sprintf("AssertTrue : %v", value), func(t *testing.T) {
		if value != true {
			t.Error(preFormattedErrorString(value, true))
		}
	})
}

func AssertFalse(value bool, t *testing.T) {
	t.Run(fmt.Sprintf("AssertFalse - %v", value), func(t *testing.T) {
		if value != false {
			t.Error(preFormattedErrorString(value, false))
		}
	})
}

func AssertEqual[T comparable](value, expected T, t *testing.T) {
	t.Run(fmt.Sprintf("AssertEqual - %v == %v", value, expected), func(t *testing.T) {
		if value != expected {
			t.Error(preFormattedErrorString(value, expected))
		}
	})
}

func AssertEqualSlice[T comparable](value, expected []T, t *testing.T) {
	t.Run(fmt.Sprintf("AssertEqualSlice - %v == %v", value, expected), func(t *testing.T) {
		if !Equal(value, expected) {
			t.Error(preFormattedErrorString(value, expected))
		}
	})
}

func AssertNotEqualSlice[T comparable](value, expected []T, t *testing.T) {
	t.Run(fmt.Sprintf("AssertNotEqualSlice - %v != %v", value, expected), func(t *testing.T) {
		if Equal(value, expected) {
			t.Error(preFormattedErrorString(value, expected))
		}
	})
}

func AssertNotEqual[T comparable](value, expected T, t *testing.T) {
	t.Run(fmt.Sprintf("AssertNotEqual - %v != %v", value, expected), func(t *testing.T) {
		if value == expected {
			t.Error(preFormattedErrorString(value, expected))
		}
	})
}

func AssertNil[T comparable](value *T, t *testing.T) {
	t.Run(fmt.Sprintf("AssertNil - %v", value), func(t *testing.T) {
		if value != nil {
			t.Error(preFormattedErrorString(value, nil))
		}
	})

}

func AssertNotNil[T comparable](value *T, t *testing.T) {
	t.Run(fmt.Sprintf("AssertNotNil - %v", value), func(t *testing.T) {
		if value == nil {
			t.Error(preFormattedErrorString("nil", "not nil"))
		}
	})
}
