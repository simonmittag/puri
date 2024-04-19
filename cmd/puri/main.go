package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/simonmittag/puri"
	"net/url"
	"os"
)

type Mode uint8

const (
	All Mode = 1 << iota
	Scheme
	Host
	Port
	Path
	Param
	Version
	Usage
)

func main() {
	mode := Usage
	uri := ""

	p := flag.String("p", "", "extract uri param")
	o := flag.Bool("o", false, "extract host")
	r := flag.Bool("r", false, "extract port")
	s := flag.Bool("s", false, "extract scheme")
	v := flag.Bool("v", false, "print puri version")
	h := flag.Bool("h", false, "print usage instructions")
	flag.Usage = printUsage
	err := ParseFlags()
	if err != nil || *h {
		mode = Usage
	} else {
		a := flag.Args()
		uri, err = parseArgs(a)

		if *v {
			mode = Version
		} else if *p != "" {
			mode = Param
		} else if *s {
			mode = Scheme
		} else if *o {
			mode = Host
		} else if *r {
			mode = Port
		}
	}

	switch mode {
	case Param:
		handleOutput(puri.ExtractParam(uri, *p))
	case Scheme:
		handleOutput(puri.ExtractScheme(uri))
	case Port:
		handleOutput(puri.ExtractPort(uri))
	case Host:
		handleOutput(puri.ExtractHost(uri))
	case Usage:

		printUsage()
	case Version:

		printVersion()
	}
}

func handleOutput(res string, err error) {
	if err == nil {
		fmt.Println(res)
	} else {
		printUsage()
	}
}

func printUsage() {
	printVersion()
	fmt.Printf("Usage: puri [-h]|[-v][-p name]| scheme://host:port?k=v\n")
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Printf("puri[%s]\n", puri.Version)
}

// ParseFlags parses the command line args, allowing flags to be
// specified after positional args.
func ParseFlags() error {
	return ParseFlagSet(flag.CommandLine, os.Args[1:])
}

// ParseFlagSet works like flagset.Parse(), except positional arguments are not
// required to come after flag arguments.
func ParseFlagSet(flagset *flag.FlagSet, args []string) error {
	var positionalArgs []string
	for {
		if err := flagset.Parse(args); err != nil {
			return err
		}
		// Consume all the flags that were parsed as flags.
		args = args[len(args)-flagset.NArg():]
		if len(args) == 0 {
			break
		}
		// There's at least one flag remaining and it must be a positional arg since
		// we consumed all args that were parsed as flags. Consume just the first
		// one, and retry parsing, since subsequent args may be flags.
		positionalArgs = append(positionalArgs, args[0])
		args = args[1:]
	}
	// Parse just the positional args so that flagset.Args()/flagset.NArgs()
	// return the expected value.
	// Note: This should never return an error.
	return flagset.Parse(positionalArgs)
}

func parseArgs(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("must supply uri")
	}
	_, err := url.Parse(args[0])
	if err != nil {
		return "", err
	}
	return args[0], nil
}
