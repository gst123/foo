package foo

import "fmt"

func Foo(s string) string {
	return fmt.Sprintf("foo v2:%s", s)
}
