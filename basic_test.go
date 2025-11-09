package head_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/head"
)

func ExampleHead_basic() {
	// echo "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n12" | head
	gloo.MustRun(
		Head(strings.NewReader("1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n12")),
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
