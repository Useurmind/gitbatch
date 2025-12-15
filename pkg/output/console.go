package output

import (
	"fmt"
	"os"
)

func CheckErrf(err error, msg string, args... interface{}) {
	if err != nil {
		allArgs := append(args, err)
		Writeln("Error: " + msg + " failed with %v", allArgs...)
		os.Exit(1)
	}
}

func Writef(msg string, args... interface{}) {
	fmt.Printf(msg, args...)
}

func Writeln(msg string, args... interface{}) {
	Writef(msg + "\r\n", args...)
}

func WriteSeparator() {
	Writeln("-----------------------------------------------------------------------")
}