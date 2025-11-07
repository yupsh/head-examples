package head_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/head"
)

func ExampleHead_fromFile_lines() {
	// head -n 3 testdata/numbers.txt
	yup.MustRun(
		Head(Lines(3), yup.File("testdata/numbers.txt")),
	)
	// Output:
	// 1
	// 2
	// 3
}

