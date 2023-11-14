package svc

import "fmt"

var GVerbose bool

func verboseFln(format string, a ...any) {
	if GVerbose {
		_, _ = fmt.Printf(format+"\n", a...)
	}
}
