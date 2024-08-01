package util

import (
	"fmt"
	"os"
)

func Fatal(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, "fatal:", err)
	os.Exit(1)
}
