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
	err := ParseFlagSet(flagset, []string{"-p", "test"})
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
