package zerosixty_test

import "testing"

/*
	lexical scope in Go follows any {} pair
	most nested scope will always override a more outter scope
*/

func TestLexicalScope(t *testing.T) {
	lexical := "initial scope"
	{
		lexical := "second scope"
		t.Log(lexical)
		{
			lexical := "third scope"
			t.Log(lexical)
			{
				lexical := "fourth scope"
				t.Log(lexical)
				if lexical := "if scope"; lexical != "" {
					t.Log(lexical)
				}
			}
			t.Log(lexical)
		}
		t.Log(lexical)
	}
	t.Log(lexical)
}

/*
	func types can be used in similar ways to an interface
	achieve polymorphic behavior without structs w/ methods on it
*/

type aggregator func(...int) int

func sum(ints ...int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

func max(ints ...int) int {
	max := 0
	for _, v := range ints {
		if v > max {
			max = v
		}
	}
	return max
}

func min(ints ...int) int {
	min := ints[0]
	for _, v := range ints[1:] {
		if min > v {
			min = v
		}
	}
	return min
}

func count(ints ...int) int {
	return len(ints)
}

func avg(ints ...int) int {
	return sum(ints...) / len(ints)
}

func TestFuncTypes(t *testing.T) {
	ints := []int{0, 3, 1, 2, 31, -3, 11}

	// running table tests
	aggregators := []struct {
		expected int
		name     string
		aggregator
	}{
		{45, "sum", sum},
		{-3, "min", min},
		{31, "max", max},
		{7, "count", count},
		{6, "avg", avg},
	}

	for _, tc := range aggregators {
		if output := tc.aggregator(ints...); output == tc.expected {
			t.Logf("%s:\t%d", tc.name, output)
		}
	}
}
