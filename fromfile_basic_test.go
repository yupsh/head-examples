package head_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/head"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleHead_fromFile_basic() {
	// head testdata/numbers.txt
	yup.MustRun(
		Head(yup.File("testdata/numbers.txt")),
	)
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}

