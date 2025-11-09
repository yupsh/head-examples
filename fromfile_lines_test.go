package head_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/head"
)

func ExampleHead_fromFile_lines() {
	// head -n 3 testdata/numbers.txt
	gloo.MustRun(
		Head(Lines(3), gloo.File("testdata/numbers.txt")),
	)
	// Output:
	// 1
	// 2
	// 3
}
