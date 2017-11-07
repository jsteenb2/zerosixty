// go test -run none -bench . -benchtime 3s -benchmem
package testing_test

import (
	"fmt"
	"testing"
)

var str string

// BenchmarkSprint tests the performance of using Sprint.
func BenchmarkSprint(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}

	str = s
}

// BenchmarkSprint tests the performance of using Sprintf.
func BenchmarkSprintf(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}

	str = s
}
