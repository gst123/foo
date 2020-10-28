package foo

import "fmt"

func Foo(s string) string {
	return fmt.Sprintf("foo:%s", s)
}
