package main

import (
	"flag"
	"testing"
)

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
