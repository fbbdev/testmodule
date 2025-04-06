package testmodule

import "github.com/fbbdev/testmodule/subpkg"

func Greeting() string {
	return subpkg.Greeting
}
