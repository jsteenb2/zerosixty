package zerosixty_test

import (
	"fmt"
	"testing"
)

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

func odder(anInt int) int {
	return anInt % 1
}

func adder(anInt, anotherInt, thirdInt int) int {
	// input params that share types can have the type associated to the last param
	// all params without type in front of that param get same type associated
	return anInt + anotherInt + thirdInt
}

func TestBasicFuncs(t *testing.T) {
	t.Log(odder(3))
	t.Log(adder(1, 2, 3))
}

func variadicFunc(msg string, ints ...int) (output string) {
	// named output type used here to showcase you can indeed name an output
	var sum int
	for _, v := range ints {
		sum += v
	}
	output = fmt.Sprintf("%s: %d", msg, sum)
	return
}

func TestVariadicFunc(t *testing.T) {
	fn := variadicFunc // can set funcs to a var

	t.Log(fn("what's my total", 1, 2, 3, 4))
}

func TestFuncClosure(t *testing.T) {
	msg := "I'm getting closured"

	fn := func(anInt int) string {
		return fmt.Sprintf("%s\n\tinput: %d", msg, anInt)
	}

	t.Log(fn(3333))
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
