package main

import (
	"github.com/pkg/errors"
)

func main() {
	errors.Wrapf(nil, "cmd+click me")
}
