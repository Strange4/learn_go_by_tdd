package reflection

import (
	"testing"
)

type Person struct {
	Name string
	Info Info
}

type Info struct {
	Age      int
	LastName string
}

func TestWalk(t *testing.T) {

	testCases := []struct {
		name          string
		input         any
		expectedCalls []string
	}{
		{
			"With only 1 value",
			struct {
				Name string
			}{"Me"},
			[]string{"Me"},
		},
		{
			"With multiple string values",
			struct {
				A string
				B string
				C string
			}{
				A: "something",
				B: "something else",
				C: "last thing",
			},
			[]string{"something", "something else", "last thing"},
		},
		{
			"With non string field",
			struct {
				MyName   string
				YourName string
				MyAge    int
				YourAge  int
			}{
				"me",
				"you",
				22,
				23,
			},
			[]string{"me", "you"},
		},
		{
			"nested structs",
			Person{
				"me",
				Info{
					22,
					"myself",
				},
			},
			[]string{"me", "myself"},
		},
		{
			"pointer to struct",
			&Person{
				"Chris",
				Info{33, "Beckham"},
			},
			[]string{"Chris", "Beckham"},
		},
		{
			"slices",
			[]Info{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"arrays",
			[2]Info{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"maps",
			map[string]string{
				"Coffee":  "Cow",
				"La Vaca": "Moo",
			},
			[]string{"Coffee", "Cow", "La Vaca", "Moo"},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			var got []string
			Walk(test.input, func(input string) {
				got = append(got, input)
			})
			arraysAreEqualWithoutOrder(t, test.expectedCalls, got)
		})
	}
}

func arraysAreEqualWithoutOrder(t *testing.T, want, got []string) {
	t.Helper()
	if len(want) != len(got) {
		t.Fatalf("expected %d elements but got %d", len(want), len(got))
	}

	stringCount := map[string]int{}
	for _, item := range got {
		count, exists := stringCount[item]
		if exists {
			stringCount[item] = count + 1
		} else {
			stringCount[item] = 1
		}
	}
	for _, item := range want {
		count, exists := stringCount[item]
		if count == 0 {
			t.Fatalf("item %q appeared less times than wanted", item)
		}
		if !exists {
			t.Fatalf("wanted item %q in the list but wasn't there", item)
		}
		stringCount[item] = count - 1
	}

	for key, count := range stringCount {
		if count != 0 {
			t.Fatalf("item %q appeared more times than wanted", key)
		}
	}
}
