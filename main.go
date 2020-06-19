package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// use the binary name to determine operating mode
//
// h2d: hexadecimal to decimal conversion
// other: decimal to hexadecimal conversion
func mode() int {
	bn := filepath.Base(strings.ToLower(os.Args[0]))
	bin := strings.TrimSuffix(bn, filepath.Ext(bn))

	if bin == "h2d" {
		return 1
	}

	return 0
}

func needful(s string, base int, format string, printNewline bool) {
	i64, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse input as base%d integer: %s\n", base, err)
		os.Exit(1)
	}

	fmt.Printf(format, i64)

	if printNewline {
		fmt.Println()
	}
}

func main() {
	var noNewline bool

	flag.BoolVar(&noNewline, "n", false, "don't print newline with output")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "must provide either an integer or hexadecimal value as the first positional argument")
		os.Exit(1)
	}

	input := flag.Args()[0]

	switch mode() {
	case 0:
		needful(input, 10, "%X", !noNewline) // dec2hex
	case 1:
		needful(input, 16, "%d", !noNewline) // hex2dec
	}
}
