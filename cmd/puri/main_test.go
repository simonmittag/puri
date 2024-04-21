package main

import (
	"flag"
	"os"
	"testing"
)

func TestMainFuncWithHelp(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = []string{os.Args[0], "-h"}
	main()
}

func TestMainFuncWithVersion(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = []string{os.Args[0], "-v"}
	main()
}

func TestMainFuncWithPayloadandUriParam(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = []string{os.Args[0], "-q", "k", "https://www.google.com?k=v"}
	main()
}

func TestMainFuncWithPayloadandScheme(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = []string{os.Args[0], "-s", "https://www.google.com?k=v"}
	main()
}

func TestMainFuncWithPayloadandHost(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = []string{os.Args[0], "-o", "https://www.google.com?k=v"}
	main()
}

func TestMainFuncWithPayloadandPath(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = []string{os.Args[0], "-p", "https://www.google.com/path/?k=v"}
	main()
}

func TestMainFuncWithPayloadandPort(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = []string{os.Args[0], "-r", "https://www.google.com:8080?k=v"}
	main()
}

func TestParseFlags(t *testing.T) {
	err := ParseFlags()
	if err != nil {
		t.Error(err)
	}
}

func TestParseFlagSet(t *testing.T) {
	flagset := flag.NewFlagSet("test", flag.ExitOnError)
	flagset.String("p", "", "extract uri param")
	err := flagset.Parse([]string{"-p", "test"})
	if err != nil {
		t.Error(err)
	}
}

func TestParseArgs(t *testing.T) {
	valid := []string{"http://localhost"}
	_, err := parseArgs(valid)
	if err != nil {
		t.Error(err)
	}

	invalid := []string{}
	_, err = parseArgs(invalid)
	if err == nil {
		t.Error("should fail")
	}
}

func TestPrintUsage(t *testing.T) {
	printUsage()
}

func TestPrintVersion(t *testing.T) {
	printVersion()
}
