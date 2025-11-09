package head_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/head"
)

func ExampleHead_lines() {
	// echo "1\n2\n3\n4\n5" | head -n 3
	gloo.MustRun(
		Head(Lines(3), strings.NewReader("1\n2\n3\n4\n5")),
	)
	// Output:
	// 1
	// 2
	// 3
}
