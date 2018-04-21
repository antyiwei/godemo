package main

import (
	"github.com/mpvl/errc"
	"github.com/pkg/errors"
)

func main() {

	err := errors.New(" no err ....")

	// 暂时不知道怎么使用errc
	errc.Catch(&err, nil)
}
